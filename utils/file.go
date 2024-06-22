package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"projectCrawler/models"
)

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return fmt.Errorf("error while creating %s directory: %v", dir, err)
		}
	}
	return nil
}

func Contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func ReadIgnoreFile() (models.IgnoreData, error) {
	var data models.IgnoreData

	file, err := os.Open("./crawlerIgnore.json")
	if err != nil {
		if os.IsNotExist(err) {
			return models.IgnoreData{Dirs: []string{}, Files: []string{}}, nil
		}
		return data, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}
