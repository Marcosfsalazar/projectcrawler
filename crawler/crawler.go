package crawler

import (
	"io/ioutil"
	"path/filepath"
	"projectCrawler/models"
	"projectCrawler/utils"
)

func ReadDir(root string, dirsToExclude, filesToExclude []string) (models.Directory, error) {
	var result models.Directory
	result.SubDirs = make(map[string]models.Directory)
	result.Files = make(map[string]models.FileContent)

	entries, err := ioutil.ReadDir(root)
	if err != nil {
		return result, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if utils.Contains(dirsToExclude, entry.Name()) {
				continue
			}
			subDirPath := filepath.Join(root, entry.Name())
			subDir, err := ReadDir(subDirPath, dirsToExclude, filesToExclude)
			if err != nil {
				return result, err
			}
			result.SubDirs[entry.Name()] = subDir
		} else {
			if utils.Contains(filesToExclude, entry.Name()) {
				continue
			}
			filePath := filepath.Join(root, entry.Name())
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				return result, err
			}
			result.Files[entry.Name()] = models.FileContent{Content: string(content)}
		}
	}
	return result, nil
}
