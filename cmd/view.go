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
	-id show all the information for the specific task
	-status show all the tasks with a specific status`

func viewFunc(stderr io.Writer, database *sql.DB, args []string, cmd *BaseCommand) int {

	idFlagValue := cmd.flags.Lookup("id").Value.String()
	statusFlagValue := cmd.flags.Lookup("status").Value.String()

	// if idFlagValue == "" && statusFlagValue == "" {
	// 	tasks, err := task.GetAllTasks(database)
	// 	if err != nil {
	// 		fmt.Fprintf(cmd.errOutput, "database error: %v\n", err)
	// 		return 1
	// 	}
	// 	tw := tabwriter.NewWriter(cmd.output, 0, 8, 2, ' ', 0)
	// 	fmt.Fprintln(tw, "ID\tTask Description\tStatus")

	// 	for _, storedTask := range tasks {
	// 		fmt.Fprintf(tw, "%d\t%s\t%t\n", storedTask.ID, storedTask.Description, storedTask.IsCompleted)
	// 	}

	// 	tw.Flush()
	// 	return 0
	// }

	if idFlagValue != "" && statusFlagValue != "" {
		fmt.Fprintln(stderr, "passing an id & status argument causes conflict, try again passing one or the other")
		return 1
	}

	if idFlagValue != "" {
		// storedTask, err := task.GetATaskByID(database, in)
	}

	// if len(args) == 0 || args[0] == "" {
	// 	tasks, err := task.GetAllTasks(database)
	// 	if err != nil {
	// 		fmt.Fprintf(cmd.errOutput, "database error: %v\n", err)
	// 		return 1
	// 	}
	// 	tw := tabwriter.NewWriter(cmd.output, 0, 8, 2, ' ', 0)
	// 	fmt.Fprintln(tw, "ID\tTask Description\tStatus")

	// 	for _, storedTask := range tasks {
	// 		fmt.Fprintf(tw, "%d\t%s\t%t\n", storedTask.ID, storedTask.Description, storedTask.IsCompleted)
	// 	}

	// 	tw.Flush()
	// 	return 0
	// } 

	

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
			*exitCode = viewFunc(cmd.errOutput, db, args, cmd)
		},
	}

	command.flags.Int64("id", 0, "id fo the specific task to look at")
	command.flags.String("status", "", "completion status to filter the list on")

	command.flags.Usage = func() {
		fmt.Fprintln(command.output, ViewUsage)
	}

	return command
}
