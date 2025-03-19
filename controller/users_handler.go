package controller


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/seyhmusoksak/to-do-api/database"
	"github.com/seyhmusoksak/to-do-api/models"
)


// Test Func.
func GetAllUsers(c *gin.Context) {
	var users []models.User
	err := database.DB.Preload("Collections").Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

