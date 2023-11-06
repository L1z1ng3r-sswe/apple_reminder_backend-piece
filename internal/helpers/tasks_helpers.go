package helpers

import (
	"apple-reminder_backend/db"
	"apple-reminder_backend/internal/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetALLTask_help() ([]models.TGetTask, error) {
	var tasks []models.TGetTask
	tasks_pstgr, err := db.DB.Query("select * from tasks")
	if err != nil {
		return nil, err
	}
	defer tasks_pstgr.Close()

	for tasks_pstgr.Next() {
		var task models.TGetTask

		err := tasks_pstgr.Scan(&task.ID, &task.Text, &task.Timer, &task.Result, &task.CreatedDate, &task.IsDone)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTaskByID_help(c *gin.Context, id int) (models.TGetTask, error) {
	var task models.TGetTask
	err := db.DB.QueryRow("SELECT * FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Text, &task.Timer, &task.Result, &task.CreatedDate, &task.IsDone)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return task, err
	}

	return task, nil
}