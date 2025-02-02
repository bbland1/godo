package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
)

const CompleteUsage = `mark a task from the list as completed
usage
	goDo complete [options]
	
options:
	-h show helpful tips for the complete command
	-id the id of the task to mark completed
	-d the description of the task to mark completed`

func completeFunc(w io.Writer, database *sql.DB, cmd *Command) int {
	return 0
}

func NewCompleteCommand(w io.Writer, db *sql.DB, exitCode *int) *Command {
	command := &Command{
		flags: flag.NewFlagSet("complete", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			*exitCode = completeFunc(w, db, cmd)
		},
	}

	command.flags.Usage = func() {
		fmt.Fprintln(w, CompleteUsage)
	}
	return command
}