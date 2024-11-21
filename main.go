package main

import (
	"fmt"
	"os"

	"github.com/bbland1/goDo/cmd"
)

func usageAndExit(msg string, code int) int {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintln(os.Stderr)
	}

	return code
}

func runAppLogic(args []string) int {
	if len(args) < 2 {
		cmd.DisplayGreeting()
		return 0
	}

	var command *cmd.Command

	passedCommand := args[1]
	passedArgs := args[2:]

	switch passedCommand {
	case "help":
		command = cmd.NewHelpCommand()
	default:
		usageAndExit(fmt.Sprintf("unknown command passed to goDo: %s\n", passedCommand), 1)
	}

	if command == nil {
		return 1
	}

	command.Init(passedArgs)
	command.Run()
	return 0
}

func main() {
	exitCode := runAppLogic(os.Args)

	os.Exit(exitCode)

}