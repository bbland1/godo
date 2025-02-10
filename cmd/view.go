package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
)

const ViewUsage = `lists out all the tasks or a specific task
usage:
	godo view [options]
	
options:
	-h 	show the usage for the command
	-id show all the information for the specific task`

func NewViewCommand(stdout, stderr io.Writer, database *sql.DB, exitCode *int) *BaseCommand {
	command := &BaseCommand{
		name: "view",
		description: "view all tasks or a specific task",
		flags: flag.NewFlagSet("view", flag.ExitOnError),
		output: stdout,
		errOutput: stderr,
		execute: func(cmd *BaseCommand, args []string) {
			// *exitCode = viewFunc()
		},
	}

	command.flags.Usage = func() {
		fmt.Fprintln(command.output, ViewUsage)
	}

	return command
}