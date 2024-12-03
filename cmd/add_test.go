package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/bbland1/goDo/task"
)

func TestAddUsageFlag(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := AddUsage

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&buffer, db)

	addCommand.flags.Usage()

	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}