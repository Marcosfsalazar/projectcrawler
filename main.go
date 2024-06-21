package main

import (
	"io/ioutil"
	"path/filepath"
)

type FileContent struct {
	Content string `json:"content"`
}

type 	Directory struct {
	SubDirs map[string]Directory `json:"sub_dirs,omitempty"`
	Files map[string]FileContent `json:"files,omitempty"`
}

func main(){}

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