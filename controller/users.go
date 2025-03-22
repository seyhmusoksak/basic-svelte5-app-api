package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/seyhmusoksak/to-do-api/models"
	"github.com/seyhmusoksak/to-do-api/service"
)

type UserController struct {
	userService        *service.UserService
	collectionsService *service.CollectionsService
	serviceTasks       *service.TasksService
}

func NewUserController(userService *service.UserService,
	collectionService *service.CollectionsService, tasksService *service.TasksService) *UserController {
	return &UserController{
		userService:        userService,
		collectionsService: collectionService,
		serviceTasks:       tasksService,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.userService.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, "User created successfully")
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	user, err := c.userService.GetUserByID(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var schema models.UserUpdate
	if err := ctx.ShouldBindJSON(&schema); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	updatedUser, err := c.userService.UpdateUser(schema, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = c.userService.DeleteUser(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "User deleted successfully")
}

func (c *UserController) GetUserCollectionsByID(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	collections, err := c.collectionsService.GetUserCollectionsByID(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, collections)
}


func (c *UserController) GetUserTasksByID(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	tasks, err := c.serviceTasks.GetUserTasksByID(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}
