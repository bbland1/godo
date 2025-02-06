package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/bbland1/goDo/task"
)

func TestEditUsageFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	expectedOutput := EditUsage

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	editCommand := NewEditCommand(&bufferOut, &bufferErr, db, &exitCode)

	editCommand.flags.Usage()

	output := strings.TrimSpace(bufferOut.String())
	if exitCode != 0 {
		t.Errorf("Expected exit code to be 0, got: %d", &exitCode)
	}

	if output != expectedOutput {
		t.Errorf("Expected output: %q, go: %q", expectedOutput, output)
	}
}

func TestEditFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	editCommand := NewEditCommand(&bufferOut, &bufferErr, db, &exitCode)

	if editCommand.flags.Name() != "edit" {
		t.Errorf("NewStatusCommand flag name = %q, want to be %q", editCommand.flags.Name(), "edit")
	}
}

func TestEditCommandNoArgs(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	editCommand := NewEditCommand(&bufferOut, &bufferErr, db, &exitCode)

	editCommand.execute(editCommand, nil)

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedOutput := "an id or task description needs to be passed to edit a task"
	output := strings.TrimSpace(bufferErr.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestEditCommandById(t *testing.T) {
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

	editCommand := NewEditCommand(&bufferOut, &bufferErr, db, &exitCode)

	editCommand.Init([]string{"-id=1", "a change"})
	editCommand.Run()

	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	task, err := task.GetATaskByID(db, 1)
	if err != nil {
		t.Errorf("Error get the task from db to check the status change %v", err)
	}

	expectedOutput := "a change"
	output := task.Description

	if output != expectedOutput {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, output)
	}
}

func TestEditNoEdit(t *testing.T) {
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

	editCommand := NewEditCommand(&bufferOut, &bufferErr, db, &exitCode)

	editCommand.Init([]string{"-id=1", ""})
	editCommand.Run()

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedOutput := "a status needs to be passed"
	output := strings.TrimSpace(bufferErr.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, output)
	}
}

func TestEditBadId(t *testing.T) {
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

	editCommand := NewEditCommand(&bufferOut, &bufferErr, db, &exitCode)

	editCommand.Init([]string{"-id=t", "a change"})
	editCommand.Run()

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedOutput := "an int needs to be passed for the id"
	output := strings.TrimSpace(bufferErr.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, output)
	}
}

func TestEditCommandByDescription(t *testing.T) {
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

	editCommand := NewEditCommand(&bufferOut, &bufferErr, db, &exitCode)

	editCommand.Init([]string{"-d=tester", "a change"})
	editCommand.Run()

	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	task, err := task.GetATaskByID(db, 1)
	if err != nil {
		t.Errorf("Error get the task from db to check the status change %v", err)
	}

	expectedOutput := "a change"
	output := task.Description

	if output != expectedOutput {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, output)
	}
}

func TestEditBadDescription(t *testing.T) {
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

	editCommand := NewEditCommand(&bufferOut, &bufferErr, db, &exitCode)

	editCommand.Init([]string{"-d=&", "a change"})
	editCommand.Run()

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedOutput := "a task with that description wasn't found"
	output := strings.TrimSpace(bufferErr.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, output)
	}
}