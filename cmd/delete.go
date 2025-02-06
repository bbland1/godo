package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"strconv"

	"github.com/bbland1/goDo/task"
)

const DeleteUsage = `delete a task from the the list
usage:
	goDo delete [options]
	
options:
	-h show helpful tips for the delete command
	-id the id value of the task to be deleted
	-d the description of the task to be deleted`

func deleteFunc(stderr io.Writer, database *sql.DB, cmd *BaseCommand) int {
	idFlagValue := cmd.flags.Lookup("id").Value.String()
	descriptionValue := cmd.flags.Lookup("d").Value.String()

	if idFlagValue == "" && descriptionValue == "" {
		fmt.Fprintln(stderr, "an id or task description needs to be passed for deletion to process")
		return 1
	}

	if idFlagValue != "" {
		idNum, err := strconv.ParseInt(idFlagValue, 10, 64)
		if err != nil {
			fmt.Fprintln(stderr, "an int needs to be passed for the id")
			return 1
		}

		if err := task.DeleteTask(database, idNum); err != nil {
			fmt.Fprintf(stderr, "database error: %v\n", err)
			return 1
		}
	}

	if descriptionValue != "" {
		storedTask, err := task.GetATaskByDescription(database, descriptionValue)
		if err != nil {
			fmt.Fprintln(stderr, "an int needs to be passed for the id")
			return 1
		}

		if err := task.DeleteTask(database, storedTask.ID); err != nil {
			fmt.Fprintf(stderr, "database error: %v\n", err)
			return 1
		}
	}

	return 0
}

func NewDeleteCommand(stdout, stderr io.Writer, db *sql.DB, exitCode *int) *BaseCommand {
	command := &BaseCommand{
		name: "delete",
		description: "remove a task from the list",
		flags: flag.NewFlagSet("delete", flag.ExitOnError),
		output: stdout,
		errOutput: stderr,
		execute: func(cmd *BaseCommand, args []string) {
			*exitCode = deleteFunc(cmd.errOutput, db, cmd)
		},
	}

	command.flags.String("id", "", "the id of the task to be deleted")
	command.flags.String("d", "", "the description of the task to be deleted")

	command.flags.Usage = func() {
		fmt.Fprintln(command.output, DeleteUsage)
	}
	return command
}
