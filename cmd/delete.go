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

func deleteFunc(w io.Writer, database *sql.DB,  cmd *Command) int {
	idFlagValue := cmd.flags.Lookup("id").Value.String()
	descriptionValue := cmd.flags.Lookup("d").Value.String()

	if idFlagValue == "" && descriptionValue == "" {
		fmt.Fprintln(w, "an id or task description needs to be passed for deletion to process")
		return 1
	}

	if idFlagValue != "" {
		idNum, err := strconv.Atoi(idFlagValue)
		if err != nil {
			fmt.Fprintln(w, "an int needs to be passed for the id")
			return 1
		}

		if err := task.DeleteTask(database, idNum); err != nil {
			fmt.Fprintf(w, "database error: %v\n", err)
			return 1
		}
	}

	if descriptionValue != "" {
		storedTask, err := task.GetATaskByDescription(database, descriptionValue)
		if err != nil {
			fmt.Fprintln(w, "an int needs to be passed for the id")
			return 1
		}

		if err := task.DeleteTask(database, storedTask.ID); err != nil {
			fmt.Fprintf(w, "database error: %v\n", err)
			return 1
		}
	}
	
	return 0
}

func NewDeleteCommand(w io.Writer, db *sql.DB, exitCode *int) *Command {
	command := &Command{
		flags: flag.NewFlagSet("delete", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			*exitCode = deleteFunc(w, db, cmd)
		},
	}

	command.flags.String("id", "", "the id of the task to be deleted")
	command.flags.String("d", "", "the description of the task to be deleted")

	command.flags.Usage = func ()  {
		fmt.Fprintln(w, DeleteUsage)
	}
	return command
}