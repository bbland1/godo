package task

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateTask(t *testing.T) {
	name := "a sample task"

	task := CreateTask(name)

	if task.ID == "" {
		t.Errorf("expected ID to be generated, but got an empty string")
	}

	if _, err := uuid.Parse(task.ID); err != nil {
		t.Errorf("expected a valid UUID for ID, but got: %v, error: %v", task.ID, err)
	}

	if task.Name != name {
		t.Errorf("expected name of the task to be %s, and got %s", name, task.Name)
	}

	if task.IsCompleted != false {
		t.Errorf("completion status should be %t, but got %t", false, task.IsCompleted)
	}

	timeNow := time.Now()
	if task.DateAdded.Before(timeNow.Add(-time.Second)) {
		t.Errorf("expected DateAdded to be close to now, got %v", task.DateAdded)
	}

	if task.DateCompleted != nil {
		t.Errorf("when task is completed date completed should be nil since now done, got %v", task.DateCompleted)
	}
}
