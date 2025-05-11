package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/bbland1/goDo/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddToDBError(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&bufferOut, &bufferErr, db, &exitCode)

	addCommand.execute(addCommand, []string{" "})

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedOutput := "database error:"
	output := strings.TrimSpace(bufferErr.String())

	if !strings.Contains(output, expectedOutput) {
		t.Errorf("Expected output to contain: %q, got: %q", expectedOutput, output)
	}
}

func TestAddCommand(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		setup        func(cmd *BaseCommand)
		expectedCode int
		expectedOut  string
		expectedErr  string
		useContains  bool
	}{
		{
			name: "usage flag set",
			setup: func(cmd *BaseCommand) {
				cmd.flags.Usage()
			},
			expectedCode: 0,
			expectedOut:  AddUsage,
		},
		{
			name:         "add command with no args",
			expectedCode: 1,
			expectedErr:  "a description string needs to be passed to add a task",
		},
		{
			name:         "valid description",
			args:         []string{"tester"},
			expectedCode: 0,
			expectedOut:  "",
		},
		{
			name:         "invalid description (whitespace only)",
			args:         []string{" "},
			expectedCode: 1,
			expectedOut:  "database error:",
			useContains:  true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			var bufferOut bytes.Buffer
			var bufferErr bytes.Buffer
			var exitCode int

			db, err := task.InitDatabase(":memory:")
			require.NoError(t, err)
			defer db.Close()

			addCommand := NewAddCommand(&bufferOut, &bufferErr, db, &exitCode)

			if testCase.setup != nil {
				testCase.setup(addCommand)
			} else {
				addCommand.execute(addCommand, testCase.args)
			}

			require.Equal(t, testCase.expectedCode, exitCode, "exit codes do not match")

			if testCase.expectedErr != "" {
				errOutput := strings.TrimSpace(bufferErr.String())

				if testCase.useContains {
					assert.Contains(t,  errOutput, testCase.expectedErr, "unexpected stderr message")
				}
				assert.Equal(t, testCase.expectedErr, errOutput, "unexpected stderr message")
			}

			if testCase.expectedOut != "" {

				assert.Equal(t, testCase.expectedOut, strings.TrimSpace(bufferOut.String()), "unexpected stdout message")
			}
		})
	}
}
