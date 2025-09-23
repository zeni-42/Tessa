package main

import "tessa/internal/config"

func main() {
	if !config.IsConfigured() {
		config.Configure()
	}
}