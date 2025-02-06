package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/bbland1/goDo/task"
)

const AddUsage = `add a new task to your list
usage:
	goDo add [options] "a task description"
	
options:
	-h 	show the usage info for the command.`

func addFunc(stderr io.Writer, database *sql.DB, args []string) int {
	if len(args) == 0 || args[0] == "" {
		fmt.Fprintf(stderr, "a description string needs to be passed to add a task\n")
		return 1
	}

	newTask := task.CreateTask(strings.TrimSpace(args[0]))
	// ? this would maybe need to use the id that is now returned
	if _, err := task.AddTask(database, newTask); err != nil {
		fmt.Fprintf(stderr, "database error: %v\n", err)
		return 1
	}
	return 0
}

func NewAddCommand(stdout, stderr io.Writer, database *sql.DB, exitCode *int) *BaseCommand {
	command := &BaseCommand{
		name: "add",
		description: "add a task to list",
		flags: flag.NewFlagSet("add", flag.ExitOnError),
		output: stdout,
		errOutput: stderr,
		execute: func(cmd *BaseCommand, args []string) {
			*exitCode = addFunc(cmd.errOutput, database, args)
		},
	}

	command.flags.Usage = func() {
		fmt.Fprintln(command.output, AddUsage)
	}

	return command
}
