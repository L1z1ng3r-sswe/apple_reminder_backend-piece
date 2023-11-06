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

func UpdateTaskIsDone(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error: ID '%s' is not a valid integer", id)})
		return
	}

	var updatedTask models.TUpdateTask
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	result, err := db.DB.Exec("UPDATE tasks SET isDone = $1, result = $2 WHERE id = $3", true, updatedTask.Result, intID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Item with ID '%d' not found", intID)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": "task IsDone field is updated"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error: ID '%s' is not a valid integer", id)})
		return
	}

	
	
	if _, err := helpers.GetTaskByID_help(c, intID); err != nil {
		return
	}

	result, err := db.DB.Exec("DELETE FROM tasks WHERE id = $1", intID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete task with ID '%d'", intID)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": fmt.Sprintf("Task with ID '%d' is deleted", intID)})
}
