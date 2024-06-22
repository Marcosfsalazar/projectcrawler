package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"projectCrawler/crawler"
	"projectCrawler/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: projectCrawler <root-directory>")
		return
	}

	rootDir := os.Args[1]

	data, err := utils.ReadIgnoreFile()
	if err != nil {
		fmt.Println("error reading ignore file!", err)
		return
	}

	dirsToExclude := data.Dirs
	filesToExclude := data.Files

	projectStructure, err := crawler.ReadDir(rootDir, dirsToExclude, filesToExclude)
	if err != nil {
		fmt.Println("Error reading directory: ", err)
		return
	}

	jsonData, err := json.MarshalIndent(projectStructure, "", " ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	dir := "crawledContent"
	if err := utils.CreateDirIfNotExist(dir); err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile("./crawledContent/project_structure.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("Success!")
}
