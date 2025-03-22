package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/seyhmusoksak/to-do-api/service"
)

type TasksController struct
{
	TasksService *service.TasksService
}

func NewTasksController(service *service.TasksService) *TasksController {
	return &TasksController{
		TasksService: service,
	}
}


func (c *TasksController) GetAllTasks(ctx *gin.Context) {
	tasks, err := c.TasksService.GetAllTasks()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, tasks)
}

