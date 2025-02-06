package main

import (
	// "database/sql"
	"fmt"
	"io"
	"os"

	"github.com/bbland1/goDo/cmd"
	"github.com/bbland1/goDo/task"
)

const dbFile = "goDo.db"

func main() {
	var exitCode int
	db, err := task.InitDatabase(dbFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize database: %v\n", err)
		os.Exit(1)
	}
	
	defer db.Close()

	// register the commands to be used manually
	// todo: maybe move to the init once make sure new command logic works for all
	cmd.RegisterCommand(cmd.NewHelpCommand(os.Stdout, os.Stderr, &exitCode))
	cmd.RegisterCommand(cmd.NewVersionCommand(os.Stdout, os.Stderr, &exitCode))
	cmd.RegisterCommand(cmd.NewAddCommand(os.Stdout, os.Stderr, db, &exitCode))
	cmd.RegisterCommand(cmd.NewStatusCommand(os.Stdout, os.Stderr, db, &exitCode))
	cmd.RegisterCommand(cmd.NewDeleteCommand(os.Stdout, os.Stderr, db, &exitCode))
	cmd.RegisterCommand(cmd.NewEditCommand(os.Stdout, os.Stderr, db, &exitCode))
	
	exitCode = runAppLogic(os.Stdout, os.Stderr, os.Args)
	
	os.Exit(exitCode)
	
}

func runAppLogic(stdout, stderr io.Writer, args []string) int {

	if len(args) < 2 {
		cmd.DisplayGreeting(stdout)
		return 0
	}

	passedCommand := args[1]
	passedArgs := args[2:]

	command, exists := cmd.GetCommand(passedCommand)
	if !exists {
		fmt.Fprintf(stderr, "Unknown command: %s\n", passedCommand)
		// cmd.DisplayUserManual(os.Stderr)
		return 1
	}

	if err := command.Init(passedArgs); err != nil {
		fmt.Fprintf(stderr, "Error initializing command: %v\n", err)
		return 1
	}
	command.Run()
	return 0
}
