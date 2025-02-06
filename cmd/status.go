package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"strconv"

	"github.com/bbland1/goDo/task"
)

const StatusUsage = `update the status of a task
usage
	goDo status [options]
	
options:
	-h show helpful tips for the status command
	-id the id of the task to change status
	-d the description of the task to change status`

func statusFunc(stderr io.Writer, database *sql.DB, args []string, cmd *BaseCommand) int {
	idFlagValue := cmd.flags.Lookup("id").Value.String()
	descriptionValue := cmd.flags.Lookup("d").Value.String()

	if idFlagValue == "" && descriptionValue == "" {
		fmt.Fprintln(stderr, "an id or task description needs to be passed to mark something as complete")
		return 1
	}

	if len(args) == 0 || args[0] == "" {
		fmt.Fprintln(stderr, "a status needs to be passed")
		return 1
	}

	statusValue, err := strconv.ParseBool(args[0])
	if err != nil {
		fmt.Fprintln(stderr, "status has to be 'true' or 'false' to update")
		return 1
	}
	if idFlagValue != "" {
		idNum, err := strconv.ParseInt(idFlagValue, 10, 64)
		if err != nil {
			fmt.Fprintln(stderr, "an int needs to be passed for the id")
			return 1
		}

		if err := task.UpdateTaskStatus(database, idNum, statusValue); err != nil {
			fmt.Fprintf(stderr, "database error: %v\n", err)
			return 1
		}
	}

	if descriptionValue != "" {
		storedTask, err := task.GetATaskByDescription(database, descriptionValue)
		if err != nil {
			fmt.Fprintln(stderr, "a task with that description wasn't found")
			return 1
		}

		if err := task.UpdateTaskStatus(database, storedTask.ID, statusValue); err != nil {
			fmt.Fprintf(stderr, "database error: %v\n", err)
			return 1
		}
	}

	return 0
}

func NewStatusCommand(stdout, stderr io.Writer, db *sql.DB, exitCode *int) *BaseCommand {
	command := &BaseCommand{
		name: "status",
		description: "update the completion status of a task",
		flags: flag.NewFlagSet("status", flag.ExitOnError),
		output: stdout,
		errOutput: stderr,
		execute: func(cmd *BaseCommand, args []string) {
			*exitCode = statusFunc(cmd.errOutput, db, args, cmd)
		},
	}

	command.flags.String("id", "", "the id of the task to have status change")
	command.flags.String("d", "", "the description of the task to have status change")

	command.flags.Usage = func() {
		fmt.Fprintln(command.output, StatusUsage)
	}
	return command
}
