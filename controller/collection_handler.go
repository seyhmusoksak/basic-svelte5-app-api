package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seyhmusoksak/to-do-api/database"
	"github.com/seyhmusoksak/to-do-api/models"
)

// func GetCollectionsByUserId(c *gin.Context) {
// 	// Get user_id from the request
// 	userId := c.Param("user_id")
// 	var user []models.User

// 	// Preloading the data
// 	if err := database.DB.Preload("Tasks").Find(&user, userId).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }

func GetAllCollections(c *gin.Context) {
	// Get all collections
	var collections []models.Collection
	err := database.DB.Preload("Tasks").Find(&collections).Error
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, collections)
}

func GetCollectionsByID (c *gin.Context) {
	// Get collection_id from the request
	Id := c.Param("id")
	var collections []models.Collection
	// Join data from tasks table
	err := database.DB.Preload("Tasks").Find(&collections, Id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, collections)
}
