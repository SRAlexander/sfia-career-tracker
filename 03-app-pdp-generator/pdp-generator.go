package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Config struct {
	TemplateLocation      string `json:"templateLocation"`
	ProcessedOutputFolder string `json:"processedOutputFolder"`
	SkillColumn           string `json:"skillColumn"`
	SkillTitleColumn      string `json:"skillTitleColumn"`
	SFIAColumn            string `json:"sfiaColumn"`
	KeyCodeColumn         string `json:"keyColumn"`
	KeyDescriptionColumn  string `json:"keyDescriptionColumn"`
}

type SkillDataModel struct {
	SkillCode           string
	SkillTitle          string
	SFIALevel           string
	KeyPointNumber      string
	KeyPointDescription string
}

type SkillMap struct {
	Code  string
	Title string
}

type SkillResponse struct {
	Skills         []SkillMap
	Specialisms    []SkillMap
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

	fmt.Println("Created " + config.ProcessedOutputFolder + outputFilename)
}

func loadJSONFileAsByteString(file string) []byte {

	// Open our jsonFile
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

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

func skillMapContains(s []SkillMap, value string) bool {
	for _, v := range s {
		if v.Code == value {
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
	var skillTitleColumn string = config.SkillTitleColumn
	var sfiaColumn string = config.SFIAColumn
	var keyCodeColumn string = config.KeyCodeColumn
	var keyDescriptionColumn string = config.KeyDescriptionColumn

	var rowCount int = 2

	var skills []SkillMap
	var specialisms []SkillMap
	var detailedSkills []SkillDataModel
	var blankLinesFound int = 0

	for {

		columnSkillCode, err := file.GetCellValue("Sheet1", skillColumn+strconv.Itoa(rowCount))
		if err != nil {
			log.Fatal(err)
		}

		if columnSkillCode == "" {
			blankLinesFound++
			if blankLinesFound > 1 {
				break
			}
		} else {

			columnSkillTitle, err := file.GetCellValue("Sheet1", skillTitleColumn+strconv.Itoa(rowCount))
			if err != nil {
				log.Fatal(err)
			}

			if !skillMapContains(skills, columnSkillCode) {

				var skillMap SkillMap
				skillMap.Code = columnSkillCode
				skillMap.Title = columnSkillTitle
				if blankLinesFound > 0 {
					skills = append(skills, skillMap)
					specialisms = append(specialisms, skillMap)
				} else {
					skills = append(skills, skillMap)
				}
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
			foundSkill.SkillTitle = columnSkillTitle
			foundSkill.SFIALevel = columnSFIALevel
			foundSkill.KeyPointNumber = columnKeyNumber
			foundSkill.KeyPointDescription = columnKeyDescription
			detailedSkills = append(detailedSkills, foundSkill)

		}
		rowCount++

	}

	var skillResponse SkillResponse
	skillResponse.Skills = skills
	skillResponse.Specialisms = specialisms
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
		var skillTitle = skillModel.Skills[skillIndex]
		var displayTitle = "Unknown Skill"
		if skillMapContains(skillModel.Specialisms, skillTitle.Code) {
			displayTitle = skillTitle.Title + " | " + skillTitle.Code + " (Specialism)"
		} else {
			displayTitle = skillTitle.Title + " | " + skillTitle.Code
		}
		prefixedSkills = append(prefixedSkills, "* "+displayTitle)
	}

	var skillsInsert = strings.Join(prefixedSkills, "\n")
	fileContent = strings.Replace(fileContent, "@SKILLS@", skillsInsert, -1)

	// SKILL CHECKLISTS
	var skillChecklists string
	for skillIndex := 0; skillIndex < len(skillModel.Skills); skillIndex++ {

		var checkingSkill = skillModel.Skills[skillIndex]
		var skillTitle = checkingSkill.Title
		if skillMapContains(skillModel.Specialisms, checkingSkill.Code) {
			skillTitle = skillTitle + " (Specialism)"
		}

		skillChecklists += "  \n  \n"
		skillChecklists += "### Skill Group: " + skillTitle + "  \n  \n"
		skillChecklists += "| ID  | Description  | Date  | Signed off By  |  \n"
		skillChecklists += "|---|---|---|---|  \n"

		for detailedSkillIndex := 0; detailedSkillIndex < len(skillModel.DetailedSkills); detailedSkillIndex++ {
			var checkingDetailedSkill = skillModel.DetailedSkills[detailedSkillIndex]
			if checkingDetailedSkill.SkillCode == checkingSkill.Code {

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
