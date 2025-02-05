package cmd

import (
	"bytes"
	"flag"
	"testing"
)

func TestCommandInit(t *testing.T) {
	cmd := &BaseCommand{
		flags: flag.NewFlagSet("tester", flag.ContinueOnError),
	}

	cmd.flags.String("name", "", "a test of string flag")

	sampleArgs := []string{"-name", "your name"}

	err := cmd.Init(sampleArgs)

	if err != nil {
		t.Errorf("Init() returned an error: %v", err)
	}

	nameFlag := cmd.flags.Lookup("name").Value.String()

	if nameFlag != "your name" {
		t.Errorf("Init() did not properly parse flag 'name'. Got = %q, want = %q", nameFlag, "your name")
	}

}

func TestCommandCalled(t *testing.T) {
	cmd := &BaseCommand{
		flags: flag.NewFlagSet("tester", flag.ContinueOnError),
	}

	cmd.flags.String("name", "", "a test of string flag")

	if cmd.Called() {
		t.Errorf("Called() should return false if method is before Init(), got= %t", cmd.Called())
	}

	sampleArgs := []string{"-name", "your name"}

	err := cmd.Init(sampleArgs)
	if err != nil {
		t.Errorf("Init() returned an error: %v", err)
	}

	if !cmd.Called() {
		t.Errorf("Called() should return true if after Init() is called, got= %t", cmd.Called())
	}

}

func TestCommandRun(t *testing.T) {
	var executed bool
	var passedArgs []string

	cmd := &BaseCommand{
		flags: flag.NewFlagSet("tester", flag.ContinueOnError),
		execute: func(cmd *BaseCommand, args []string) {
			executed = true
			passedArgs = args
		},
	}

	cmd.flags.String("name", "", "a test of string flag")

	sampleArgs := []string{"-name", "your name", "extra", "non-flag", "args"}

	err := cmd.Init(sampleArgs)
	if err != nil {
		t.Errorf("Init() returned an error: %v", err)
	}

	cmd.Run()

	if !executed {
		t.Errorf("Run() should execute the command, but executed = %t", executed)
	}

	expectedArgs := []string{"extra", "non-flag", "args"}
	if len(passedArgs) != len(expectedArgs) {
		t.Errorf("Expected %d arguments, got = %d. Args: %v", len(expectedArgs), len(passedArgs), passedArgs)
	}

	for i, expected := range expectedArgs {
		if passedArgs[i] != expected {
			t.Errorf("Expected arg[%d] to be %q, got %q", i, expected, passedArgs[i])
		}

	}

}

func TestRegisterCommand(t *testing.T) {
	testCmd := &BaseCommand{
		name:        "tester",
		description: "a tester command",
		flags:       flag.NewFlagSet("test", flag.ContinueOnError),
		execute:     func(cmd *BaseCommand, args []string) {},
	}

	RegisterCommand(testCmd)

	retrievedCmd, exists := GetCommand("tester")
	if !exists {
		t.Fatalf("Expected command 'tester' to be registered, but it was not.")
	}

	if retrievedCmd.GetName() != "tester" {
		t.Errorf("Expected command name to be 'mock', got %q", retrievedCmd.GetName())
	}

	if retrievedCmd.GetDescription() != "a tester command" {
		t.Errorf("Expected description to be 'Mock command for testing', got %q", retrievedCmd.GetDescription())
	}
}

func TestGetCommand_NotFOund(t *testing.T) {
	_, exists := GetCommand("fake")
	if exists {
		t.Errorf("GetCommand() should return false for an unregistered command, but it returned true.")
	}
}

func TestListCommands(t *testing.T) {
	registeredCommands = make(map[string]Command)

	RegisterCommand(&BaseCommand{
		name:        "cmd1",
		description: "command 1",
		flags:       flag.NewFlagSet("cmd1", flag.ContinueOnError),
		execute:     func(cmd *BaseCommand, args []string) {},
	})

	RegisterCommand(&BaseCommand{
		name:        "cmd2",
		description: "command 2",
		flags:       flag.NewFlagSet("cmd2", flag.ContinueOnError),
		execute:     func(cmd *BaseCommand, args []string) {},
	})

	var buffer bytes.Buffer

	ListCommands(&buffer)

	expectedOutput := "Available commands:\n cmd1 - command 1\n cmd2 - command 2\n"

	if buffer.String() != expectedOutput {
		t.Errorf("ListCommands() output mismatch.\nGot:\n%q\nWant:\n%q", buffer.String(), expectedOutput)
	}
}
