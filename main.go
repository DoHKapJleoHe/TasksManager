package main

import (
	"GoProject/db"
	"GoProject/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.NewDatabase()
	if err != nil {
		return
	}

	defer db.DB.Db.Close()

	router := gin.Default()
	router.GET("/tasks", handlers.GetTasks)
	router.POST("/tasks", handlers.CreateTask)

	router.Run("localhost:8080")
}
