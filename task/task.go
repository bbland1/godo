package task

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID            string
	Name          string
	IsCompleted   bool
	DateAdded     time.Time
	DateCompleted *time.Time
}

func CreateTask(name string) *Task {
	return &Task{
		ID: uuid.New().String(),
		Name: name,
		IsCompleted: false,
		DateAdded: time.Now(),
		DateCompleted: nil,
	}
}