package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/seyhmusoksak/to-do-api/models"
	"github.com/seyhmusoksak/to-do-api/repository"
	"net/http"
	"strconv"
)

// I will add user_id and collection_id check to this function
func CreateTask(c *gin.Context) {
	var task models.Tasks
	// Validate the request body
	err := c.ShouldBindJSON(&task)
	// Error Checking
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.CreateTask(&task, c)
}

func GetAllTasks(c *gin.Context) {
	var tasks []models.Tasks

	// Sending to repository
	repository.GetAllTasks(&tasks, c)
}

func GetTaskById(c *gin.Context) {
	var task models.Tasks
	// Getting task id
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.GetTaskById(&task, taskID, c)
}

func PatchTaskById(c *gin.Context) {
	var task map[string]interface{}

	// Check the data from client
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Check if the body is empty
	if len(task) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty data."})
		return
	}

	// Creating a map, because we can update only few things
	changeableField := map[string]bool{
		"title":         true,
		"description":   true,
		"created_at":    false,
		"is_completed":  true,
		"collection_id": false,
		"user_id":       false,
	}

	// Now Creating a filtered map
	filteredUpdates := make(map[string]interface{}, len(task))
	for key, value := range task {
		if changeableField[key] {
			filteredUpdates[key] = value
		}
	}

	// Check empty values
	for key, value := range filteredUpdates {
		if value == "" {
			if key != "description" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request."})
				return
			}
		}
	}

	// Checking the type of id
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	// Sending to repository
	repository.PatchTaskById(&filteredUpdates, taskID, c)
}

func DeleteTaskByID(c *gin.Context) {
	// Getting task id
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Sending to repository
	repository.DeleteTaskByID(taskID, c)
}
