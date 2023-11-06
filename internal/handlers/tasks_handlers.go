package handlers

import (
	"apple-reminder_backend/db"
	"apple-reminder_backend/internal/helpers"
	"apple-reminder_backend/internal/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	AllTasks, err := helpers.GetALLTask_help()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, AllTasks)
}

func GetTaskByID(c *gin.Context) {	
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error: ID '%s' is not a valid integer", id)})
		return
	}

	task, err := helpers.GetTaskByID_help(c, intId)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var newTask models.TAddTask

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.DB.Exec("insert into tasks (text, timer) values ($1, $2)", newTask.Text, newTask.Timer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"ok": "user is created"})
}