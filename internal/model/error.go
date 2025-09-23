package model

import (
	"os"

	"github.com/fatih/color"
)

type Err struct {
	message	string
}

func (e *Err) ShowInfo (message string) {
	e.message = message
	color.Blue(string(e.message))
}

func (e *Err) ShowWarning(message string) {
	e.message = message
	color.Yellow(string(e.message))
	os.Exit(1)
}

func (e *Err) ShowError(message string) {
	e.message = message
	color.Red(string(e.message))
	os.Exit(1)
}
