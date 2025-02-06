package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"strconv"

	"github.com/bbland1/goDo/task"
)

const EditUsage = `edit the description of a task
usage:
	goDo edit [edit]
	
options:
	-h show helpful tips for the edit command
	-id the id value of the task to be edited
	-d the description of the task to be edited`

func editFunc(w io.Writer, database *sql.DB, args []string, cmd *BaseCommand) int {
	idFlagValue := cmd.flags.Lookup("id").Value.String()
	descriptionValue := cmd.flags.Lookup("d").Value.String()

	if idFlagValue == "" && descriptionValue == "" {
		fmt.Fprintln(w, "an id or task description needs to be passed to edit a task")
		return 1
	}

	if len(args) == 0 || args[0] == "" {
		fmt.Fprintln(w, "a status needs to be passed")
		return 1
	}

	newDescription := args[0]

	if idFlagValue != "" {
		idNum, err := strconv.ParseInt(idFlagValue, 10, 64)
		if err != nil {
			fmt.Fprintln(w, "an int needs to be passed for the id")
			return 1
		}

		if err := task.UpdateTaskDescription(database, idNum, newDescription); err != nil {
			fmt.Fprintf(w, "database error: %v\n", err)
			return 1
		}
	}

	if descriptionValue != "" {
		storedTask, err := task.GetATaskByDescription(database, descriptionValue)
		if err != nil {
			fmt.Fprintln(w, "a task with that description wasn't found")
			return 1
		}

		if err := task.UpdateTaskDescription(database, storedTask.ID, newDescription); err != nil {
			fmt.Fprintf(w, "database error: %v\n", err)
			return 1
		}
	}
	return 0
}

func NewEditCommand(stdout, stderr io.Writer, db *sql.DB, exitCode *int) *BaseCommand {
	command := &BaseCommand{
		name: "edit",
		description: "edit the description of a task",
		flags: flag.NewFlagSet("edit", flag.ExitOnError),
		output: stdout,
		errOutput: stderr,
		execute: func(cmd *BaseCommand, args []string) {
			*exitCode = editFunc(cmd.errOutput, db, args, cmd)
		},
	}

	command.flags.String("id", "", "the id of the task to be edited")
	command.flags.String("d", "", "the description of the task to be edited")

	command.flags.Usage = func() {
		fmt.Fprintln(command.output, EditUsage)
	}
	return command
}
