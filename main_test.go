package main

import (
	"testing"
)

func TestGoDoHello(t *testing.T) {
	got := openingMessage()
	want := "Welcome to goDo your todo list in the terminal allowing you to keep your fingers on the keys."

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGoDoUnknownCommand(t *testing.T) {

}
