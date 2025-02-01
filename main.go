package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"

	"github.com/bbland1/goDo/cmd"
	"github.com/bbland1/goDo/task"
)
const dbFile = "goDo.db"

func usageAndExit(w io.Writer, msg string, code int) int {
	if msg != "" {
		fmt.Fprint(w, msg)
		fmt.Fprintln(w)
	}

	return code
}

func runAppLogic(w io.Writer, args []string, database *sql.DB) int {
	var exitCode int
	var command *cmd.Command
	
	if len(args) < 2 {
		cmd.DisplayGreeting(w)
		return 0
	}


	passedCommand := args[1]
	passedArgs := args[2:]

	switch passedCommand {
	case "help":
		command = cmd.NewHelpCommand(w)
	case "version":
		command = cmd.NewVersionCommand(w)
	case "add":
		command = cmd.NewAddCommand(w, database, &exitCode)
	case "delete":
		command = cmd.NewDeleteCommand(w, database, &exitCode)
	default:
		usageAndExit(w, fmt.Sprintf("unknown command passed to goDo: %s\n", passedCommand), 1)
	}

	if command == nil {
		return 1
	}

	command.Init(passedArgs)
	command.Run()
	return exitCode
}

func main() {
	db, err := task.InitDatabase(dbFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize database: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()

	exitCode := runAppLogic(os.Stdout, os.Args, db)

	os.Exit(exitCode)

}
