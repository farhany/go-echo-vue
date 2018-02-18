package models

import (
	"database/sql"

	// following the tutorial for the blank import
	_ "github.com/mattn/go-sqlite3"
)

// Task struct
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TaskCollection is a collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTasks gets all tasks from the database
func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := TaskCollection{}

	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name)

		if err2 != nil {
			panic(err2)
		}

		result.Tasks = append(result.Tasks, task)

	}
	return result
}

// PutTask puts a task into the db
func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"

	// create a prepared sql statement
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	// replace the ? in our prepare statement with name
	result, err2 := stmt.Exec(name)

	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// DeleteTask deletes a task
func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
