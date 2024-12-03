package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bbland1/goDo/task"
)

const AddUsage = `add a new task to your list
usage:
	goDo add [options]
	
options:
	-h 	passed to pull up more info on how to use the add command further`

func addFunc(w io.Writer, database *sql.DB, args []string) {
	if len(args) == 0 {
		fmt.Fprintf(w, "this is an add command, and this is everything else, %s\n", args)
		return
	}

	newTask := task.CreateTask(strings.TrimSpace(args[0]))
	if err := task.AddTask(database, newTask); err != nil {
		fmt.Fprintf(os.Stderr, "task did not add to the database: %v", err)
	}
}

func NewAddCommand(w io.Writer, database *sql.DB) *Command {
	command := &Command{
		flags: flag.NewFlagSet("add", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			addFunc(w, database, args)
		},
	}

	command.flags.Usage = func() {
		fmt.Fprintln(w, AddUsage)
	}

	return command
}
