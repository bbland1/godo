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
)

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

			testCmd := &BaseCommand{
				flags: flag.NewFlagSet("tester", flag.ContinueOnError),
			}

			testCmd.flags.String("name", "", "a test of string flag")

			err := testCmd.Init(testCase.args)

			if testCase.expectedErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			nameFlag := testCmd.flags.Lookup("name").Value.String()
			assert.Equal(t, testCase.expectedOut, nameFlag)
		})
	}
}

func TestCommandCalled(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectedOut bool
		expectedErr bool
	}{
		{
			name:        "valid called response",
			args:        []string{"-name", "your name"},
			expectedOut: true,
			expectedErr: false,
		},
		{
			name:        "called method done before init",
			args:        []string{"-name", "your name"},
			expectedOut: false,
			expectedErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCmd := &BaseCommand{
				flags: flag.NewFlagSet("tester", flag.ContinueOnError),
			}

			testCmd.flags.String("name", "", "a test of string flag")

			if testCase.expectedErr {
				require.Equal(t, testCase.expectedOut, testCmd.Called())
				return
			}

			err := testCmd.Init(testCase.args)

			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedOut, testCmd.Called())
		})
	}
}

func TestCommandRegistration(t *testing.T) {
	tests := []struct {
		name         string
		command      *BaseCommand
		register     bool
		expectExists bool
		expectedErr  error
	}{
		{
			name: "command is registered properly",
			command: &BaseCommand{
				name:        "tester",
				description: "a test command",
				flags:       flag.NewFlagSet("test", flag.ContinueOnError),
				execute:     func(cmd *BaseCommand, args []string) {},
			},
			register:     true,
			expectExists: true,
		},
		{
			name: "called method before init",
			command: &BaseCommand{
				name:        "tester",
				description: "a test command",
				flags:       flag.NewFlagSet("test", flag.ContinueOnError),
				execute:     func(cmd *BaseCommand, args []string) {},
			},
			register:     false,
			expectExists: false,
		},
		{
			name: "get unregistered command",
			command: &BaseCommand{
				name:        "unknown",
				description: "a test command",
				flags:       flag.NewFlagSet("test", flag.ContinueOnError),
				execute:     func(cmd *BaseCommand, args []string) {},
			},
			register:     false,
			expectExists: false,
		},
		{
			name: "empty command name",
			command: &BaseCommand{
				name:        "",
				description: "no name",
				flags:       flag.NewFlagSet("test", flag.ContinueOnError),
				execute:     func(cmd *BaseCommand, args []string) {},
			},
			register:     true,
			expectExists: false,
			expectedErr:  ErrEmptyCmdName,
		},
		{
			name:         "nil passed to command",
			command:      nil,
			register:     true,
			expectExists: false,
			expectedErr:  ErrCommandNil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			ClearCommandRegistry()

			var err error
			if testCase.register {
				err = RegisterCommand(testCase.command)
				assert.ErrorIs(t, err, testCase.expectedErr)
			}

			var cmdName string
			if testCase.command != nil {
				cmdName = testCase.command.name
			}

			retrievedCmd, exists := GetCommand(cmdName)
			assert.Equal(t, testCase.expectExists, exists)

			if testCase.expectExists {
				assert.Equal(t, testCase.command.GetName(), retrievedCmd.GetName())
				assert.Equal(t, testCase.command.GetDescription(), retrievedCmd.GetDescription())
			}
		})
	}
}

func TestCommandListMethod(t *testing.T) {
	tests := []struct {
		name           string
		commands       []*BaseCommand
		expectedOutput string
	}{
		{
			name: "multiple commands",
			commands: []*BaseCommand{
				{
					name:        "cmd1",
					description: "command 1",
					flags:       flag.NewFlagSet("cmd1", flag.ContinueOnError),
					execute:     func(cmd *BaseCommand, args []string) {},
				},
				{
					name:        "cmd2",
					description: "command 2",
					flags:       flag.NewFlagSet("cmd2", flag.ContinueOnError),
					execute:     func(cmd *BaseCommand, args []string) {},
				},
			},
		},
		{
			name:     "no commands",
			commands: []*BaseCommand{},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			ClearCommandRegistry()

			for _, cmd := range testCase.commands {
				err := RegisterCommand(cmd)
				require.NoError(t, err)
			}

			var bufferOut bytes.Buffer
			ListCommands(&bufferOut)

			var bufferExpectedOutput bytes.Buffer

			tw := tabwriter.NewWriter(&bufferExpectedOutput, 0, 8, 2, ' ', 0)

			fmt.Fprintln(tw, "commands:")
			for _, cmd := range testCase.commands {
				fmt.Fprintf(tw, "  %s\t- %s\n", cmd.GetName(), cmd.GetDescription())

			}

			tw.Flush()

			expected := strings.TrimSpace(bufferExpectedOutput.String())
			actual := strings.TrimSpace(bufferOut.String())

			assert.Equal(t, expected, actual)
		})
	}
}

func TestCommandRun9(t *testing.T) {
	tests := []struct {
		name             string
		args             []string
		expectedArgs     []string
		expectedExecuted bool
		expectedErr      bool
		errorContains    string
	}{
		{
			name:             "valid flags and positional args",
			args:             []string{"-name", "your name", "extra", "non-flag", "args"},
			expectedArgs:     []string{"extra", "non-flag", "args"},
			expectedExecuted: true,
			expectedErr:      false,
		},
		{
			name:             "only flags, no positional args",
			args:             []string{"-name", "value"},
			expectedArgs:     []string{},
			expectedExecuted: true,
			expectedErr:      false,
		},
		{
			name:             "missing flag value",
			args:             []string{"-name"},
			expectedArgs:     nil,
			expectedExecuted: false, 
			expectedErr:      true, 
			errorContains:    "flag needs an argument",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			var executed bool
			var passedArgs []string

			testCmd := &BaseCommand{
				flags: flag.NewFlagSet("tester", flag.ContinueOnError),
				execute: func(cmd *BaseCommand, args []string) {
					executed = true
					passedArgs = args
				},
			}

			testCmd.flags.String("name", "", "a test of string flag")

			err := testCmd.Init(testCase.args)
			if testCase.expectedErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), testCase.errorContains)
			} else {
				assert.NoError(t, err)
				testCmd.Run()
			}


			assert.Equal(t, testCase.expectedExecuted, executed)

			if testCase.expectedExecuted {
				assert.Equal(t, testCase.expectedArgs, passedArgs)
			}
		})
	}
}
