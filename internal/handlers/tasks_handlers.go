package handlers

import (
	"apple-reminder_backend/internal/helpers"
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