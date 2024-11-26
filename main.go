package main

import (
	"fmt"
	"io"
	"os"

	"github.com/bbland1/goDo/cmd"
)

func usageAndExit(w io.Writer, msg string, code int) int {
	if msg != "" {
		fmt.Fprint(w, msg)
		fmt.Fprintln(w)
	}

	return code
}

func runAppLogic(w io.Writer, args []string) int {
	if len(args) < 2 {
		cmd.DisplayGreeting(w)
		return 0
	}

	var command *cmd.Command

	passedCommand := args[1]
	passedArgs := args[2:]

	switch passedCommand {
	case "help":
		command = cmd.NewHelpCommand(w)
	case "version":
		command = cmd.NewVersionCommand(w)
	default:
		usageAndExit(w, fmt.Sprintf("unknown command passed to goDo: %s\n", passedCommand), 1)
	}

	if command == nil {
		return 1
	}

	command.Init(passedArgs)
	command.Run()
	return 0
}

func main() {
	exitCode := runAppLogic(os.Stdout, os.Args)

	os.Exit(exitCode)

}
