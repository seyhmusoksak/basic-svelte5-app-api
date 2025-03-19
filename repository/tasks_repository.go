package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/seyhmusoksak/to-do-api/database"
	"github.com/seyhmusoksak/to-do-api/models"
	"net/http"
)

// I will add user_id and collection_id check to this function
func CreateTask(task *models.Tasks, c *gin.Context) {
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

func GetAllTasks(tasks *[]models.Tasks, c *gin.Context) {
	// Join with collection table wiyh Preload method
	err := database.DB.Find(tasks).Error
	// Checking Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return it..
	c.JSON(http.StatusOK, tasks)
}

func GetTaskById(task *models.Tasks, taskID int, c *gin.Context) {
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

func PatchTaskById(filteredUpdates *map[string]interface{}, taskID int, c *gin.Context) {

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

func DeleteTaskByID(taskID int, c *gin.Context) {
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
