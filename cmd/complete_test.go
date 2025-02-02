package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/bbland1/goDo/task"
)

func TestCompleteUsageFlag(t *testing.T) {
	var buffer bytes.Buffer
	var exitCode int

	expectedOutput := CompleteUsage

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	completeCommand := NewCompleteCommand(&buffer, db, &exitCode)

	completeCommand.flags.Usage()

	output := strings.TrimSpace(buffer.String())
	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got %q", expectedOutput, output)
	}
}

func TestCompleteFlag(t *testing.T) {
	var buffer bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	completeCommand := NewCompleteCommand(&buffer, db, &exitCode)

	if completeCommand.flags.Name() != "delete" {
		t.Errorf("NewDeleteCommand flag name = %q, want to be %q", completeCommand.flags.Name(), "delete")
	}
}