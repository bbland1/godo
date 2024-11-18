package main

import (
	"bytes"
	"fmt"
	"strings"

	// "os"
	"os/exec"
	"testing"

	"github.com/bbland1/goDo/cmd"
)

func TestNoArgs(t *testing.T) {
	testCmd := exec.Command("go", "run", "main.go")
	var stdout bytes.Buffer
	testCmd.Stdout = &stdout

	err := testCmd.Run()
	if err != nil {
		t.Errorf("Command execution failed: %v", err)
	}

	expectedOutput := cmd.Greeting
	output := strings.TrimSpace(stdout.String())
	if output != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, output)
	}
}

func TestUnknownCommand(t *testing.T) {
	testCmd := exec.Command("go", "run", "main.go", "unknown")
	var stderr bytes.Buffer
	testCmd.Stderr = &stderr

	err := testCmd.Run()

	if exitErr, ok := err.(*exec.ExitError); ok {
		if exitErr.ExitCode() != 1 {
			t.Errorf("Expected exit code 1, but got: %d", exitErr.ExitCode())
		}
	} else if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else {
		t.Errorf("Expected non-zero exit code for unknown command, but command succeeded")
	}

	expectedError := fmt.Sprintf("unknown command passed to goDo: %s\n", "unknown")
	actualError := stderr.String()
	if !strings.Contains(actualError, expectedError) {
		t.Errorf("Expected output: %q, but got: %q", expectedError, stderr.String())
	}
}
