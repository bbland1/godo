package task

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateTask(t *testing.T) {
	description := "a sample task"

	task := CreateTask(description)

	if task.ID == "" {
		t.Errorf("expected ID to be generated, but got an empty string")
	}

	if _, err := uuid.Parse(task.ID); err != nil {
		t.Errorf("expected a valid UUID for ID, but got: %v, error: %v", task.ID, err)
	}

	if task.Description != description {
		t.Errorf("expected name of the task to be %s, and got %s", description, task.Description)
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

func TestMarkingTaskAsComplete(t *testing.T) {
	description := "a sample task"

	task := CreateTask(description)

	task.MarkTaskCompleted()

	timeNow := time.Now()

	if task.DateCompleted.Before(timeNow.Add(-time.Second)) {
		t.Errorf("when task is completed date completed should be nil since now done, got %v", task.DateCompleted)
	}

	if task.IsCompleted != true {
		t.Errorf("completion status should be %t, but got %t", true, task.IsCompleted)
	}
}

func TestUpdatingTaskName(t *testing.T) {
	description := "a sample task"
	newDescription := "a new name"

	task := CreateTask(description)

	if task.Description != description {
		t.Errorf("there was an error in creating the task got = %s, want = %s", task.Description, description)
	}

	task.UpdateTaskName(newDescription)

	if task.Description != newDescription {
		t.Errorf("the update to the task name should have happened, got = %s, want %s", task.Description, newDescription)
	}
}