package controller

import (
	"net/http"
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/seyhmusoksak/to-do-api/database"
	"github.com/seyhmusoksak/to-do-api/models"
)

type taskSchema struct {
	Title        string    `json:"title" binding:"required"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	CollectionID int       `json:"collection_id" binding:"required"`
	UserID       int       `json:"user_id" binding:"required"`
}


// I will add user_id and collection_id check to this function
func CreateTask(c *gin.Context) {
	var task taskSchema
	// Validate the request body
	err := c.ShouldBindJSON(&task)
	// Error Checking
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Insert into database
	result := database.DB.Select("title", "description", "created_at",
		"collection_id", "user_id").Create(&task)
	// Check if there is an error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
	}
	// Now, We can respond with 201 Created status code
	c.JSON(http.StatusCreated, task)
}

func GetAllTasks(c *gin.Context) {
	var tasks []models.Tasks
	// Join with collection table wiyh Preload method
	err := database.DB.Preload("Collections").Find(&tasks).Error
	// Checking Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return it..
	c.JSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	var task models.Tasks
	// Getting task id
	taskID := c.Param("id")
	if _ , err := strconv.Atoi(taskID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Using GORM first method to get only one record.
	result := database.DB.First(&task, taskID).Error
	// Checking Error
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error()})
		return
	}
	// Okey.. Returning it
	c.JSON(http.StatusOK, task)
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
	changeableField := map[string]bool {
		"title": true,
		"description": true,
		"created_at": false,
		"is_completed": true,
		"collection_id": false,
		"user_id": false,
	}

	// Now Creating a filtered map
	filteredUpdates := make(map[string]interface{}, len(task))
	for key, value := range task {
		if changeableField[key] {
			filteredUpdates[key] = value
		}
	}

	// Checking the type of id
	taskID := c.Param("id")
    _ , err := strconv.Atoi(taskID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

	// Making changes on database
	result := database.DB.Model(&models.Tasks{}).Where("id = ?", taskID).
		Updates(filteredUpdates)

	// Some controlling stuff
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or no changes made"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated."})
}

func DeleteTaskByID(c *gin.Context) {

	// Getting task id
	taskID := c.Param("id")
	if _ , err := strconv.Atoi(taskID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Delete the task
	result := database.DB.Delete(&models.Tasks{}, taskID)
	// Some Controlling
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// StatusOK
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})

}


