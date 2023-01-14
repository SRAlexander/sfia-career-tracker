package main

import (
	"fmt"
	"os"

	// "strings"
	// "time"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	pdpJobSpecsLocation     string `json:"pdpJobSpecs"`
	pdpTitleColumn          string `json:"pdpTitleColumn "`
	pdpCoreSkillsColumn            string `json:"sfiaColumn"`
	pdpSpecialismSkillsColumn            string `json:"sfiaColumn"`
}

func main() {

	var config Config
	json.Unmarshal(loadJSONFileAsByteString("config.json"), &config)

	var loadPDPRoleCritieria(config.)

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