package main

import (
	"fmt"
	"os"
	"strings"
	"tessa/internal/config"
	"tessa/internal/db"
	"tessa/internal/model"

	"github.com/fatih/color"
)

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

	args := os.Args;
	if len(args) == 1 {
		e.ShowWarning("not enough arguments")
	}

	command := args[1];
	switch command {
		case "save":
			data := ParseArgs(args[2:])
			if err := db.SaveData(data); err != nil {
				fmt.Printf("%v\n", err);
				e.ShowError("failed to save data")
			}
			e.ShowRes("cool")
		case "get":
			if len(args) >= 3 {
				e.ShowInfo("get requires no args")
			}
			storedData, err := db.GetData()
			if err != nil {
				e.ShowWarning("something is wrong lets debug!")
				fmt.Println(err.Error())
			}
			for _, indiData := range storedData {
				fmt.Printf("%d. %s\n", indiData.Id, indiData.Data)
			}
		case "clean":
			if err := db.CleanData(); err != nil {
				e.ShowWarning(err.Error())
			}
			e.ShowRes("cleared")
		case "whoisdeno":
			fmt.Printf("\n")
			fmt.Println("-------------- DENO --------------")
			color.Cyan(">> https://www.github.com/zeni-42")
			fmt.Printf("\n")
		default:
			newMsg := "command not found: " + args[1];
			e.ShowWarning(newMsg);
	}
}
