package main

import (
	"fmt"
	"github.com/seyhmusoksak/to-do-api/database"
	"github.com/seyhmusoksak/to-do-api/controller"
	"github.com/gin-gonic/gin"
)


func main() {

	err := database.ConnectDB()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	r := gin.Default()
	r.GET("/collections", controller.GetCollections)
	r.GET("/collections/:id", controller.GetCollectionsByID)
	r.POST("/collections", controller.CreateCollection)
	r.PUT("/collections/:id", controller.UpdateCollection)
	r.DELETE("/collections/:id", controller.DeleteCollection)
	r.Run(":8080")
}
