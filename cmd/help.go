package cmd

import (
	"flag"
	"fmt"
	"io"
)

const HelpUsage = `print the user manual of goDo to given an overview of how to use the app

usage:
	goDo help

there are no additional options for help`

// The message that is displayed when the app starts with no commands passed
const Greeting = `welcome to goDo your todo list in the terminal allowing you to keep your fingers on the keys

to learn more about how to use:
	goDo help`

// The base help message for the app, showing an overview of how it works
const UserManual = `usage: 
	goDo [command] [options]

options:
	-h		used to get more information about a command
	-verbose 	if a command has this option it will print the long form of the information
	
commands:
	help		show this message with an overview of all options and commands
	version 	display the version info of installed goDo
	add		add a new item to your list
	edit 		edit the description of a task
	status 		change the status of a task
	delete 		remove a task from the list

use "goDo [command] -help" for more information about a command`

// Prints the UserManual to the terminal to show user how to use app
func DisplayUserManual(w io.Writer) {
	fmt.Fprintln(w, UserManual)
}

// Prints the welcome message to the terminal when the app is called with no commands passed
func DisplayGreeting(w io.Writer) {
	fmt.Fprintln(w, Greeting)
}

// NewHelpCommand is called to pull up the usage or userManual of how to use goDo
func NewHelpCommand(stdout, stderr io.Writer) *BaseCommand {
	command := &BaseCommand{
		name: "help",
		description: "show this message with an overview of all options and commands",
		flags: flag.NewFlagSet("help", flag.ExitOnError),
		output: stdout,
		errOutput: stderr,
		execute: func(cmd *BaseCommand, args []string) {
			DisplayUserManual(cmd.output)
		},
	}

	command.flags.Usage = func() {
		fmt.Fprintln(command.output, HelpUsage)
	}

	RegisterCommand(command)

	return command
}
