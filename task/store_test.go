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

	query := `SELECT name FROM sqlite_master WHERE type=table AND name=tasks`
	row := db.QueryRow(query)

	var tableName string
	if err := row.Scan(&tableName); err != nil {
		t.Fatalf("Failed to query database schema: %v", err)
	}

	if tableName != "tasks" {
		t.Errorf("Expected table 'tasks', but got: %v", tableName)
	}
}

func TestGetAllTasks(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()
}

func TestGetATask(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()
}

func TestUpdateTaskName(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()
}

func TestUpdateTaskComplete(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()
}

func TestDeleteTask(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()
}

func TestAddTask(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	testName := "tester name"

	testTask := CreateTask(testName)

	err = AddTask(db, testTask)
	if err != nil {
		t.Fatalf("AddTask failed: %v", err)
	}

	query := `SELECT id, name, is_completed FROM tasks WHERE id = ?`
	row := db.QueryRow(query)

	var id, name string
	var isCompleted bool
	err = row.Scan(&id, &name, &isCompleted)

	if id != testTask.ID || name != testTask.Name || isCompleted != testTask.IsCompleted {
		t.Errorf("Expected task (%v), got (%v, %v, %v)", testTask, id, name, isCompleted)
	}
}

func TestAddingDuplicateTask(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	testName := "tester name"

	testTask := CreateTask(testName)

	err = AddTask(db, testTask)
	if err != nil {
		t.Fatalf("AddTask failed: %v", err)
	}

	duplicateTestTask := &Task{
		ID:            testTask.ID,
		Name:          "duplication",
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
		Name:          "duplication",
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
		Name:          "",
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
