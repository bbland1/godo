package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/bbland1/goDo/task"
)

func TestDeleteUsageFlag(t *testing.T) {
	var buffer bytes.Buffer
	var exitCode int

	expectedOutput := DeleteUsage

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	deleteCommand := NewDeleteCommand(&buffer, db, &exitCode)

	deleteCommand.flags.Usage()

	output := strings.TrimSpace(buffer.String())
	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestDeleteFlag(t *testing.T) {
	var buffer bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	deleteCommand := NewDeleteCommand(&buffer, db, &exitCode)

	if deleteCommand.flags.Name() != "delete" {
		t.Errorf("NewDeleteCommand flag name = %q, want to be %q", deleteCommand.flags.Name(), "delete")
	}
}
