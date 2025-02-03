package task

import (
	"time"
)

type Task struct {
	ID            int64 // sql lite returns task id as int64 so have currently made type match across project
	Description   string
	IsCompleted   bool
	DateAdded     time.Time
	DateCompleted *time.Time
}

// CreateTask creates and new task with the passed description.
func CreateTask(description string) *Task {
	return &Task{
		Description:   description,
		IsCompleted:   false,
		DateAdded:     time.Now().UTC(),
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
