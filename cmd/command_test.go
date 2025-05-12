package cmd

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
	"testing"
	"text/tabwriter"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	// "github.com/bbland1/goDo/task"
	// "github.com/stretchr/testify/require"
)

// func TestCommandInit(t *testing.T) {
// 	cmd := &BaseCommand{
// 		flags: flag.NewFlagSet("tester", flag.ContinueOnError),
// 	}

// 	cmd.flags.String("name", "", "a test of string flag")

// 	sampleArgs := []string{"-name", "your name"}

// 	err := cmd.Init(sampleArgs)

// 	if err != nil {
// 		t.Errorf("Init() returned an error: %v", err)
// 	}

// 	nameFlag := cmd.flags.Lookup("name").Value.String()

// 	if nameFlag != "your name" {
// 		t.Errorf("Init() did not properly parse flag 'name'. Got = %q, want = %q", nameFlag, "your name")
// 	}

// }

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

	var bufferOut bytes.Buffer

	ListCommands(&bufferOut)

	var bufferExpectedOutput bytes.Buffer

	tw := tabwriter.NewWriter(&bufferExpectedOutput, 0, 8, 2, ' ', 0)

	fmt.Fprintln(tw, "commands:")
	fmt.Fprintln(tw, "  cmd1\t- command 1")
	fmt.Fprintln(tw, "  cmd2\t- command 2")

	tw.Flush()

	output := strings.TrimSpace(bufferOut.String())
	expectedOutput := strings.TrimSpace(bufferExpectedOutput.String())

	if output != expectedOutput {
		t.Errorf("ListCommands() output mismatch.\nGot:\n%q\nWant:\n%q", bufferOut.String(), expectedOutput)
	}
}

func TestCommandInit(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectedOut string
		expectedErr bool
	}{
		{
			name:        "valid flag name",
			args:        []string{"-name", "your name"},
			expectedOut: "your name",
		},
		{
			name:        "missing flag name",
			args:        []string{},
			expectedErr: false,
		},
		{
			name:        "invalid flag name",
			args:        []string{"-invalid"},
			expectedErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {

			cmd := &BaseCommand{
				flags: flag.NewFlagSet("tester", flag.ContinueOnError),
			}

			cmd.flags.String("name", "", "a test of string flag")

			err := cmd.Init(testCase.args)

			if testCase.expectedErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			nameFlag := cmd.flags.Lookup("name").Value.String()
			assert.Equal(t, testCase.expectedOut, nameFlag)
		})
	}
}
