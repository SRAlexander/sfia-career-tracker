package main

import (
	"os"
    "fmt"
    "log"
	"strings"
	"time"
	"io/ioutil"
	"strconv"
	"encoding/json"
    "github.com/xuri/excelize/v2"
)

type Config struct {
    SFIASpreadsheetLocation string `json:"sfiaSpreadsheetLocation"`
    SkillMappingsLocation string `json:"skillMappingsLocation"`
	ResponsibilitiesMappingsLocation string `json:"responsibilitiesMappingsLocation"`
	ProcessedOutputFolder string `json:"processedOutputFolder"`
	ExportFormat string `json:"exportFormat"`
}

type SheetMapping struct {
    SheetName string `json:"SHEETNAME"`
    SkillCodeColumn string `json:"SKILLCODE"`
    LevelIndications []string `json:"LEVEL-INDICATORS"`
    LevelDescriptions []string `json:"LEVEL-DESCRIPTIONS"`
}

type PostSkillDataModel struct {
    SkillCode string
    SFIALevel int
    KeyPointNumber int
    KeyPointDescription string
}

func main() {

	var config Config
	json.Unmarshal(loadJSONFileAsByteString("config.json"), &config)

	var skillSheetMappings SheetMapping
	json.Unmarshal(loadJSONFileAsByteString(config.SkillMappingsLocation), &skillSheetMappings)

	var responsibilitiesSheetMappings SheetMapping
	json.Unmarshal(loadJSONFileAsByteString(config.ResponsibilitiesMappingsLocation), &responsibilitiesSheetMappings)

    file, err := excelize.OpenFile(config.SFIASpreadsheetLocation)
    if err != nil {
        log.Fatal(err)
    }

	var skillModels []PostSkillDataModel
	skillModels = append(skillModels, processSFIAResponsibilities(file, responsibilitiesSheetMappings)...)
	skillModels = append(skillModels, processSFIASkills(file, skillSheetMappings)...)
	
	// output JSON file
	var currentTimestamp = time.Now().Unix()
	if config.ExportFormat == "JSON" {
		exportAsJSON(skillModels, config.ProcessedOutputFolder, currentTimestamp)
	}

	if config.ExportFormat == "EXCEL" {
		exportAsExcel(skillModels, config.ProcessedOutputFolder, currentTimestamp)
	}
}

// The behaviours sheet has a different format to skills
func processSFIAResponsibilities(file *excelize.File, responsibilitiesSheetMappings SheetMapping) []PostSkillDataModel {

	var rowOffset int = 2;
	var skillModels [] PostSkillDataModel

	for levelPosition:=0; levelPosition < len(responsibilitiesSheetMappings.LevelIndications); levelPosition ++ {

		var trueLevel int = levelPosition+1
		var rowOffsetLevel int = levelPosition + rowOffset

		for descriptionPosition:=0; descriptionPosition < len(responsibilitiesSheetMappings.LevelDescriptions); descriptionPosition++ {
			columnLevelDescription, err := file.GetCellValue(responsibilitiesSheetMappings.SheetName, responsibilitiesSheetMappings.LevelDescriptions[descriptionPosition] + strconv.Itoa(rowOffsetLevel))
			if err != nil {
				log.Fatal(err)
			}

			columnAttribute, err := file.GetCellValue(responsibilitiesSheetMappings.SheetName, responsibilitiesSheetMappings.LevelDescriptions[descriptionPosition] + "1")
			if err != nil {
				log.Fatal(err)
			}

			var stage1Formatting string = strings.ReplaceAll(columnLevelDescription, "\r\n", "\n")
			var stage2Formatting string = strings.ReplaceAll(stage1Formatting, "\n", "")
			var keyPoints []string = deleteEmpty(strings.Split(stage2Formatting, "."))
			
			// form into json model
			for keypoint:=0; keypoint < len(keyPoints); keypoint++ {
				var postSkillDataModel PostSkillDataModel
				postSkillDataModel.SkillCode = columnAttribute
				postSkillDataModel.SFIALevel = trueLevel
				postSkillDataModel.KeyPointNumber = keypoint + 1
				postSkillDataModel.KeyPointDescription = strings.TrimSpace(keyPoints[keypoint])
				skillModels = append(skillModels, postSkillDataModel)
			}

		}
	}

	return skillModels
}

// The behaviours sheet has a different format to skills
func processSFIASkills(file *excelize.File, skillSheetMappings SheetMapping) []PostSkillDataModel {

	// SFIA documents have a header column which is not required, start all row counting with +1
	var rowCount int = 2;
	var skillModels [] PostSkillDataModel

	for {
		columnSkill, err := file.GetCellValue(skillSheetMappings.SheetName, skillSheetMappings.SkillCodeColumn + strconv.Itoa(rowCount))
		if err != nil {
			log.Fatal(err)
		}

		if (columnSkill == "") {
			break;
		}

		// Lets work through each SFIA and create keypoints against applicable levels.
		for level:=0; level < len(skillSheetMappings.LevelIndications); level++ {
			var levelIndicationLocation string = skillSheetMappings.LevelIndications[level]
			columnLevelIndicator, err := file.GetCellValue(skillSheetMappings.SheetName, levelIndicationLocation + strconv.Itoa(rowCount))
			if err != nil {
				log.Fatal(err)
			}

			if (columnLevelIndicator == ""){
				continue
			}

			var levelDescriptionLocation string = skillSheetMappings.LevelDescriptions[level]
			columnLevelDescription, err := file.GetCellValue(skillSheetMappings.SheetName, levelDescriptionLocation + strconv.Itoa(rowCount))
			if err != nil {
				log.Fatal(err)
			}

			var stage1Formatting string = strings.ReplaceAll(columnLevelDescription, "\r\n", "\n")
			var stage2Formatting string = strings.ReplaceAll(stage1Formatting, "\n", "")
			var keyPoints []string = deleteEmpty(strings.Split(stage2Formatting, "."))
			
			// form into json model
			for keypoint:=0; keypoint < len(keyPoints); keypoint++ {
				var postSkillDataModel PostSkillDataModel
				postSkillDataModel.SkillCode = columnSkill
				postSkillDataModel.SFIALevel = level
				postSkillDataModel.KeyPointNumber = keypoint + 1
				postSkillDataModel.KeyPointDescription = strings.TrimSpace(keyPoints[keypoint])
				skillModels = append(skillModels, postSkillDataModel)
			}
		}

		rowCount++
	}

	return skillModels
}

func loadJSONFileAsByteString(file string) []byte {

	// Open our jsonFile
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened: " + file) 
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()
	return byteValue
}

func exportAsJSON(skillModels []PostSkillDataModel, outputLocation string, currentTimestamp int64) {

	jsonContent, err := json.Marshal(skillModels)
	if err != nil {
		fmt.Println(err)
	}

	var jsonOutputLocation string = outputLocation + strconv.Itoa(int(currentTimestamp)) + "-SFIA-Skill-Criteria.json"
	err = ioutil.WriteFile(jsonOutputLocation, jsonContent, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func exportAsExcel(skillModels []PostSkillDataModel, outputLocation string, currentTimestamp int64) {
	file := excelize.NewFile()

	file.SetCellValue("Sheet1", "A1", "SkillCode")
    file.SetCellValue("Sheet1", "B1", "SFIA Level")
	file.SetCellValue("Sheet1", "C1", "Keycode")
	file.SetCellValue("Sheet1", "D1", "Description")

	for skillItem := 0; skillItem < len(skillModels); skillItem++ {
		var insertLocation int = skillItem + 2
		file.SetCellValue("Sheet1", "A" + strconv.Itoa(insertLocation), skillModels[skillItem].SkillCode)
		file.SetCellValue("Sheet1", "B" + strconv.Itoa(insertLocation), skillModels[skillItem].SFIALevel)
		file.SetCellValue("Sheet1", "C" + strconv.Itoa(insertLocation), skillModels[skillItem].KeyPointNumber)
		file.SetCellValue("Sheet1", "D" + strconv.Itoa(insertLocation), skillModels[skillItem].KeyPointDescription)
	}

	var jsonOutputLocation string = outputLocation + strconv.Itoa(int(currentTimestamp)) + "-SFIA-Skill-Criteria.xlsx"
	if err := file.SaveAs(jsonOutputLocation); err != nil {
        log.Fatal(err)
	}
}

func deleteEmpty (s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}



