package config

import (
	"os"
	"path/filepath"
	"tessa/internal/model"
)

const (
	NAME = "tessa"
)

func getE() *model.Err {
	e := &model.Err{}
	return e
}

func getHomeDir() string {
	path, err := os.UserHomeDir()
	if err != nil {
		getE().ShowWarning(err.Error())
	}
	return path
}

func IsConfigured() bool {
	return false
}

func createFiles(_ string) bool {
	return false
}

func createDir() string {
	path := getHomeDir()
	dirName := "."+NAME
	newPath := filepath.Join(path, dirName)

	if err := os.MkdirAll(newPath, 0700); err != nil {
		getE().ShowWarning(err.Error())
	}

	return newPath;
}

func Configure() {
	path := createDir()
	res := createFiles(path)
	if !res {
		getE().ShowError("failed to configure")
	}
}