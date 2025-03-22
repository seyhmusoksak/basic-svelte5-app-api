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

	tasksRepository := repository.NewTasksRepository(database.DB)
	tasksService := service.NewTasksService(tasksRepository)
	tasksController := controller.NewTasksController(tasksService)

	collectionsRepository := repository.NewCollectionsRepository(database.DB)
	collectionsService := service.NewCollectionsService(collectionsRepository)
	collectionsController := controller.NewCollectionsController(collectionsService)

	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService, collectionsService, tasksService)

	r := gin.Default()
	r.GET("/users/:id", userController.GetUserByID)
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id/tasks", userController.GetUserTasksByID)
	r.GET("/users/:id/collections", userController.GetUserCollectionsByID)
	r.GET("/tasks", tasksController.GetAllTasks)
	r.GET("/collections", collectionsController.GetAllCollections)
	r.Run(":8081")
}
