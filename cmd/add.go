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

func addFunc(w io.Writer, database *sql.DB, args []string) int {
	if len(args) == 0 || args[0] == ""{
		fmt.Fprintf(w, "a description string needs to be passed to add a task\n")
		return 1
	}

	newTask := task.CreateTask(strings.TrimSpace(args[0]))
	// ? this would maybe need to use the id that is now returned
	if _, err := task.AddTask(database, newTask); err != nil {
		fmt.Fprintf(w, "database error: %v\n", err)
		return 1
	}
	return 0
}

func NewAddCommand(w io.Writer, database *sql.DB, exitCode *int) (*Command) {
	command := &Command{
		flags: flag.NewFlagSet("add", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			*exitCode = addFunc(w, database, args)
		},
	}

	command.flags.Usage = func() {
		fmt.Fprintln(w, AddUsage)
	}

	return command
}
