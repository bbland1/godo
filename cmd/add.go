package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"io"

	// "github.com/bbland1/goDo/task"
)

const AddUsage = `add a new task to your list
usage:
	goDo add [options]
	
options:
	-h 	passed to pull up more info on how to use the add command further`

func NewAddCommand(w io.Writer, database *sql.DB) *Command {
	command := &Command{
		flags: flag.NewFlagSet("add", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			// newTask := task.CreateTask(args[0])
			fmt.Fprintf(w, "this is an add command, and this is everything else, %s\n", args)
			// fmt.Fprintf(w, "this is the task %v", newTask)
		},
	}

	command.flags.Usage = func() {
		fmt.Fprintln(w, AddUsage)
	}

	return command
}
