package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"tessa/internal/config"
	"tessa/internal/db"
	"tessa/internal/model"

	"github.com/atotto/clipboard"
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
			if strings.TrimSpace(data) == "" {
				e.ShowWarning("message cannot be empty")
				return
			}
			if err := db.SaveData(data); err != nil {
				fmt.Printf("%v\n", err);
				e.ShowError("failed to save data")
			}
			e.ShowRes("cool")
		case "show":
			if len(args) >= 3 {
				e.ShowInfo("show requires no args")
			}
			storedData, err := db.GetData()
			if len(storedData) == 0 {
				e.ShowWarning("no previous data")
			}

			if err != nil {
				e.ShowWarning("something is wrong lets debug!")
				fmt.Println(err.Error())
			}
			for _, indiData := range storedData {
				fmt.Printf("%d) %s\n", indiData.Id, indiData.Data)
			}
		case "get":
			if len(args) != 3 {
				e.ShowWarning("expecting two args, got more")
			}
			strNo, err := strconv.Atoi(args[2])
			if err != nil {
				e.ShowError("invalid args")
			}

			res, err := db.GetDataById(strNo)
			if err != nil {
				e.ShowError("failed to retieve data")
			}

			if err := clipboard.WriteAll(res.Data); err != nil {
				e.ShowError("failed to save data")
			}

			e.ShowRes("copied")
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
