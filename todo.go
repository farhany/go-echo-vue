package main

import (
	"database/sql"
	"fmt"
	"go-echo-vue/handlers"
	"os"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDB("storage.db")
	migrate(db)

	e := echo.New()

	e.File("/", "public/index.html")

	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTasks(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Print("Port not in env, setting it to 8090")
		port = "8090"
	}

	e.Logger.Fatal(e.Start(":" + port))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(191) NOT NULL
	);
	`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
