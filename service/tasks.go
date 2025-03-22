package service

import (
	"github.com/seyhmusoksak/to-do-api/models"
	"github.com/seyhmusoksak/to-do-api/repository"
)

type TasksService struct
{
	TasksRepository *repository.TasksRepository
}

func NewTasksService(repository *repository.TasksRepository) *TasksService {
	return &TasksService{
		TasksRepository: repository,
	}
}

func (s *TasksService) GetAllTasks() (*[]models.Tasks, error) {
	tasks, err := s.TasksRepository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, err
}

func (c *TasksService) GetUserTasksByID(userID int) (*[]models.Tasks, error) {
	tasks, err := c.TasksRepository.GetUserTasksByID(userID)
	if
		err != nil {
		return nil, err
	}
	return tasks, nil
}

