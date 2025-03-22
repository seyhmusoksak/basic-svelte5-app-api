package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/seyhmusoksak/to-do-api/models"
	"github.com/seyhmusoksak/to-do-api/service"
)

type CollectionsController struct
{
	TasksService *service.CollectionsService
}

func NewCollectionsController(service *service.CollectionsService) *CollectionsController {
	return &CollectionsController{
		TasksService: service,
	}
}

func (c *CollectionsController) GetAllCollections(ctx *gin.Context) {
	collections, err := c.TasksService.GetAllCollections()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, collections)
}

func (c *CollectionsController) GetUserCollectionsByID(userID int) (*[]models.Collection, error) {
	collections, err := c.TasksService.GetUserCollectionsByID(userID)
	if err != nil {
		return nil, err
	}
	return collections, nil
}
