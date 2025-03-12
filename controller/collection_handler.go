package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/seyhmusoksak/to-do-api/database"
	"github.com/seyhmusoksak/to-do-api/models"
)

func GetCollections(c *gin.Context) {
	collections, err := models.GetAllCollections(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch collections"})
		return
	}
	c.JSON(http.StatusOK, collections)
}

func GetCollectionsByID(c *gin.Context) {
	id := c.Param("id")
	collection, err := models.GetCollectionByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Collection not found"})
		return
	}
	c.JSON(http.StatusOK, collection)
}

func CreateCollection(c *gin.Context) {
	var collection models.Collection
	if err := c.ShouldBindJSON(&collection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	newCollection, err := models.CreateCollection(database.DB, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create collection"})
		return
	}
	c.JSON(http.StatusCreated, newCollection)
}

func UpdateCollection(c *gin.Context) {
	id := c.Param("id")
	var collection models.Collection
	if err := c.ShouldBindJSON(&collection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	updatedCollection, err := models.UpdateCollection(database.DB, id, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update collection"})
		return
	}
	c.JSON(http.StatusOK, updatedCollection)
}

func DeleteCollection(c *gin.Context) {
	id := c.Param("id")
	if err := models.DeleteCollection(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete collection"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
