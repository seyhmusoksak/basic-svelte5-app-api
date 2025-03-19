package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/seyhmusoksak/to-do-api/controller"
	"github.com/seyhmusoksak/to-do-api/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	err = database.ConnectDB()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}
	r := gin.Default()
	// r.GET("/collections/:user_id", controller.GetCollectionsByUserId)
	// r.GET("/collections/:id", controller.GetCollectionsByID)
	// r.GET("/collections", controller.GetAllCollections)
	r.GET("/tasks", controller.GetAllTasks)
	r.GET("/tasks/:id", controller.GetTaskById)
	r.PATCH("/tasks/:id", controller.PatchTaskById)
	r.POST("/tasks", controller.CreateTask)
	r.Run(":8081")
}
