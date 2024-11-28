package task

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "modernc.org/sqlite" // Import SQLite driver
)

func TestDatabaseInit(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	query := `SELECT name FROM sqlite_master WHERE type='table' AND name='tasks'`
	row := db.QueryRow(query)

	var tableName string
	if err := row.Scan(&tableName); err != nil {
		t.Fatalf("Failed to query database schema: %v", err)
	}

	if tableName != "tasks" {
		t.Errorf("Expected table 'tasks', but got: %v", tableName)
	}
}

func TestAddTask(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	testDescription := "tester name"

	testTask := CreateTask(testDescription)

	err = AddTask(db, testTask)
	if err != nil {
		t.Fatalf("AddTask failed: %v", err)
	}

	query := `SELECT id, name, is_completed FROM tasks WHERE id = ?`
	row := db.QueryRow(query)

	var id, description string
	var isCompleted bool
	err = row.Scan(&id, &description, &isCompleted)

	if id != testTask.ID || description != testTask.Description || isCompleted != testTask.IsCompleted {
		t.Errorf("Expected task (%v), got (%v, %v, %v)", testTask, id, description, isCompleted)
	}
}

func TestAddingDuplicateTask(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	testDescription := "tester name"

	testTask := CreateTask(testDescription)

	err = AddTask(db, testTask)
	if err != nil {
		t.Fatalf("AddTask failed: %v", err)
	}

	duplicateTestTask := &Task{
		ID:            testTask.ID,
		Description:   "duplication",
		IsCompleted:   false,
		DateAdded:     time.Now(),
		DateCompleted: nil,
	}

	err = AddTask(db, duplicateTestTask)
	if err == nil {
		t.Errorf("Expected there to be an error when adding a duplicate task but got none")
	}

	if err != nil && !strings.Contains(err.Error(), "UNIQUE constraint failed") {
		t.Errorf("Expected unique constraint error, got: %v", err)
	}
}

func TestAddInvalidIDTask(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	testTask := &Task{
		ID:            "7",
		Description:   "duplication",
		IsCompleted:   false,
		DateAdded:     time.Now(),
		DateCompleted: nil,
	}

	err = AddTask(db, testTask)
	if err == nil {
		t.Errorf("Expected there to be an error when adding a task with an invalid ID (non-UUID), but got none")
	}

	if err != nil && strings.Contains(err.Error(), "invalid UUID") {
		t.Errorf("Expected unique constraint error, got: %v", err)
	}
}

func TestAddEmptyNameTask(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	testTask := &Task{
		ID:            uuid.New().String(),
		Description:   "",
		IsCompleted:   false,
		DateAdded:     time.Now(),
		DateCompleted: nil,
	}

	err = AddTask(db, testTask)
	if err == nil {
		t.Errorf("Expected there to be an error when adding a task with an empty Name, but got none")
	}

	if err != nil && strings.Contains(err.Error(), "CHECK constraint failed") {
		t.Errorf("Expected CHECK constraint error for Name, got: %v", err)
	}
}

func TestDeleteTask(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	testDescription := "tester name"

	testTask := CreateTask(testDescription)

	err = AddTask(db, testTask)
	if err != nil {
		t.Fatalf("AddTask failed: %v", err)
	}

	err = DeleteTask(db, testTask.ID)
	if err != nil {
		t.Errorf("Expected tasked to be deleted and error to be nil, but got %v", err)
	}

	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM tasks WHERE id = ?`, testTask.ID).Scan(&count)
	if err != nil {
		t.Fatalf("Error querying task after deletion: %v", err)
	}
	if count != 0 {
		t.Errorf("Expected task to be deleted, but found %d tasks", count)
	}
}

func TestDeleteNonExistentTask(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	err = DeleteTask(db, "no-task-here")
	if err == nil {
		t.Errorf("Expected an error when attempting to delete a non-exist tas but got none.")
	}

	if !strings.Contains(err.Error(), "task with ID nonexistent-task-id not found") {
		t.Errorf("Expected 'task not found' error, but got: %v", err)
	}
}

func TestGetAllTasks(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	CreateTask("test 1")
	CreateTask("test 2")
	CreateTask("test 3")

	tasks, err := GetAllTasks(db)
	if err != nil {
		t.Fatalf("GetAllTasks failed: %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected there to be 3 tasks in DB, got %d", len(tasks))
	}
}

func TestGetATask(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	CreateTask("test 1")
	testTask2 := CreateTask("test 2")
	CreateTask("test 3")

	task, err := GetATask(db, testTask2.ID)
	if err != nil {
		t.Fatalf("GetATask failed: %v", err)
	}

	if task.Description != testTask2.Description {
		t.Errorf("Expected the retrieved task to have the name %s, but got %s", testTask2.Description, task.Description)
	}
}

func TestUpdateTaskName(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	CreateTask("test 1")
	testTask2 := CreateTask("test 2")
	CreateTask("test 3")
}

func TestUpdateTaskComplete(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	CreateTask("test 1")
	testTask2 := CreateTask("test 2")
	CreateTask("test 3")
}
