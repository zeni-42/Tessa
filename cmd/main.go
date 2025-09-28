package main

import (
	"fmt"
	"os"
	"strings"
	"tessa/internal/config"
	"tessa/internal/db"
	"tessa/internal/model"
)

func ParseCommands(s string) string {
	switch s {
		case "-s":
			return "save"
		case "-g":
			return "get"
		default:
			return ""
	}
}

func ParseArgs(s []string) string {
	return strings.Join(s, " ")
}

func main() {
	e := &model.Err{}

	if !config.IsConfigured() {
		config.Configure()
	}

	db.GetCon()
	defer db.CloseCon()

	if err := db.Init(); err != nil {
		e.ShowError("database init failed")
	}

	command := ParseCommands(os.Args[1]);
	if command == "" {
		newMsg := "command not found: " + os.Args[1];
		e.ShowWarning(newMsg);
	}
	
	pureString := ParseArgs(os.Args[2:])

	fmt.Printf("%s\n", pureString)
}
