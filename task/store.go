package task

import (
	"cmp"
	"database/sql"
	"fmt"

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
		description TEXT NOT NULL,
		is_completed INTEGER CHECK(is_completed IN (0,1)) NOT NULL,
		date_added TEXT NOT NULL,
		date_completed TEXT
	)`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create tasks table: %w", err)
	}
	return db, nil
}

func AddTask(db *sql.DB, task *Task) error {
	return nil
}

func GetAllTasks(db *sql.DB){
}

// func GetATask(db *sql.DB, taskID *Task.ID) (*Task, error) {
// 	return task, nil
// }