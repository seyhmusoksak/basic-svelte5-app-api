package repository

import (
	"gorm.io/gorm"
	"github.com/seyhmusoksak/to-do-api/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) CreateUser(user interface{}) error {
	err := u.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) CreateCollection(collection interface{}) error {
	err := u.DB.Create(collection).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetUserCollectionByID(userID int) (interface{}, error) {
	var schema []models.Collection
	err := u.DB.Preload("Tasks").Where("user_id = ?", userID).Find(&schema).Error
	if err != nil {
		return nil, err
	}
	return schema, nil
}

func (u *UserRepository) GetUser(userId string) (interface{}, error) {

	var schema models.UserResponse

	err := u.DB.Table("users").Select("id, name").Where("id = ?", userId).First(&schema).Error

	if err != nil {
		return nil, err
	}
	return schema, nil
}

