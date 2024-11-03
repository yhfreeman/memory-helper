package main

import (
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v3"
)

// YAML structure for persisting data
type YAMLData struct {
    StudyItems []StudyItem `yaml:"study_items"`
    Events     []Event     `yaml:"events"`
}

// LoadEvents reads events from a YAML file.
func LoadDB(filePath string) (YAMLData, error) {
	var storage YAMLData

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return storage, err
	}

	err = yaml.Unmarshal(data, &storage)
	return storage, err
}

// SaveEvents writes events to a YAML file.
func WriteDB(filePath string, storage YAMLData) error {
	data, err := yaml.Marshal(&storage)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, data, os.ModePerm)
}
