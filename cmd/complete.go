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

func statusFunc(w io.Writer, database *sql.DB, args []string, cmd *Command) int {
	idFlagValue := cmd.flags.Lookup("id").Value.String()
	descriptionValue := cmd.flags.Lookup("d").Value.String()
	
	if idFlagValue == "" && descriptionValue == "" {
		fmt.Fprintln(w, "an id or task description needs to be passed to mark something as complete")
		return 1
	}
	
	if len(args) == 0 || args[0] == ""{
		fmt.Fprintln(w, "a status needs to be passed")
		return 1
	}
	
	statusValue, err := strconv.ParseBool(args[0])
	if err != nil {
		fmt.Fprintln(w, "status has to be 'true' or 'false' to update")
		return 1
	}
	if idFlagValue != "" {
		idNum, err := strconv.Atoi(idFlagValue)
		if err != nil {
			fmt.Fprintln(w, "an int needs to be passed for the id")
			return 1
		}

		if err := task.UpdateTaskCompletionStatus(database, idNum, statusValue); err != nil {
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

func NewStatusCommand(w io.Writer, db *sql.DB, exitCode *int) *Command {
	command := &Command{
		flags: flag.NewFlagSet("status", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			*exitCode = statusFunc(w, db, args, cmd)
		},
	}

	command.flags.String("id", "", "the id of the task to have status change")
	command.flags.String("d", "", "the description of the task to have status change")

	command.flags.Usage = func() {
		fmt.Fprintln(w, StatusUsage)
	}
	return command
}
