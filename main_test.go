package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"testing"
	"text/tabwriter"

	"github.com/bbland1/goDo/cmd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunAppLogic(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		setup        func(out *bytes.Buffer, err *bytes.Buffer, exit *int)
		expectedCode int
		expectedOut  string
		expectedErr  string
	}{
		{
			name:         "no command args passed",
			args:         []string{"godo"},
			expectedOut:  cmd.Greeting,
			expectedCode: 0,
		},
		{
			name:         "unregistered command passed",
			args:         []string{"godo", "unknown"},
			expectedCode: 1,
			expectedErr:  "Unknown command: unknown",
		},
		{
			name: "pass a valid command",
			args: []string{"godo", "help"},
			setup: func(out, err *bytes.Buffer, exit *int) {
				cmd.RegisterCommand(cmd.NewHelpCommand(out, err, exit))
			},
			expectedCode: 0,
			expectedOut: func() string {
				var tempBuf bytes.Buffer
				GenerateTestUserManual([]cmd.Command{
					cmd.NewHelpCommand(&tempBuf, &bytes.Buffer{}, new(int)),
				}, &tempBuf)
				return strings.TrimSpace(tempBuf.String())
			}(),
		},
		{
			name: "command init fails",
			args: []string{"goDo", "fail"},
			setup: func(out, err *bytes.Buffer, exit *int) {
				failure := &failingCommand{
					name:        "fail",
					description: "a fail",
				}
				cmd.RegisterCommand(failure)
			},
			expectedCode: 1,
			expectedErr:  "Error initializing command: fail",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			cmd.ClearCommandRegistry()

			var bufferOut bytes.Buffer
			var bufferErr bytes.Buffer
			var exitCode int

			if testCase.setup != nil {
				testCase.setup(&bufferOut, &bufferErr, &exitCode)
			}

			exitCode = runAppLogic(&bufferOut, &bufferErr, testCase.args)

			require.Equal(t, testCase.expectedCode, exitCode, "exit codes do not match")

			if testCase.expectedErr != "" {

				assert.Equal(t, testCase.expectedErr, strings.TrimSpace(bufferErr.String()), "unexpected stderr message")
			}

			if testCase.expectedOut != "" {

				assert.Equal(t, testCase.expectedOut, strings.TrimSpace(bufferOut.String()), "unexpected stdout message")
			}
		})
	}
}

type failingCommand struct {
	name        string
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

func GenerateTestUserManual(commands []cmd.Command, out *bytes.Buffer) {

	tw := tabwriter.NewWriter(out, 0, 8, 2, ' ', 0)

	fmt.Fprintln(tw, "Usage:\n  goDo [command] [options]")

	fmt.Fprintln(tw, "\nOptions:")
	fmt.Fprintln(tw, "  -h\tShow more information about a command")
	fmt.Fprintln(tw, "  -verbose\tPrint detailed output when available")
	fmt.Fprintln(tw, "\nCommands:")

	tempRegisteredCommands := map[string]cmd.Command{}
	for _, c := range commands {
		tempRegisteredCommands[c.GetName()] = c
	}

	// Sort command names for consistent display
	var names []string
	for name := range tempRegisteredCommands {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		cmd := tempRegisteredCommands[name]
		fmt.Fprintf(tw, "  %s\t- %s\n", cmd.GetName(), cmd.GetDescription())
	}

	tw.Flush()
}
