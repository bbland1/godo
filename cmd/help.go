package cmd

import "fmt"

var greeting = "Welcome to goDo your todo list in the terminal allowing you to keep your fingers on the keys."

var unknown = "You have entered an unknown command please try again."

var helpMessage = `

usage: goDo command [options]

options:
	-h, -help	used to get more information about a command
	
commands:
	help	show this message with an overview of all options and commands

`

func GetHelpMessage() string {
	return helpMessage
}

func GetGreeting() string {
	return greeting
}

func GetUnknown() string {
	return unknown
}

func DisplayHelpMessage() {
	fmt.Println(GetHelpMessage())
}

func DisplayGreeting() {
	fmt.Println(GetGreeting())
}

func DisplayUnknown() {
	fmt.Println(GetUnknown())
}
