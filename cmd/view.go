package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/bbland1/goDo/task"
)

const ViewUsage = `lists out all the tasks or a specific task
usage:
	godo view [options]
	
options:
	-h 	show the usage for the command
	-id show all the information for the specific task`

func viewFunc(database *sql.DB, args []string, cmd *BaseCommand) int {
	tasks, err := task.GetAllTasks(database)
	if err != nil {
		fmt.Fprintf(cmd.errOutput, "database error: %v\n", err)
		return 1
	}
	tw := tabwriter.NewWriter(cmd.output, 0, 8, 2, ' ', 0)
	fmt.Fprintln(tw, "ID\tTask Description\tStatus")

	for _, storedTask := range tasks {
		fmt.Fprintf(tw, "%d\t%s\t%t\n", storedTask.ID, storedTask.Description, storedTask.IsCompleted)
	}

	tw.Flush()
	return 0
}

func NewViewCommand(stdout, stderr io.Writer, db *sql.DB, exitCode *int) *BaseCommand {
	command := &BaseCommand{
		name: "view",
		description: "view all tasks or a specific task",
		flags: flag.NewFlagSet("view", flag.ExitOnError),
		output: stdout,
		errOutput: stderr,
		execute: func(cmd *BaseCommand, args []string) {
			*exitCode = viewFunc(db, args, cmd)
		},
	}

	command.flags.Usage = func() {
		fmt.Fprintln(command.output, ViewUsage)
	}

	return command
}