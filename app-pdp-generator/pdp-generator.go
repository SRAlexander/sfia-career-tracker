package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	// "strings"
	// "time"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Config struct {
	TemplateLocation      string `json:"templateLocation"`
	ProcessedOutputFolder string `json:"processedOutputFolder"`
	SkillColumn           string `json:"skillColumn"`
	SFIAColumn            string `json:"sfiaColumn"`
	KeyCodeColumn         string `json:"keyColumn"`
	KeyDescriptionColumn  string `json:"keyDescriptionColumn"`
}

type SkillDataModel struct {
	SkillCode           string
	SFIALevel           string
	KeyPointNumber      string
	KeyPointDescription string
}

type SkillResponse struct {
	Skills         []string
	DetailedSkills []SkillDataModel
}

func main() {

	var config Config
	json.Unmarshal(loadJSONFileAsByteString("config.json"), &config)

	var sfiaSkillsFile string
	flag.StringVar(&sfiaSkillsFile, "skill-list", "./", "Where is your skills file?")

	var outputFilename string
	flag.StringVar(&outputFilename, "output-filename", "./", "What would you like to name the output file?")
	flag.Parse()

	var skillModel = generateSkills(config, sfiaSkillsFile)
	var createdPDP = generatePDP(skillModel, config.TemplateLocation)
	savePDP(config, createdPDP, outputFilename)

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

func generateSkills(config Config, sfiaSkillsFile string) SkillResponse {

	file, err := excelize.OpenFile(sfiaSkillsFile)
	if err != nil {
		log.Fatal(err)
	}

	var skillColumn string = config.SkillColumn
	var sfiaColumn string = config.SFIAColumn
	var keyCodeColumn string = config.KeyCodeColumn
	var keyDescriptionColumn string = config.KeyDescriptionColumn

	var rowCount int = 2

	var skills []string
	var detailedSkills []SkillDataModel

	for {

		columnSkillCode, err := file.GetCellValue("Sheet1", skillColumn+strconv.Itoa(rowCount))
		if err != nil {
			log.Fatal(err)
		}

		if columnSkillCode == "" {
			break
		}

		if !contains(skills, columnSkillCode) {
			skills = append(skills, columnSkillCode)
		}

		columnKeyNumber, err := file.GetCellValue("Sheet1", keyCodeColumn+strconv.Itoa(rowCount))
		if err != nil {
			log.Fatal(err)
		}

		columnKeyDescription, err := file.GetCellValue("Sheet1", keyDescriptionColumn+strconv.Itoa(rowCount))
		if err != nil {
			log.Fatal(err)
		}

		columnSFIALevel, err := file.GetCellValue("Sheet1", sfiaColumn+strconv.Itoa(rowCount))
		if err != nil {
			log.Fatal(err)
		}

		var foundSkill SkillDataModel
		foundSkill.SkillCode = columnSkillCode
		foundSkill.SFIALevel = columnSFIALevel
		foundSkill.KeyPointNumber = columnKeyNumber
		foundSkill.KeyPointDescription = columnKeyDescription
		detailedSkills = append(detailedSkills, foundSkill)

		rowCount++
	}

	var skillResponse SkillResponse
	skillResponse.Skills = skills
	skillResponse.DetailedSkills = detailedSkills

	return skillResponse

}

func generatePDP(skillModel SkillResponse, templateLocation string) string {

	file, err := ioutil.ReadFile(templateLocation)
	if err != nil {
		fmt.Printf("Could not read template file due to this %s error \n", err)
	}

	// convert the file binary into a string using string
	fileContent := string(file)

	// Create file content

	// SFIA LEVEL
	var sfiaLevel = skillModel.DetailedSkills[0].SFIALevel
	fileContent = strings.Replace(fileContent, "@SFIA@", sfiaLevel, -1)

	// SKILLS
	var prefixedSkills []string
	for skillIndex := 0; skillIndex < len(skillModel.Skills); skillIndex++ {
		prefixedSkills = append(prefixedSkills, "* "+skillModel.Skills[skillIndex])
	}

	var skillsInsert = strings.Join(prefixedSkills, "\n")
	fileContent = strings.Replace(fileContent, "@SKILLS@", skillsInsert, -1)

	// SKILL CHECKLISTS
	var skillChecklists string
	for skillIndex := 0; skillIndex < len(skillModel.Skills); skillIndex++ {

		var checkingSkill = skillModel.Skills[skillIndex]

		skillChecklists += "  \n  \n"
		skillChecklists += "### Skill Group: " + checkingSkill + "  \n  \n"
		skillChecklists += "| ID  | Description  | Date  | Signed off By  |  \n"
		skillChecklists += "|---|---|---|---|  \n"

		for detailedSkillIndex := 0; detailedSkillIndex < len(skillModel.DetailedSkills); detailedSkillIndex++ {
			var checkingDetailedSkill = skillModel.DetailedSkills[detailedSkillIndex]
			if checkingDetailedSkill.SkillCode == checkingSkill {

				skillChecklists += "| " + checkingDetailedSkill.KeyPointNumber + " | " + checkingDetailedSkill.KeyPointDescription + " | | |  \n"
				skillChecklists += "| Evidence: <td colspan=3> Insert evidence and reference here... |  \n"
			}
		}

		skillChecklists += "  \n  \n"
	}

	fileContent = strings.Replace(fileContent, "@CRITERIA@", skillChecklists, -1)
	return fileContent
}

func savePDP(config Config, pdpContents string, outputFilename string) {

	f, err := os.Create(config.ProcessedOutputFolder + outputFilename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(pdpContents)
	if err2 != nil {
		log.Fatal(err2)
	}
}

// func exportAsJSON(skillModels []SkillDataModel, outputLocation string, currentTimestamp int64) {

// 	jsonContent, err := json.Marshal(skillModels)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var jsonOutputLocation string = outputLocation + strconv.Itoa(int(currentTimestamp)) + "-SFIA-Skill-PDP.json"
// 	err = ioutil.WriteFile(jsonOutputLocation, jsonContent, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// func exportAsExcel(skillModels []SkillDataModel, outputLocation string, currentTimestamp int64) {
// 	file := excelize.NewFile()

// 	file.SetCellValue("Sheet1", "A1", "SkillCode")
// 	file.SetCellValue("Sheet1", "B1", "SFIA Level")
// 	file.SetCellValue("Sheet1", "C1", "Keycode")
// 	file.SetCellValue("Sheet1", "D1", "Description")

// 	for skillItem := 0; skillItem < len(skillModels); skillItem++ {
// 		var insertLocation int = skillItem + 2
// 		file.SetCellValue("Sheet1", "A"+strconv.Itoa(insertLocation), skillModels[skillItem].SkillCode)
// 		file.SetCellValue("Sheet1", "B"+strconv.Itoa(insertLocation), skillModels[skillItem].SFIALevel)
// 		file.SetCellValue("Sheet1", "C"+strconv.Itoa(insertLocation), skillModels[skillItem].KeyPointNumber)
// 		file.SetCellValue("Sheet1", "D"+strconv.Itoa(insertLocation), skillModels[skillItem].KeyPointDescription)
// 	}

// 	var jsonOutputLocation string = outputLocation + strconv.Itoa(int(currentTimestamp)) + "-SFIA-Skill-PDP.xlsx"
// 	if err := file.SaveAs(jsonOutputLocation); err != nil {
// 		log.Fatal(err)
// 	}
// }
