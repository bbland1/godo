package cmd

import (
	"flag"
	"fmt"
	"io"
)

const AddUsage = `add a new task to your list
usage:
	goDo add [options]
	
options:
	-h 	passed to pull up more info on how to use the add command further`

func NewAddCommand(w io.Writer) *Command {
	command := &Command{
		flags: flag.NewFlagSet("add", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			fmt.Fprintf(w, "this is an add command, and this is everything else, %s\n", args)
		},
	}
	
	command.flags.Usage = func() {
		fmt.Fprintln(w, AddUsage)
	}

	return command
}