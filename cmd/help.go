package cmd

import (
	"flag"
	"fmt"
	"io"
	"text/tabwriter"
)

const HelpUsage = `print the user manual of goDo to given an overview of how to use the app

usage:
	goDo help

there are no additional options for help`

// The message that is displayed when the app starts with no commands passed
const Greeting = `welcome to goDo your todo list in the terminal allowing you to keep your fingers on the keys

to learn more about how to use:
	goDo help`


// Prints the UserManual to the terminal to show user how to use app
func DisplayUserManual(w io.Writer) int {
	tw := tabwriter.NewWriter(w, 0, 8, 2, ' ', 0)

	fmt.Fprintln(tw, "Usage:\n  goDo [command] [options]")

	fmt.Fprintln(tw, "\nOptions:")
	fmt.Fprintln(tw, "  -h\tShow more information about a command")
	fmt.Fprintln(tw, "  -verbose\tPrint detailed output when available")
	fmt.Fprintln(tw, "\nCommands:")

	for _, cmd := range registeredCommands {
		fmt.Fprintf(tw, "  %s\t- %s\n", cmd.GetName(), cmd.GetDescription())
	}

	tw.Flush()
	return 0
}

// Prints the welcome message to the terminal when the app is called with no commands passed
func DisplayGreeting(w io.Writer) int {
	fmt.Fprintln(w, Greeting)
	return 0
}

// NewHelpCommand is called to pull up the usage or userManual of how to use goDo
func NewHelpCommand(stdout, stderr io.Writer, exitCode *int) *BaseCommand {
	command := &BaseCommand{
		name: "help",
		description: "show this message with an overview of all options and commands",
		flags: flag.NewFlagSet("help", flag.ExitOnError),
		output: stdout,
		errOutput: stderr,
		execute: func(cmd *BaseCommand, args []string) {
			*exitCode = DisplayUserManual(cmd.output)
		},
	}

	command.flags.Usage = func() {
		fmt.Fprintln(command.output, HelpUsage)
	}

	return command
}
