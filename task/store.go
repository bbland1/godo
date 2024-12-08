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
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL UNIQUE CHECK(description != ''),
		is_completed INTEGER CHECK(is_completed IN (0,1)) DEFAULT 0 NOT NULL,
		date_added DATETIME NOT NULL,
		date_completed DATETIME
	)`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create tasks table: %w", err)
	}
	return db, nil
}

// AddTask will add the passed task into the sqlite DB.
func AddTask(db *sql.DB, task *Task) error {
	addRowQuery := `INSERT INTO tasks (description, is_completed, date_added, date_completed) VALUES (?, ?, ?, ?)`

	if task.Description == "" {
		return fmt.Errorf("description can not be empty")
	}

	completedValue := 0
	if task.IsCompleted {
		completedValue = 1
	}

	_, err := db.Exec(addRowQuery, task.Description, completedValue, task.DateAdded, task.DateCompleted)
	if err != nil {
		return fmt.Errorf("error when adding task to the db: %w", err)
	}

	return nil
}

func GetAllTasks(db *sql.DB) ([]Task, error) {
	getAllTasksQuery := `SELECT id, description, is_completed, date_added, date_completed FROM tasks`

	rows, err := db.Query(getAllTasksQuery)
	if err != nil {
		return nil, fmt.Errorf("error when adding task to the db: %w", err)
	}

	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Description, &task.IsCompleted, &task.DateAdded, &task.DateCompleted);
		if err != nil {
			return nil, fmt.Errorf("error scanning task: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}
	
	return tasks, nil
}

func GetATaskByID(db *sql.DB, id string) (*Task, error) {
	return nil, nil
}

func UpdateTaskCompletionStatus(db *sql.DB, id string, isCompleted bool) error {
	return nil
}

// DeleteTask will remove the task from the sqlite DB.
func DeleteTask(db *sql.DB, id int) error {
	deleteTaskQuery := `DELETE FROM tasks WHERE id = ?`

	result, err := db.Exec(deleteTaskQuery, id)
	if err != nil {
		return fmt.Errorf("error deleting task (id = %d)from the db: %w", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected in the delete in getting : %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task with id = %d not found", id)
	}
	return nil
}
