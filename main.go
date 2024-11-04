package main

import (
	// "fmt"
	"os"

	"github.com/bbland1/goDo/cmd"
)

func openingMessage() string {
	return "Welcome to goDo your todo list in the terminal allowing you to keep your fingers on the keys."
}

func main() {
	if len(os.Args) < 2 {
		cmd.DisplayGreeting()
	}

	switch os.Args[1] {
	case "help":
		cmd.DisplayHelpMessage()
	}
}
