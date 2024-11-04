package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bbland1/goDo/cmd"
)

func openingMessage() string {
	return "Welcome to goDo your todo list in the terminal allowing you to keep your fingers on the keys."
}

func main() {

	flag.Usage = func() {
		cmd.DisplayUserManual()
	}

	helpFlag := flag.Bool("help", false, "show help message")

	flag.Parse()

	if len(os.Args) < 2 || *helpFlag {
		cmd.DisplayUserManual()
		os.Exit(0)
	}

	command := os.Args[1]


	switch command {
	case "help":
		cmd.DisplayUserManual()
	default:
		fmt.Printf("unknown command passed to goDo: %s\n\n", command)
		flag.Usage()
		os.Exit(1)
	}

}
