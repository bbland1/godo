package task

import (
	"fmt"
	"strings"
	"testing"
	"time"

	_ "modernc.org/sqlite" // Import SQLite driver
)

func TestDatabaseInit(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	query := `SELECT name FROM sqlite_master WHERE type='table' AND name= ?`
	row := db.QueryRow(query, "tasks")

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

	query := `SELECT description, is_completed FROM tasks WHERE description = ?`

	var description string
	var isCompleted bool
	err = db.QueryRow(query, testTask.Description).Scan(&description, &isCompleted)
	if err != nil {
		t.Fatalf("Error in finding the task in the db, %v", err)
	}

	if description != testTask.Description || isCompleted != testTask.IsCompleted {
		t.Errorf("Expected task (%v), got (%v, %v)", testTask, description, isCompleted)
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

	duplicateTestTask := CreateTask(testDescription)

	err = AddTask(db, duplicateTestTask)
	if err == nil {
		t.Errorf("Expected there to be an error when adding a duplicate task but got none")
	}

	if err != nil && !strings.Contains(err.Error(), "UNIQUE constraint failed: tasks.description") {
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
		Description:   "",
		IsCompleted:   false,
		DateAdded:     time.Now(),
		DateCompleted: nil,
	}

	err = AddTask(db, testTask)
	if err == nil {
		t.Errorf("Expected there to be an error when adding a task with an empty Name, but got none")
	}

	if err != nil && !strings.Contains(err.Error(), "description can not be empty") {
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

	err = DeleteTask(db, 1)
	if err != nil {
		t.Errorf("Expected tasked to be deleted and error to be nil, but got %v", err)
	}
	query := `SELECT COUNT(*) FROM tasks WHERE id = ?`

	var count int
	err = db.QueryRow(query, 1).Scan(&count)
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

	nonExistentId := 1

	err = DeleteTask(db, nonExistentId)
	if err == nil {
		t.Errorf("Expected an error when attempting to delete a non-exist tas but got none.")
	}

	expectedErrMsg := fmt.Sprintf("task with id = %d not found", nonExistentId)
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error '%s', but got: %v", expectedErrMsg, err)
	}
}

func TestGetAllTasks(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	testTask1 := CreateTask("test 1")
	testTask2 := CreateTask("test 2")
	testTask3 := CreateTask("test 3")

	err = AddTask(db, testTask1)
	if err != nil {
		t.Fatalf("AddTask testTask1 failed: %v", err)
	}
	err = AddTask(db, testTask2)
	if err != nil {
		t.Fatalf("AddTask testTask2 failed: %v", err)
	}

	err = AddTask(db, testTask3)
	if err != nil {
		t.Fatalf("AddTask testTask3 failed: %v", err)
	}
	
	tasks, err := GetAllTasks(db)
	if err != nil {
		t.Fatalf("GetAllTasks failed: %v", err)
	}

	if len(tasks) != 3 {
		t.Errorf("Expected there to be 3 tasks in DB, got %d", len(tasks))
	}
}

func TestGetATaskById(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	testTask1 := CreateTask("test 1")
	testTask2 := CreateTask("test 2")
	testTask3 := CreateTask("test 3")

	err = AddTask(db, testTask1)
	if err != nil {
		t.Fatalf("AddTask testTask1 failed: %v", err)
	}
	err = AddTask(db, testTask2)
	if err != nil {
		t.Fatalf("AddTask testTask2 failed: %v", err)
	}

	err = AddTask(db, testTask3)
	if err != nil {
		t.Fatalf("AddTask testTask3 failed: %v", err)
	}

	task, err := GetATaskByID(db, 2)
	if err != nil {
		t.Fatalf("GetATaskById failed: %v", err)
	}

	if task.Description != testTask2.Description {
		t.Errorf("Expected the retrieved task to have the name %s, but got %s", testTask2.Description, task.Description)
	}
}

func TestGetATaskByDescription(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	testTask1 := CreateTask("test 1")
	testTask2 := CreateTask("test 2")
	testTask3 := CreateTask("test 3")

	err = AddTask(db, testTask1)
	if err != nil {
		t.Fatalf("AddTask testTask1 failed: %v", err)
	}
	err = AddTask(db, testTask2)
	if err != nil {
		t.Fatalf("AddTask testTask2 failed: %v", err)
	}

	err = AddTask(db, testTask3)
	if err != nil {
		t.Fatalf("AddTask testTask3 failed: %v", err)
	}

	task, err := GetATaskByDescription(db, testTask2.Description)
	if err != nil {
		t.Fatalf("GetATaskById failed: %v", err)
	}

	if task.Description != testTask2.Description {
		t.Errorf("Expected the retrieved task to have the name %s, but got %s", testTask2.Description, task.Description)
	}
}

func TestUpdateTaskComplete(t *testing.T) {
	db, err := InitDatabase(":memory:")
	if err != nil {
		t.Fatalf("InitDatabase failed at creating the db, %v", err)
	}

	defer db.Close()

	testTask := CreateTask("test 2")

	err = AddTask(db, testTask)
	if err != nil {
		t.Fatalf("AddTask testTask failed: %v", err)
	}

	if err := UpdateTaskCompletionStatus(db, testTask.ID, true); err != nil {
		t.Fatalf("UpdateTaskCompletionStatus failed: %v", err)
	}

	query := `SELECT id, name, is_completed FROM tasks WHERE id = ?`

	var description string
	var isCompleted bool
	var dateCompleted time.Time
	err = db.QueryRow(query, testTask.ID).Scan(&description, &isCompleted, &dateCompleted)
	if err != nil {
		t.Fatalf("Error in finding the task in the db, %v", err)
	}

	if description != testTask.Description || isCompleted != testTask.IsCompleted || testTask.DateCompleted == nil {
		t.Errorf("Expected task to be marked as completed and have a date completed value not as nil(%v), got (%v, %v, %v)", testTask, description, isCompleted, dateCompleted)
	}

}
