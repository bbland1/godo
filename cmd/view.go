package cmd

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/bbland1/goDo/task"
)

var (
	ErrMultiFlagsPassed = errors.New("passing an id & status argument causes conflict, try again passing one or the other")
)

const ViewUsage = `lists out all the tasks or a specific task
usage:
	godo view [options]
	
options:
	-h 	show the usage for the command
	-id show all the information for the specific task
	-status show all the tasks with a specific status`

func viewFunc(database *sql.DB, args []string, cmd *BaseCommand) int {
	idFlag := cmd.passedFlags["id"]
	statusFlag := cmd.passedFlags["status"]

	if idFlag && statusFlag {
		fmt.Fprintln(cmd.errOutput, ErrMultiFlagsPassed)
		return 1
	}

	if idFlag {
		idFlagValue := cmd.flags.Lookup("id").Value.(flag.Getter).Get().(int64)

		storedTask, err := task.GetATaskByID(database, idFlagValue)
		if err != nil {
			fmt.Fprintf(cmd.errOutput, "database error: %v\n", err)
			return 1

		}

		tw := tabwriter.NewWriter(cmd.output, 0, 8, 2, ' ', 0)
		fmt.Fprintln(tw, "ID\tTask Description\tStatus")
		fmt.Fprintf(tw, "%d\t%s\t%t\n", storedTask.ID, storedTask.Description, storedTask.IsCompleted)

		tw.Flush()
		return 0
	}

	if statusFlag {
		statusFlagValue := cmd.flags.Lookup("status").Value.(flag.Getter).Get().(bool)

		storedTasks, err := task.GetTasksByStatus(database, statusFlagValue)
		if err != nil {
			fmt.Fprintf(cmd.errOutput, "database error: %v\n", err)
			return 1
		}

		tw := tabwriter.NewWriter(cmd.output, 0, 8, 2, ' ', 0)
		fmt.Fprintln(tw, "ID\tTask Description\tStatus")
		
		for _, storedTask := range storedTasks {
			fmt.Fprintf(tw, "%d\t%s\t%t\n", storedTask.ID, storedTask.Description, storedTask.IsCompleted)
		}

		tw.Flush()
		return 0
	}

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
		name:        "view",
		description: "view all tasks or a specific task",
		flags:       flag.NewFlagSet("view", flag.ExitOnError),
		output:      stdout,
		errOutput:   stderr,
		execute: func(cmd *BaseCommand, args []string) {
			*exitCode = viewFunc(db, args, cmd)
		},
		passedFlags: make(map[string]bool),
	}

	command.flags.Int64("id", 0, "id fo the specific task to look at")
	command.flags.Bool("status", false, "completion status to filter the list on")

	// loop through all of the registered flags and see if it was set by user
	// when user does pass it will set up the flag to be tracked when passed
	command.flags.VisitAll(func(f *flag.Flag) {
		f.Value = &TrackedFlags{
			Value:         f.Value,
			name:          f.Name,
			passedCommand: command,
		}
	})

	command.flags.Usage = func() {
		fmt.Fprintln(command.output, ViewUsage)
	}

	return command
}
