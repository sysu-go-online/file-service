package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sysu-go-online/file-service/model"

	projectModel "github.com/sysu-go-online/project-service/model"
	"github.com/sysu-go-online/public-service/tools"
	userModel "github.com/sysu-go-online/user-service/model"
)

// UpdateFileHandler is a handler for update file or rename
func UpdateFileHandler(w http.ResponseWriter, r *http.Request) error {
	// Read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	type temp struct {
		Operation string `json:"operation"`
		Content   string `json:"content"`
	}

	// get operation from body
	op := &temp{}
	err = json.Unmarshal(body, op)
	if err != nil {
		w.WriteHeader(400)
		return nil
	}

	// get username from jwt
	ok, username := tools.GetUserNameFromToken(r.Header.Get("Authorization"), AuthRedisClient)
	if !ok {
		w.WriteHeader(401)
		return nil
	}
	// Read project name and file path from uri
	vars := mux.Vars(r)
	projectName := vars["projectname"]
	filePath := vars["filepath"]

	// Get project information
	session := MysqlEngine.NewSession()
	u := userModel.User{Username: username}
	ok, err = u.GetWithUsername(session)
	if !ok {
		w.WriteHeader(400)
		return nil
	}
	if err != nil {
		return err
	}
	p := projectModel.Project{Name: projectName, UserID: u.ID}
	has, err := p.GetWithUserIDAndName(session)
	if !has {
		w.WriteHeader(204)
		return nil
	}
	if err != nil {
		return err
	}

	// Check if the file path is valid

	switch op.Operation {
	case "update":
		err = model.UpdateFileContent(username, p.Path, projectName, filePath, op.Content, false, false)
	case "rename":
		err = model.RenameFile(username, p.Path, projectName, filePath, op.Content)
	default:
		w.WriteHeader(400)
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

// CreateFileHandler is a handler for create file
func CreateFileHandler(w http.ResponseWriter, r *http.Request) error {
	// Read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	type IsDir struct {
		Dir bool `json:"dir"`
	}
	// Judge if it is dir from body
	isDir := IsDir{}
	err = json.Unmarshal(body, &isDir)
	if err != nil {
		return err
	}
	dir := isDir.Dir

	// get username from jwt
	ok, username := tools.GetUserNameFromToken(r.Header.Get("Authorization"), AuthRedisClient)
	if !ok {
		w.WriteHeader(401)
		return nil
	}
	// Read project id and file path from uri
	vars := mux.Vars(r)
	projectName := vars["projectname"]
	filePath := vars["filepath"]

	// Get project information
	session := MysqlEngine.NewSession()
	u := userModel.User{Username: username}
	ok, err = u.GetWithUsername(session)
	if !ok {
		w.WriteHeader(400)
		return nil
	}
	if err != nil {
		return err
	}
	p := projectModel.Project{Name: projectName, UserID: u.ID}
	has, err := p.GetWithUserIDAndName(session)
	if !has {
		w.WriteHeader(204)
		return nil
	}
	if err != nil {
		return err
	}

	if ok {
		// Save file
		err := model.UpdateFileContent(username, p.Path, projectName, filePath, "", true, dir)
		if err != nil {
			return err
		}
		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
	}
	return nil
}

// GetFileContentHandler is a handler for reading file content
func GetFileContentHandler(w http.ResponseWriter, r *http.Request) error {
	// get username from jwt
	ok, username := tools.GetUserNameFromToken(r.Header.Get("Authorization"), AuthRedisClient)
	if !ok {
		w.WriteHeader(401)
		return nil
	}

	// Read project id and file path from uri
	vars := mux.Vars(r)
	projectName := vars["projectname"]
	filePath := vars["filepath"]

	// Get project information
	session := MysqlEngine.NewSession()
	u := userModel.User{Username: username}
	ok, err := u.GetWithUsername(session)
	if !ok {
		w.WriteHeader(400)
		return nil
	}
	if err != nil {
		return err
	}
	p := projectModel.Project{Name: projectName, UserID: u.ID}
	has, err := p.GetWithUserIDAndName(session)
	if !has {
		w.WriteHeader(204)
		return nil
	}
	if err != nil {
		return err
	}

	// Load file
	content, err := model.GetFileContent(username, p.Path, projectName, filePath)
	if err != nil {
		return err
	}
	w.WriteHeader(200)
	w.Write(content)
	return nil
}

// DeleteFileHandler is a handler for delete file
func DeleteFileHandler(w http.ResponseWriter, r *http.Request) error {
	// get username from jwt
	ok, username := tools.GetUserNameFromToken(r.Header.Get("Authorization"), AuthRedisClient)
	if !ok {
		w.WriteHeader(401)
		return nil
	}
	// Read project id and file path from uri
	vars := mux.Vars(r)
	projectName := vars["projectname"]
	filePath := vars["filepath"]

	// Get project information
	session := MysqlEngine.NewSession()
	u := userModel.User{Username: username}
	ok, err := u.GetWithUsername(session)
	if !ok {
		w.WriteHeader(400)
		return nil
	}
	if err != nil {
		return err
	}
	p := projectModel.Project{Name: projectName, UserID: u.ID}
	has, err := p.GetWithUserIDAndName(session)
	if !has {
		w.WriteHeader(204)
		return nil
	}
	if err != nil {
		return err
	}

	// Load file
	err = model.DeleteFile(username, p.Path, projectName, filePath)
	if err != nil {
		return err
	}
	w.WriteHeader(200)
	return nil
}

// GetFileStructureHandler is handler for get project structure
func GetFileStructureHandler(w http.ResponseWriter, r *http.Request) error {
	// get username from jwt
	ok, username := tools.GetUserNameFromToken(r.Header.Get("Authorization"), AuthRedisClient)
	if !ok {
		w.WriteHeader(401)
		return nil
	}
	// Read project id
	vars := mux.Vars(r)
	projectName := vars["projectname"]

	// Get project information
	session := MysqlEngine.NewSession()
	u := userModel.User{Username: username}
	ok, err := u.GetWithUsername(session)
	if !ok {
		w.WriteHeader(400)
		return nil
	}
	if err != nil {
		return err
	}
	p := projectModel.Project{Name: projectName, UserID: u.ID}
	has, err := p.GetWithUserIDAndName(session)
	if !has {
		w.WriteHeader(204)
		return nil
	}
	if err != nil {
		return err
	}

	// Get file structure
	structure, err := model.GetFileStructure(username, p.Path, projectName)
	if err != nil {
		return err
	}
	ret, err := json.Marshal(structure)
	if err != nil {
		return err
	}
	w.Write(ret)
	return nil
}
