package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/bbland1/goDo/task"
)

var exitCode int

func TestAddUsageFlag(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := AddUsage

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&buffer, db, &exitCode)

	addCommand.flags.Usage()

	output := strings.TrimSpace(buffer.String())
	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestAddCommandFlag(t *testing.T) {
	var buffer bytes.Buffer

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&buffer, db, &exitCode)

	if addCommand.flags.Name() != "add" {
		t.Errorf("NewAddCommand flag name = %q, want to be %q", addCommand.flags.Name(), "add")
	}
}

func TestAddCommandNoArgs(t *testing.T) {
	var buffer bytes.Buffer

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&buffer, db, &exitCode)

	addCommand.Execute(addCommand, nil)

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedOutput := "a description string needs to be passed to add a task"
	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestAddCommandWithDescription(t *testing.T) {
	var buffer bytes.Buffer

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&buffer, db, &exitCode)

	addCommand.Execute(addCommand, []string{"tester"})

	expectedOutput := ""
	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestAddToDBError(t *testing.T) {
	var buffer bytes.Buffer

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&buffer, db, &exitCode)

	addCommand.Execute(addCommand, []string{" "})

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedOutput := "database error:"
	output := strings.TrimSpace(buffer.String())

	if !strings.Contains(output, expectedOutput) {
		t.Errorf("Expected output to contain: %q, got: %q", expectedOutput, output)
	}
}
