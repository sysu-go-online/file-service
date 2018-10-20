package model

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sysu-go-online/public-service/tools"
	"github.com/sysu-go-online/public-service/types"
)

// ROOT defines the root directory
var ROOT = "/home"

// UpdateFileContent update content with given filepath and content
func UpdateFileContent(projectName string, username string, filePath string, content string, create bool, dir bool, projectType int) error {
	// Get absolute path
	var err error
	absPath := getFilePath(username, projectName, filePath, projectType)

	// Update file, if the file not exists, judge accroding to the given param
	if create {
		if dir {
			err = os.Mkdir(absPath, os.ModeDir)
		} else {
			f, err := os.Create(absPath)
			if err != nil {
				f.Close()
			}
		}
	} else {
		err = ioutil.WriteFile(absPath, []byte(content), os.ModeAppend)
	}
	return err
}

// DeleteFile delete file accroding to the given path
func DeleteFile(projectName string, username string, filePath string, projectType int) error {
	// Get absolute path
	var err error
	absPath := getFilePath(username, projectName, filePath, projectType)

	// Delete file
	err = os.RemoveAll(absPath)
	return err
}

// GetFileContent returns required file content
func GetFileContent(projectName string, username string, filePath string, projectType int) ([]byte, error) {
	// Get absolute path
	var err error
	absPath := getFilePath(username, projectName, filePath, projectType)

	// Read file content
	content, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// GetFileStructure read file structure and return it
func GetFileStructure(projectName string, username string, projectType int) (*types.FileStructure, error) {
	// Get absolute path
	var err error
	absPath := getFilePath(username, projectName, "/", projectType)

	// Recurisively get file structure
	s, err := tools.Dfs(absPath, 0)
	if err != nil {
		return nil, err
	}
	// Add root content
	root := types.FileStructure{
		Name:       projectName,
		Type:       "dir",
		Children:   s,
		Root:       true,
		IsSelected: true,
	}
	return &root, nil
}

// RenameFile rename file
func RenameFile(projectName string, username string, rawPathName, afterName string, projectType int) error {
	// Get absolute path
	absPath := getFilePath(username, projectName, rawPathName, projectType)
	newPath := getFilePath(username, projectName, afterName, projectType)

	return os.Rename(absPath, newPath)
}

func getFilePath(username string, projectName string, filePath string, projectType int) string {
	switch projectType {
	case 0:
		// golang
		return filepath.Join(ROOT, username, "go/src/github.com", projectName, filePath)
	case 1:
		// cpp
		return filepath.Join(ROOT, username, "cpp", projectName, filePath)
	}
	return ""
}
