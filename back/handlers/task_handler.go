package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pansani/interjiajs-golang-example/database"
	"github.com/pansani/interjiajs-golang-example/models"
)

func GetTasks(c echo.Context) error {
	var tasks []models.Task
	database.DB.Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}

func CreateTask(c echo.Context) error {
	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	database.DB.Create(task)
	return c.JSON(http.StatusCreated, task)
}
