package repository

import (
	"github.com/seyhmusoksak/to-do-api/models"
	"gorm.io/gorm"
)

type TasksRepository struct
{
	db *gorm.DB
}

func NewTasksRepository(connection *gorm.DB) *TasksRepository {
	return &TasksRepository{
		db: connection,
	}
}

func (r *TasksRepository) GetAllTasks() (*[]models.Tasks, error) {
	var tasks []models.Tasks
	err := r.db.Table("tasks").Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return &tasks, nil
}

func (r *TasksRepository) GetUserTasksByID(userID int) (*[]models.Tasks, error) {
	var tasks []models.Tasks
	err := r.db.Table("tasks").Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return &tasks, nil
}
