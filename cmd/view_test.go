package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"text/tabwriter"

	"github.com/bbland1/goDo/task"
)

func TestViewUsageFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	expectedOutput := ViewUsage

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	viewCommand := NewViewCommand(&bufferOut, &bufferErr, db, &exitCode)

	viewCommand.flags.Usage()

	output := strings.TrimSpace(bufferOut.String())
	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestViewCommandFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	db, err := task.InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	viewCommand := NewViewCommand(&bufferOut, &bufferErr, db, &exitCode)

	if viewCommand.flags.Name() != "view" {
		t.Errorf("NewViewCommand flag name = %q, want to be %q", viewCommand.flags.Name(), "view")
	}
}

func TestViewCommandNoArgs(t *testing.T) {
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

	viewCommand := NewViewCommand(&bufferOut, &bufferErr, db, &exitCode)

	viewCommand.execute(viewCommand, nil)

	var bufExpOut bytes.Buffer

	tasks, err := task.GetAllTasks(db)
	if err != nil {
		t.Errorf("There has been an issue getting the tasks from the DB, %s", err)
	}
	tw := tabwriter.NewWriter(&bufExpOut, 0, 8, 2, ' ', 0)
	fmt.Fprintln(tw, "ID\tTask Description\tStatus")

	for _, storedTask := range tasks {
		fmt.Fprintf(tw, "%d\t%s\t%t\n", storedTask.ID, storedTask.Description, storedTask.IsCompleted)
	}

	tw.Flush()

	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	output := strings.TrimSpace(bufferOut.String())
	expectedOutput := strings.TrimSpace(bufExpOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestViewCommandBothArgs(t *testing.T) {
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
	addCommand.execute(addCommand, []string{"tester 1"})

	viewCommand := NewViewCommand(&bufferOut, &bufferErr, db, &exitCode)

	viewCommand.Init([]string{"-id=1"})
	viewCommand.Run()

	var bufExpOut bytes.Buffer

	tw := tabwriter.NewWriter(&bufExpOut, 0, 8, 2, ' ', 0)
	fmt.Fprintln(tw, "ID\tTask Description\tStatus")

	fmt.Fprintf(tw, "%d\t%s\t%t\n", 1, "tester", false)

	tw.Flush()

	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	output := strings.TrimSpace(bufferOut.String())
	expectedOutput := strings.TrimSpace(bufExpOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}
