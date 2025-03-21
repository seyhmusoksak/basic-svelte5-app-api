package service

import (
	"github.com/seyhmusoksak/to-do-api/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}


func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) CreateUser(user interface{}) error {
	err := s.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) CreateCollection(collection interface{}) error {
	err := s.userRepository.CreateCollection(collection)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserCollectionByID(userID int) (interface{}, error) {
	collection, err := s.userRepository.GetUserCollectionByID(userID)
	if err != nil {
		return nil, err
	}
	return collection, nil
}


func (s *UserService) GetUser(userId string) (interface{}, error) {
	user, err := s.userRepository.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}



