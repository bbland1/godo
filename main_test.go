package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/bbland1/goDo/cmd"
	"github.com/bbland1/goDo/task"
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

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	exitCode := runAppLogic(&buffer, []string{"main"}, db)
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

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	exitCode := runAppLogic(&buffer, []string{"main", "unknown"}, db)

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

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	exitCode := runAppLogic(&buffer, []string{"main", "help"}, db)

	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	expectedOutput := cmd.UserManual
	output := strings.TrimSpace(buffer.String())
	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestVersionCommand(t *testing.T) {
	var buffer bytes.Buffer

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	exitCode := runAppLogic(&buffer, []string{"main", "version"}, db)

	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	expectedOutput := "goDo v" + cmd.BuildInfo.Version
	output := strings.TrimSpace(buffer.String())
	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestAddCommand(t *testing.T) {
	var buffer bytes.Buffer

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	exitCode := runAppLogic(&buffer, []string{"main", "add"}, db)

	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	expectedOutput := "this is an add command, and this is everything else, []"
	output := strings.TrimSpace(buffer.String())
	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}