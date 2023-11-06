package main

import (
	"apple-reminder_backend/api/routes"
	"apple-reminder_backend/db"
	"log"

	"github.com/gin-gonic/gin"
)


func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize the database:", err)
	}

	defer db.CloseDB()
	router := gin.New()

	routes.SetUpRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Println("server cant launch")
		log.Fatal(err.Error())
	}
}