package service

import (
	"github.com/seyhmusoksak/to-do-api/models"
	"github.com/seyhmusoksak/to-do-api/repository"
)

type CollectionsService struct
{
	TasksRepository *repository.CollectionsRepository
}

func NewCollectionsService(repository *repository.CollectionsRepository) *CollectionsService {
	return &CollectionsService{
		TasksRepository: repository,
	}
}


func (s *CollectionsService) GetAllCollections() (*[]models.Collection, error) {
	collections, err := s.TasksRepository.GetAllCollections()
	if err != nil {
		return nil, err
	}
	return collections, err
}

func (s *CollectionsService) GetUserCollectionsByID(userID int) (*[]models.Collection, error) {
	collections, err := s.TasksRepository.GetUserCollectionsByID(userID)
	if err != nil {
		return nil, err
	}
	return collections, err
}
