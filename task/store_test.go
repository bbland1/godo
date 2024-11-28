package task

import (
	"database/sql"
	"testing"
	
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

}

func TestGetATask(t *testing.T) {

}

func TestUpdateTaskName(t *testing.T) {

}

func TestUpdateTaskComplete(t *testing.T) {

}

func TestDeleteTask(t *testing.T) {

}

func TestAddTask(t *testing.T) {

}

func TestAddingDuplicateTask(t *testing.T) {

}

func TestAddInvalidTask(t *testing.T) {
	
}