package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
)

const EditUsage = `edit the description of a task
usage:
	goDo edit [edit]
	
options:
	-h show helpful tips for the edit command
	-id the id value of the task to be edited
	-d the description of the task to be edited`

func editFunc(w io.Writer, database *sql.DB, cmd *Command) int {
	return 0
}

func NewEditCommand(w io.Writer, db *sql.DB, exitCode *int) *Command {
	command := &Command{
		flags: flag.NewFlagSet("edit", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			*exitCode = editFunc(w, db, cmd)
		},
	}

	command.flags.String("id", "", "the id of the task to be edited")
	command.flags.String("d","", "the description of the task to be edited")

	command.flags.Usage = func() {
		fmt.Fprintln(w, EditUsage)
	}
	return command
}