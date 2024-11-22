package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/bbland1/goDo/cmd"
)

func TestUsageAndExit(t *testing.T) {
	var buffer bytes.Buffer

	exitCode := usageAndExit(&buffer, "this is a test error message", 3)

	if exitCode != 3 {
		t.Errorf("Exit code of 3 was expected but got %d", exitCode)
	}

	expectedOutput := "this is a test error message"

	output := strings.TrimSpace(buffer.String())

	if !strings.Contains(output, expectedOutput) {
		t.Errorf("Expected error message to contain %q, but got %q", expectedOutput, output)
	}
}

func TestNoArgs(t *testing.T) {
	var buffer bytes.Buffer

	exitCode := runAppLogic(&buffer, []string{"main"})
	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	expectedOutput := cmd.Greeting
	output := strings.TrimSpace(buffer.String())
	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestUnknownCommand(t *testing.T) {
	var buffer bytes.Buffer

	exitCode := runAppLogic(&buffer, []string{"main", "unknown"})

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedOutput := fmt.Sprintf("unknown command passed to goDo: %s", "unknown")
	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestHelpCommand(t *testing.T) {
	var buffer bytes.Buffer
	exitCode := runAppLogic(&buffer, []string{"main", "help"})

	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	expectedOutput := cmd.UserManual
	output := strings.TrimSpace(buffer.String())
	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}
