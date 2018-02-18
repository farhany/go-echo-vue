package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"go-echo-vue/models"

	"github.com/labstack/echo"
)

// H is for JSON
type H map[string]interface{}

// GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// PutTasks endpoint
func PutTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task models.Task

		// map incoming json body to the new task
		c.Bind(task)

		// ad a task using our new model
		id, err := models.PutTask(db, task.Name)

		// return a json response if successfull
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		}

		return err
	}
}

// DeleteTask endpoint
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		// use our new model to delete a task
		_, err := models.DeleteTask(db, id)

		// return a json response on success
		if err == nil {

			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		}

		return err

	}
}
