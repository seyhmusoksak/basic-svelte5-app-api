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
	r.GET("/collections", controller.GetCollections)
	r.GET("/collections/:id", controller.GetCollectionsByID)
	r.POST("/collections", controller.CreateCollection)
	r.PUT("/collections/:id", controller.UpdateCollection)
	r.DELETE("/collections/:id", controller.DeleteCollection)
	r.Run(":8080")
}
