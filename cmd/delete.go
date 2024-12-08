package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
)

const DeleteUsage = `delete a task from the the list
usage:
	goDo delete [options]
	
options:
	-h show helpful tips for the delete command`

func deleteFunc(w io.Writer) int {
	return 0
}

func NewDeleteCommand(w io.Writer, db *sql.DB, exitCode *int) *Command {
	command := &Command{
		flags: flag.NewFlagSet("delete", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			*exitCode = deleteFunc(w)
		},
	}

	command.flags.String("id", "", "the id used to get and delete a specific task")
	command.flags.String("description", "", "the id used to get and delete a specific task")

	command.flags.Usage = func ()  {
		fmt.Fprintln(w, DeleteUsage)
	}
	return command
}