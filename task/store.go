package task

import (
	"cmp"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

const dbFile = "goDo.db"

// InitDatabase initializes the SQLite DB for the app.
func InitDatabase(dbSource string) (*sql.DB, error) {

	dbSource = cmp.Or(dbSource, dbFile)

	db, err := sql.Open("sqlite", dbSource)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		description TEXT NOT NULL CHECK(description != ''),
		is_completed INTEGER CHECK(is_completed IN (0,1)) DEFAULT 0 NOT NULL,
		date_added TEXT NOT NULL,
		date_completed TEXT
	)`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create tasks table: %w", err)
	}
	return db, nil
}

// AddTask will add the passed task into the sqlite DB.
func AddTask(db *sql.DB, task *Task) error {
	addRowQuery := `INSERT INTO tasks (id, description, is_completed, date_added, date_completed) VALUES (?, ?, ?, ?, ?)`

	if _, err := uuid.Parse(task.ID); err != nil {
		return fmt.Errorf("invalid UUID: %v", err)
	}

	if task.Description == "" {
		return fmt.Errorf("description can not be empty")
	}

	completedValue := 0
	if task.IsCompleted {
		completedValue = 1
	}

	_, err := db.Exec(addRowQuery, task.ID, task.Description, completedValue, task.DateAdded, task.DateCompleted)
	if err != nil {
		return fmt.Errorf("error when adding task to the db: %w", err)
	}

	return nil
}

func GetAllTasks(db *sql.DB) ([]Task, error) {
	return nil, nil
}

func GetATaskByID(db *sql.DB, id string) (*Task, error) {
	return nil, nil
}

func UpdateTaskCompletionStatus(db *sql.DB, id string, isCompleted bool) error {
	return nil
}

func DeleteTask(db *sql.DB, id string) error {
	return nil
}
