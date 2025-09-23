package config

import (
	"fmt"
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

// A public function which checks if the application is already configured or not
func IsConfigured() bool {
	path := getHomeDir()
	dirName := "."+NAME
	configFile := filepath.Join(path, dirName, "config.jsonc");
	databaseFile := filepath.Join(path, dirName, "database.db");
	
	if _, err := os.Stat(configFile); err != nil {
		fmt.Println("configuring..")
		return false
	}
	
	if _, err := os.Stat(databaseFile); err != nil {
		fmt.Println("configuring..");
		return false
	}
	
	return true
}

// Create the required files like configuration and database file
func createFiles(path string) bool {
	configFile := filepath.Join(path, "config.jsonc")
	databaseFile := filepath.Join(path, "database.db")

	if _, err := os.Create(configFile); err != nil {
		getE().ShowWarning("failed to create configuration")
		return false
	}
	if _, err := os.Create(databaseFile); err != nil {
		getE().ShowWarning("failed to create database")
		return false
	}
	return true
}

// A private function which create the required directories for the application
func createDir() string {
	path := getHomeDir()
	dirName := "."+NAME
	newPath := filepath.Join(path, dirName)

	if err := os.MkdirAll(newPath, 0700); err != nil {
		getE().ShowWarning(err.Error())
	}

	return newPath;
}

// A public function which internally calls the required fuctions to configure the application for the first time
func Configure() {
	path := createDir()
	res := createFiles(path)
	if !res {
		getE().ShowError("failed to configure")
	}
}