package main

import (
	"os"
    "fmt"
    "log"
	"io/ioutil"
	"strconv"
	"encoding/json"
    "github.com/xuri/excelize/v2"
	"os/exec"
	"bytes"
)

type Config struct {
    SFIAJobSpecLocation string `json:"sfiaJobSpecLocation"`
	SFIAJobSpecSheetName string `json:"sfiaJobSpecSheetName"`
    JobTitleColumn string `json:"jobTitleColumn"`
	FilenameColumn string `json:"filenameColumn"`
	SFIALevelColumn string `json:"sfiaLevelColumn"`
	CoreSkillsColumn string `json:"coreSkillsColumn"`
	SpecialismSkillsColumn string `json:"specialismSkillsColumn"`
}

type JobRole struct {
	JobTitle string
	Filename string
	SFIALevel string
	CoreSkills string
	SpecialismSkills string
}

func main() {

	var config Config
	json.Unmarshal(loadJSONFileAsByteString("config.json"), &config)

    file, err := excelize.OpenFile(config.SFIAJobSpecLocation)
    if err != nil {
        log.Fatal(err)
    }

	var jobRoles []JobRole
	jobRoles = append(jobRoles, processJobRoles(file, config)...)

	for jobRoleIndex := 0; jobRoleIndex < len(jobRoles); jobRoleIndex++ {
		runPDPCriteriaGenerator(jobRoles[jobRoleIndex])
		runPDPGenerator(jobRoles[jobRoleIndex])
	}

}

// The behaviours sheet has a different format to skills
func processJobRoles(file *excelize.File, config Config) []JobRole {

	var rowCount int = 2;
	var jobRoles []JobRole

	for {
		columnJobTitle, err := file.GetCellValue(config.SFIAJobSpecSheetName, config.JobTitleColumn + strconv.Itoa(rowCount))
		if err != nil {
			log.Fatal(err)
		}

		if (columnJobTitle == "") {
			break;
		}

		columnFilename, err := file.GetCellValue(config.SFIAJobSpecSheetName, config.FilenameColumn + strconv.Itoa(rowCount))
		if err != nil {
			log.Fatal(err)
		}

		columnSFIALevel, err := file.GetCellValue(config.SFIAJobSpecSheetName, config.SFIALevelColumn + strconv.Itoa(rowCount))
		if err != nil {
			log.Fatal(err)
		}

		columnCoreSkills, err := file.GetCellValue(config.SFIAJobSpecSheetName, config.CoreSkillsColumn + strconv.Itoa(rowCount))
		if err != nil {
			log.Fatal(err)
		}

		columnSpecialismSkills, err := file.GetCellValue(config.SFIAJobSpecSheetName, config.SpecialismSkillsColumn + strconv.Itoa(rowCount))
		if err != nil {
			log.Fatal(err)
		}

		var jobRole JobRole
		jobRole.JobTitle = columnJobTitle
		jobRole.Filename = columnFilename
		jobRole.SFIALevel = columnSFIALevel
		jobRole.CoreSkills = columnCoreSkills
		jobRole.SpecialismSkills = columnSpecialismSkills
		jobRoles = append(jobRoles, jobRole)
		
		rowCount++
	}

	return jobRoles
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

func runPDPCriteriaGenerator(jobRole JobRole) {
	command := exec.Command("go", "run", "create-pdp-criteria.go", "--sfia-level", jobRole.SFIALevel, "--output-filename", jobRole.Filename +".xlsx", "--core-skills", jobRole.CoreSkills, "--specialism-skills", jobRole.SpecialismSkills)
	fmt.Println(command)
	command.Dir = "../02-app-pdp-criteria-generator"

	var out bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stderr

	err := command.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	
}

func runPDPGenerator(jobRole JobRole) {
	command := exec.Command("go", "run", "pdp-generator.go", "--skill-list", "../outputs/app-pdp-criteria-generator/" + jobRole.Filename +".xlsx" , "--output-filename", jobRole.Filename +".MD")
	fmt.Println(command)
	command.Dir = "../03-app-pdp-generator"

	var out bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stderr

	err := command.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
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



