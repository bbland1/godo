package task

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID            string
	Description          string
	IsCompleted   bool
	DateAdded     time.Time
	DateCompleted *time.Time
}

// CreateTask creates and new task with the passed description.
func CreateTask(description string) *Task {
	return &Task{
		ID: uuid.New().String(),
		Description: description,
		IsCompleted: false,
		DateAdded: time.Now(),
		DateCompleted: nil,
	}
}

// MarkTaskCompleted will updated the completed values of a task to the now time and true.
func (t *Task) MarkTaskCompleted() {
	timeNow := time.Now()

	t.DateCompleted = &timeNow
	t.IsCompleted = true
}

// UpdateTaskName uses the passed string value to replace the task description with the new value
func (t *Task) UpdateTaskName(newDescription string) {
	t.Description = newDescription
}