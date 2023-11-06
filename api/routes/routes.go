package routes

import (
	"apple-reminder_backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	router.GET("/tasks", handlers.GetAllTasks)
	router.GET("/tasks/:id", handlers.GetTaskByID)
	router.POST("/tasks", handlers.CreateTask)
	// router.PATCH("/tasks/:id", handlers.UpdateTaskIsDone)
	// router.DELETE("/tasks/:id", handlers.DeleteTask)
}