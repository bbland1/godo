package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/bbland1/goDo/task"
)

func TestDeleteUsageFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	expectedOutput := DeleteUsage

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	deleteCommand := NewDeleteCommand(&bufferOut, &bufferErr, db, &exitCode)

	deleteCommand.flags.Usage()

	output := strings.TrimSpace(bufferOut.String())
	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestDeleteFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	deleteCommand := NewDeleteCommand(&bufferOut, &bufferErr, db, &exitCode)

	if deleteCommand.flags.Name() != "delete" {
		t.Errorf("NewDeleteCommand flag name = %q, want to be %q", deleteCommand.flags.Name(), "delete")
	}
}

func TestDeleteCommandNoArgs(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	deleteCommand := NewDeleteCommand(&bufferOut, &bufferErr, db, &exitCode)

	deleteCommand.execute(deleteCommand, nil)

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedOutput := "an id or task description needs to be passed for deletion to process"
	output := strings.TrimSpace(bufferErr.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestDeleteCommandById(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&bufferOut, &bufferErr, db, &exitCode)

	addCommand.execute(addCommand, []string{"tester"})

	deleteCommand := NewDeleteCommand(&bufferOut, &bufferErr, db, &exitCode)

	deleteCommand.Init([]string{"-id=1"})
	deleteCommand.Run()

	expectedOutput := ""
	output := strings.TrimSpace(bufferOut.String())
	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestDeleteCommandByDescription(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&bufferOut, &bufferErr, db, &exitCode)

	addCommand.execute(addCommand, []string{"tester"})

	deleteCommand := NewDeleteCommand(&bufferOut, &bufferErr, db, &exitCode)

	deleteCommand.Init([]string{"-d=tester"})
	deleteCommand.Run()

	expectedOutput := ""
	output := strings.TrimSpace(bufferOut.String())
	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}
