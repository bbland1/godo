package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/bbland1/goDo/task"
)

func TestStatusUsageFlag(t *testing.T) {
	var buffer bytes.Buffer
	var exitCode int

	expectedOutput := StatusUsage

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	statusCommand := NewStatusCommand(&buffer, db, &exitCode)

	statusCommand.flags.Usage()

	output := strings.TrimSpace(buffer.String())
	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got %q", expectedOutput, output)
	}
}

func TestStatusFlag(t *testing.T) {
	var buffer bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	statusCommand := NewStatusCommand(&buffer, db, &exitCode)

	if statusCommand.flags.Name() != "status" {
		t.Errorf("NewStatusCommand flag name = %q, want to be %q", statusCommand.flags.Name(), "complete")
	}
}

func TestStatusCommandNoArgs(t *testing.T) {
	var buffer bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	statusCommand := NewStatusCommand(&buffer, db, &exitCode)

	statusCommand.Execute(statusCommand, nil)

	if exitCode != 1 {
		t.Errorf("Exit code of 1 was expected but got %d", exitCode)
	}

	expectedOutput := "an id or task description needs to be passed to mark something as complete"
	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestStatusCommandById(t *testing.T) {
	var buffer bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&buffer, db, &exitCode)

	addCommand.Execute(addCommand, []string{"tester"})

	statusCommand := NewStatusCommand(&buffer, db, &exitCode)

	statusCommand.Init([]string{"-id=1", "true"})
	statusCommand.Run()

	task, err := task.GetATaskByID(db, 1)
	if err != nil {
		t.Errorf("Error get the task from db to check the status change %v", err)
	}

	expectedOutput := true
	output := task.IsCompleted

	if output != expectedOutput {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, output)
	}
}

func TestStatusNoStatus(t *testing.T) {
	var buffer bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&buffer, db, &exitCode)

	addCommand.Execute(addCommand, []string{"tester"})

	statusCommand := NewStatusCommand(&buffer, db, &exitCode)

	statusCommand.Init([]string{"-id=1", ""})
	statusCommand.Run()

	expectedOutput := "a status needs to be passed"
	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, output)
	}
}

func TestStatusBadStatus(t *testing.T) {
	var buffer bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&buffer, db, &exitCode)

	addCommand.Execute(addCommand, []string{"tester"})

	statusCommand := NewStatusCommand(&buffer, db, &exitCode)

	statusCommand.Init([]string{"-id=1", "hello"})
	statusCommand.Run()

	expectedOutput := "status has to be 'true' or 'false' to update"
	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, output)
	}
}

func TestStatusBadId(t *testing.T) {
	var buffer bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&buffer, db, &exitCode)

	addCommand.Execute(addCommand, []string{"tester"})

	statusCommand := NewStatusCommand(&buffer, db, &exitCode)

	statusCommand.Init([]string{"-id=t", "true"})
	statusCommand.Run()

	expectedOutput := "an int needs to be passed for the id"
	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, output)
	}
}

func TestStatusCommandByDescription(t *testing.T) {
	var buffer bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&buffer, db, &exitCode)

	addCommand.Execute(addCommand, []string{"tester"})

	statusCommand := NewStatusCommand(&buffer, db, &exitCode)

	statusCommand.Init([]string{"-d=tester", "true"})
	statusCommand.Run()

	task, err := task.GetATaskByID(db, 1)
	if err != nil {
		t.Errorf("Error get the task from db to check the status change %v", err)
	}

	expectedOutput := true
	output := task.IsCompleted

	if output != expectedOutput {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, output)
	}
}

func TestStatusBadDescription(t *testing.T) {
	var buffer bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	addCommand := NewAddCommand(&buffer, db, &exitCode)

	addCommand.Execute(addCommand, []string{"tester"})

	statusCommand := NewStatusCommand(&buffer, db, &exitCode)

	statusCommand.Init([]string{"-d=&", "true"})
	statusCommand.Run()

	expectedOutput := "a task with that description wasn't found"
	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, output)
	}
}