package cmd

import (
	"flag"
	"fmt"
)

// The message that is displayed when the app starts with no commands passed
const Greeting = `welcome to goDo your todo list in the terminal allowing you to keep your fingers on the keys

to learn more about how to use:
	goDo help`

// The base help message for the app, showing an overview of how it works
const UserManual = `usage: 
	goDo [command] [options]

options:
	-h, -help	used to get more information about a command
	
commands:
	help	show this message with an overview of all options and commands
	add	add a new itm to your todo list

use "goDo [command] -help" for more information about a command`

// Prints the UserManual to the terminal to show user how to use app
func DisplayUserManual() {
	fmt.Println(UserManual)
}

// Prints the welcome message to the terminal when the app is called with no commands passed
func DisplayGreeting() {
	fmt.Println(Greeting)
}

// NewHelpCommand is called to pull up the usage or userManual of how to use goDo
func NewHelpCommand() *Command {
	return &Command{
		flags: flag.NewFlagSet("help", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			DisplayUserManual()
		},
	}
}
