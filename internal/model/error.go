package model

import (
	"os"

	"github.com/fatih/color"
)

type Err struct {
	message string
}

func (e *Err) ShowInfo(message string) {
	updatedMsg := "INFO: " + message
	e.message = updatedMsg
	color.Blue(e.message)
}

func (e *Err) ShowRes(message string) {
	updatedMsg := ">> " + message
	e.message = updatedMsg
	color.Green(e.message)
	os.Exit(0)
}

func (e *Err) ShowWarning(message string) {
	updatedMsg := "WARN: " + message
	e.message = updatedMsg
	color.Yellow(e.message)
	os.Exit(0)
}

func (e *Err) ShowError(message string) {
	updatedMsg := "FATAL: " + message
	e.message = updatedMsg
	color.Red(e.message)
	os.Exit(0)
}
