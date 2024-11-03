package main

import (
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v3"
)

// EventStorage represents the structure of the YAML file.
type EventStorage struct {
	Events []Event `yaml:"events"`
}

// LoadEvents reads events from a YAML file.
func LoadEvents(filePath string) (EventStorage, error) {
	var storage EventStorage

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return storage, err
	}

	err = yaml.Unmarshal(data, &storage)
	return storage, err
}

// SaveEvents writes events to a YAML file.
func SaveEvents(filePath string, storage EventStorage) error {
	data, err := yaml.Marshal(&storage)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, data, os.ModePerm)
}
