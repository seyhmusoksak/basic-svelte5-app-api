package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/seyhmusoksak/to-do-api/controller"
	"github.com/seyhmusoksak/to-do-api/repository"
	"github.com/seyhmusoksak/to-do-api/service"
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

	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	r := gin.Default()
	r.GET("/users/:id", userController.GetUser)
	r.POST("/users", userController.CreateUser)
	r.POST("/users/:id/collections", userController.CreateCollection)
	r.GET("/users/:id/collections", userController.GetUserCollectionByID)
	r.Run(":8081")
}
