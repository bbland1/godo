package cmd

import "fmt"

var greeting = "welcome to goDo your todo list in the terminal allowing you to keep your fingers on the keys"

var unknown = "you have entered an unknown command please try again"

var userManual = `
usage: 
	goDo [command] [options]

options:
	-h, -help	used to get more information about a command
	
commands:
	help	show this message with an overview of all options and commands

use "goDo [command] -help" for more information about a command
`

func GetUserManual() string {
	return userManual
}

func GetGreeting() string {
	return greeting
}

func GetUnknown() string {
	return unknown
}

func DisplayUserManual() {
	fmt.Println(GetUserManual())
}

func DisplayGreeting() {
	fmt.Println(GetGreeting())
}

func DisplayUnknown() {
	fmt.Println(GetUnknown())
}
