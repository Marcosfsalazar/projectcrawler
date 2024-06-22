package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileContent struct {
	Content string `json:"content"`
}

type 	Directory struct {
	SubDirs map[string]Directory `json:"sub_dirs,omitempty"`
	Files map[string]FileContent `json:"files,omitempty"`
}

func main(){
	if len(os.Args) < 2{
		fmt.Println("Usage: projectCrawler <root-directory>")
		return
	}

	rootDir := os.Args[1]
	
	data,err := readIgnoreFile()

	if err !=nil {
		fmt.Println("error reading ignore file!", err)
		return
	}

	dirsToExclude := data.Dirs
	filesToExclude := data.Files

	projectStructure, err := readDir(rootDir, dirsToExclude, filesToExclude)
	if err != nil{
		fmt.Println("Error reading directory: ", err)
		return
	}

	jsonData, err := json.MarshalIndent(projectStructure, "", " ")
	if err != nil{
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	dir := "crawledContent"
	if err := createDirIfNotExist(dir); err != nil{
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile("./crawledContent/project_structure.json", jsonData, 0644)

	if err != nil{
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("Success!")
}

func readDir(root string, dirsToExclude, filesToExclude []string) (Directory, error){
	var result Directory
	result.SubDirs = make(map[string]Directory)
	result.Files = make(map[string]FileContent)

	entries, err := ioutil.ReadDir(root)
	
	if err != nil {
		return result, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if contains(dirsToExclude, entry.Name()) {
					continue
			}
			subDirPath := filepath.Join(root, entry.Name())
			subDir, err := readDir(subDirPath, dirsToExclude, filesToExclude)
			if err != nil{
				return result, err
			}
			result.SubDirs[entry.Name()] = subDir
		}else{
			if contains(filesToExclude, entry.Name()){
				continue
			}
			filePath := filepath.Join(root, entry.Name())
			content, err := ioutil.ReadFile(filePath)
			
			if err != nil{
				return result, err
			}

			result.Files[entry.Name()] = FileContent{Content: string(content)}
		}
	}
	return result, nil
}

func contains(slice []string, item string) bool {
	for _, a := range slice{
		if a == item {
			return true
		}
	}
	return false
}

func createDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return fmt.Errorf("error while creating %s directory: %v", dir, err)
		}
	}
	return nil
}

type IgnoreData struct{
	Dirs []string `json:"dirs"`
	Files []string `json:"files"`
}

func readIgnoreFile() (IgnoreData,error){
	var data IgnoreData

	file, err := os.Open("./crawlerIgnore.json")

	if err != nil {
		if os.IsNotExist(err) {
			return IgnoreData{Dirs: []string{}, Files: []string{}}, nil
		}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&data); err != nil{
		return data, err
	}

	return data, err
}