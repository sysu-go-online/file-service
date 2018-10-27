package model

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sysu-go-online/public-service/tools"
	"github.com/sysu-go-online/public-service/types"
)

// ROOT defines the root directory
var ROOT = "/home"

// UpdateFileContent update content with given filepath and content
func UpdateFileContent(username string, projectPath string, filePath string, content string, create bool, dir bool) error {
	// Get absolute path
	var err error
	absPath := getFilePath(username, filePath, projectPath)

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
func DeleteFile(username string, projectPath string, filePath string) error {
	// Get absolute path
	var err error
	absPath := getFilePath(username, filePath, projectPath)

	// Delete file
	err = os.RemoveAll(absPath)
	return err
}

// GetFileContent returns required file content
func GetFileContent(username string, projectPath string, filePath string) ([]byte, error) {
	// Get absolute path
	var err error
	absPath := getFilePath(username, filePath, projectPath)

	// Read file content
	f, err := os.Stat(absPath)
	if err != nil {
		return nil, err
	}
	if f.Size() > 5*1024*1024 {
		return nil, errors.New("File is too Large")
	}
	content, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// GetFileStructure read file structure and return it
func GetFileStructure(username string, projectPath string, projectName string) (*types.FileStructure, error) {
	// Get absolute path
	var err error
	absPath := getFilePath(username, "/", projectPath)

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
func RenameFile(username string, projectPath string, rawPathName, afterName string) error {
	// Get absolute path
	absPath := getFilePath(username, rawPathName, projectPath)
	newPath := getFilePath(username, afterName, projectPath)
	return os.Rename(absPath, newPath)
}

func getFilePath(username string, filePath string, projectPath string) string {
	return filepath.Join("home", username, "projects", projectPath, filePath)
}
