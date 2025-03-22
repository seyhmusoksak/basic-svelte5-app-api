package service

import (
	"time"

	"github.com/seyhmusoksak/to-do-api/models"
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

func (s *UserService) CreateUser(user *models.User) error {
	// Set Create Time
	user.CreatedAt = time.Now()
	err := s.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetAllUsers() ([]models.UserResponse, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, err
}

func (s *UserService) GetUserByID(userId int) (*models.UserResponse, error) {
	user, err := s.userRepository.GetUserByID(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}


func (s *UserService) UpdateUser(user models.UserUpdate, userId int) (*models.UserResponse, error) {
	UpdatedUser, err := s.userRepository.UpdateUser(user, userId)
	if err != nil {
		return nil, err
	}
	return UpdatedUser, nil
}

func (s *UserService) DeleteUser(userId int) error {
	err := s.userRepository.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}
