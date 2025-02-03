package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/bbland1/goDo/cmd"
	"github.com/bbland1/goDo/task"
)

// todo: need to review the tests to make sure they are effective

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

func TestNoCommandArgs(t *testing.T) {
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

func TestAddCommandNoArgs(t *testing.T) {
	var buffer bytes.Buffer

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	exitCode := runAppLogic(&buffer, []string{"main", "add"}, db)

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedOutput := "a description string needs to be passed to add a task"
	output := strings.TrimSpace(buffer.String())
	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestAddCommandArgs(t *testing.T) {
	var buffer bytes.Buffer

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	exitCode := runAppLogic(&buffer, []string{"main", "add", "tester"}, db)

	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	expectedOutput := ""
	output := strings.TrimSpace(buffer.String())
	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestDeleteCommandNoArgs(t *testing.T) {
	var buffer bytes.Buffer

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	exitCode := runAppLogic(&buffer, []string{"main", "add", "tester"}, db)

	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	exitCode = runAppLogic(&buffer, []string{"main", "delete", ""}, db)

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedOutput := "an id or task description needs to be passed for deletion to process"
	output := strings.TrimSpace(buffer.String())
	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestDeleteCommandArgs(t *testing.T) {
	var buffer bytes.Buffer

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()
	exitCode := runAppLogic(&buffer, []string{"main", "add", "tester"}, db)

	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	exitCode = runAppLogic(&buffer, []string{"main", "delete", "-id=1"}, db)

	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	expectedOutput := ""
	output := strings.TrimSpace(buffer.String())
	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestStatusCommandArgs(t *testing.T) {
	var buffer bytes.Buffer

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()
	exitCode := runAppLogic(&buffer, []string{"main", "add", "tester"}, db)

	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	exitCode = runAppLogic(&buffer, []string{"main", "status", "-id=1", "true"}, db)

	// todo: would this need to check that the status changed if that is tested in the status_test

	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	expectedOutput := ""
	output := strings.TrimSpace(buffer.String())
	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestStatusCommandNoArgs(t *testing.T) {
	var buffer bytes.Buffer

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	exitCode := runAppLogic(&buffer, []string{"main", "add", "tester"}, db)

	if exitCode != 0 {
		t.Errorf("Exit code of 0 was expected but got %d", exitCode)
	}

	exitCode = runAppLogic(&buffer, []string{"main", "status", ""}, db)

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedOutput := "an id or task description needs to be passed to mark something as complete"
	output := strings.TrimSpace(buffer.String())
	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}