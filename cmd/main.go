package main

import (
	"tessa/internal/db"
	"tessa/internal/model"
	"tessa/internal/config"
)

func getE() *model.Err {
	return &model.Err{}
}

func main() {
	if !config.IsConfigured() {
		config.Configure()
	}

	db.GetCon()
	defer db.CloseCon()

	if err := db.Init(); err != nil {
		getE().ShowError("database init failed")
	}
}
