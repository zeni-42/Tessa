package model

import (
	"os"

	"github.com/fatih/color"
)

type Err struct {
	message string
}

func (e *Err) ShowInfo(message string) {
	e.message = message
	color.Blue(e.message)
}

func (e *Err) ShowWarning(message string) {
	e.message = message
	color.Yellow(e.message)
	os.Exit(0)
}

func (e *Err) ShowError(message string) {
	e.message = message
	color.Red(e.message)
	os.Exit(0)
}
