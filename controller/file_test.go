package controller

import (
	"net/http"
	"testing"
)

func TestUpdateFileHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
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
			if err := UpdateFileHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("UpdateFileHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateFileHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
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
			if err := CreateFileHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("CreateFileHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFileContentHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
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
			if err := GetFileContentHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("GetFileContentHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteFileHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
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
			if err := DeleteFileHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFileHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFileStructureHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
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
			if err := GetFileStructureHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("GetFileStructureHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
