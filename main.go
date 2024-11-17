package main

import (
	"fmt"
	"os"

	"github.com/bbland1/goDo/cmd"
)

func Execute(args []string) {

}

func main() {

	if len(os.Args) < 2 {
		cmd.DisplayGreeting()
		os.Exit(0)
	}

	var command *cmd.Command

	passedCommand := os.Args[1]
	passedArgs := os.Args[2:]

	switch passedCommand {
	case "help":
		command = cmd.NewHelpCommand()
	default:
		fmt.Printf("unknown command passed to goDo: %s\n\n", passedCommand)
		os.Exit(1)
	}

	command.Init(passedArgs)
	command.Run()

}

func usageAndExit(msg string, code int) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintln(os.Stderr)
	}

	os.Exit(code)
}