package model

import (
	"reflect"
	"testing"

	"github.com/sysu-go-online/public-service/types"
)

func TestUpdateFileContent(t *testing.T) {
	type args struct {
		projectName string
		username    string
		filePath    string
		content     string
		create      bool
		dir         bool
		projectType int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateFileContent(tt.args.projectName, tt.args.username, tt.args.filePath, tt.args.content, tt.args.create, tt.args.dir, tt.args.projectType); (err != nil) != tt.wantErr {
				t.Errorf("UpdateFileContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteFile(t *testing.T) {
	type args struct {
		projectName string
		username    string
		filePath    string
		projectType int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteFile(tt.args.projectName, tt.args.username, tt.args.filePath, tt.args.projectType); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFileContent(t *testing.T) {
	type args struct {
		projectName string
		username    string
		filePath    string
		projectType int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFileContent(tt.args.projectName, tt.args.username, tt.args.filePath, tt.args.projectType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFileContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFileStructure(t *testing.T) {
	type args struct {
		projectName string
		username    string
		projectType int
	}
	tests := []struct {
		name    string
		args    args
		want    *types.FileStructure
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFileStructure(tt.args.projectName, tt.args.username, tt.args.projectType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileStructure() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFileStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRenameFile(t *testing.T) {
	type args struct {
		projectName string
		username    string
		rawPathName string
		afterName   string
		projectType int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RenameFile(tt.args.projectName, tt.args.username, tt.args.rawPathName, tt.args.afterName, tt.args.projectType); (err != nil) != tt.wantErr {
				t.Errorf("RenameFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getFilePath(t *testing.T) {
	type args struct {
		username    string
		projectName string
		filePath    string
		projectType int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFilePath(tt.args.username, tt.args.projectName, tt.args.filePath, tt.args.projectType); got != tt.want {
				t.Errorf("getFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
