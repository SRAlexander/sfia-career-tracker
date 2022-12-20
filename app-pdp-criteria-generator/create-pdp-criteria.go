package main

import (
	"os"
    "fmt"
    "log"
	"flag"
	"strings"
	"time"
	"io/ioutil"
	"strconv"
	"encoding/json"
    "github.com/xuri/excelize/v2"
)

type Config struct {
    SFIAProcessedSpreadsheetLocation string `json:"sfiaProcessedSpreadsheetLocation"`
	ProcessedOutputFolder string `json:"processedOutputFolder"`
	ExportFormat string `json:"exportFormat"`
	DefaultSkills string `json:"defaultSkills"`
	SkillColumn string `json:"skillColumn"`
	SFIAColumn string `json:"sfiaColumn"`
	KeyCodeColumn string `json:"keyColumn"`
	KeyDescriptionColumn string `json:"keyDescriptionColumn"`
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

	var sfiaLevel int 
	flag.IntVar(&sfiaLevel, "sfia-level", 5, "What SFIA level are you aiming for? ect. If you are a 3 you should be aiming for 4")
	
	flag.Parse()

	var skills []string = []string{}
	if contains(flag.Args(), "CORE") {
		skills = append(flag.Args(), strings.Split(config.DefaultSkills,",")...)
	} else {
		skills = append(flag.Args())
	}

    file, err := excelize.OpenFile(config.SFIAProcessedSpreadsheetLocation)
    if err != nil {
        log.Fatal(err)
    }

	var skillColumn string = config.SkillColumn
	var sfiaColumn string = config.SFIAColumn
	var keyCodeColumn string = config.KeyCodeColumn
	var keyDescriptionColumn string = config.KeyDescriptionColumn

	var rowCount int = 2
	var applicableSkills []PostSkillDataModel

	for {

		columnSkillCode, err := file.GetCellValue("Sheet1", skillColumn + strconv.Itoa(rowCount))
		if err != nil {
			log.Fatal(err)
		}

		
		if contains(skills, columnSkillCode) {
			columnSFIALevel, err := file.GetCellValue("Sheet1", sfiaColumn + strconv.Itoa(rowCount))
			if err != nil {
				log.Fatal(err)
			}

			if columnSFIALevel == strconv.Itoa(sfiaLevel) {

				columnKeyNumber, err := file.GetCellValue("Sheet1", keyCodeColumn + strconv.Itoa(rowCount))
				if err != nil {
					log.Fatal(err)
				}

				columnKeyDescription, err := file.GetCellValue("Sheet1", keyDescriptionColumn + strconv.Itoa(rowCount))
				if err != nil {
					log.Fatal(err)
				}

				var sfiaInt, errConvert = strconv.Atoi(columnKeyNumber)
				if errConvert != nil {
					fmt.Println("Error during conversion")
					return
				}

				var foundSkill PostSkillDataModel
				foundSkill.SkillCode = columnSkillCode
				foundSkill.SFIALevel = sfiaLevel
				foundSkill.KeyPointNumber = sfiaInt
				foundSkill.KeyPointDescription = columnKeyDescription
				applicableSkills = append(applicableSkills, foundSkill) 
			}
		}

		if columnSkillCode == "" {
			break
		}

		rowCount++
	}

	// output JSON file
	var currentTimestamp = time.Now().Unix()
	if config.ExportFormat == "JSON" {
		exportAsJSON(applicableSkills, config.ProcessedOutputFolder, currentTimestamp)
	}

	if config.ExportFormat == "EXCEL" {
		exportAsExcel(applicableSkills, config.ProcessedOutputFolder, currentTimestamp)
	}
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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}


func exportAsJSON(skillModels []PostSkillDataModel, outputLocation string, currentTimestamp int64) {

	jsonContent, err := json.Marshal(skillModels)
	if err != nil {
		fmt.Println(err)
	}

	var jsonOutputLocation string = outputLocation + strconv.Itoa(int(currentTimestamp)) + "-SFIA-Skill-PDP.json"
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

	var jsonOutputLocation string = outputLocation + strconv.Itoa(int(currentTimestamp)) + "-SFIA-Skill-PDP.xlsx"
	if err := file.SaveAs(jsonOutputLocation); err != nil {
        log.Fatal(err)
	}
}

// func deleteEmpty (s []string) []string {
// 	var r []string
// 	for _, str := range s {
// 		if str != "" {
// 			r = append(r, str)
// 		}
// 	}
// 	return r
// }



