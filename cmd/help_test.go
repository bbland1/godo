package cmd

import (
	"testing"
)

func TestDisplayGreeting(t *testing.T) {
	got := GetGreeting()
	want := "Welcome to goDo your todo list in the terminal allowing you to keep your fingers on the keys."

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDisplayUsage(t *testing.T) {
	got := GetUserManual()
	want := `
usage: 
	goDo [command] [options]

options:
	-h, -help	used to get more information about a command
	
commands:
	help	show this message with an overview of all options and commands

use "goDo [command] -help" for more information about a command
`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDisplayUnknown(t *testing.T) {
	got := GetUnknown()
	want := "You have entered an unknown command please try again."

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
