package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/bbland1/goDo/cmd"
	// "github.com/bbland1/goDo/task"
)

func TestNoCommandArgs(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer

	args := []string{"godo"}

	exitCode := runAppLogic(&bufferOut, &bufferErr, args)

	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	expectedOutput := cmd.Greeting
	output := strings.TrimSpace(bufferOut.String())
	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestUnknownCommand(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer

	args := []string{"godo", "unknown"}

	exitCode := runAppLogic(&bufferOut, &bufferErr, args)

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedErrMsg := fmt.Sprintf("Unknown command: %s", "unknown")
	output := strings.TrimSpace(bufferErr.String())

	if output != expectedErrMsg {
		t.Errorf("Expected output: %q, got: %q", expectedErrMsg, output)
	}
}

func TestValidCommandPassed(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	cmd.RegisterCommand(cmd.NewHelpCommand(&bufferOut, &bufferErr, &exitCode))

	args := []string{"godo", "help"}

	exitCode = runAppLogic(&bufferOut, &bufferErr, args)

	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	expectedOutput := cmd.UserManual
	output := strings.TrimSpace(bufferOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestCommandInitFail(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer

	failure := &failingCommand{
		name: "fail",
		description: "a fail",
	}
	cmd.RegisterCommand(failure)

	args := []string{"goDo", "fail"}

	exitCode := runAppLogic(&bufferOut, &bufferErr, args)

	if exitCode != 1 {
		t.Errorf("Expected exit code 1, but got %d", exitCode)
	}

	expectedErrorMsg := "Error initializing command: fail"
	output := strings.TrimSpace(bufferErr.String())
	if output != expectedErrorMsg {
		t.Errorf("Expected stderr to contain '%s', but got '%s'", expectedErrorMsg, bufferErr.String())
	}
}

type failingCommand struct {
	name string
	description string
}

func (cmd *failingCommand) Init(args []string) error {
	return fmt.Errorf("%s", cmd.name)
}

func (cmd *failingCommand) Run() {}

func (cmd *failingCommand) Called() bool {
	return false
}

func (cmd *failingCommand) GetName() string {
	return cmd.name
}

func (cmd *failingCommand) GetDescription() string {
	return cmd.description
}