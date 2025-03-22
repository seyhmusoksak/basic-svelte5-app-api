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

func (u *UserRepository) CreateUser(user *models.User) error {
	err := u.DB.Table("users").Omit("id").Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetAllUsers() ([]models.UserResponse, error) {

	var schema []models.UserResponse

	err := u.DB.Table("users").Find(&schema).Error

	if err != nil {
		return nil, err
	}

	return schema, nil
}


func (u *UserRepository) GetUserByID(userId int) (*models.UserResponse, error) {

	var schema models.UserResponse

	err := u.DB.Table("users").Select("id, name").Where("id = ?", userId).First(&schema).Error

	if err != nil {
		return nil, err
	}

	return &schema, nil
}

func (u *UserRepository) UpdateUser(user models.UserUpdate, userId int) (*models.UserResponse, error) {
	var schema models.UserResponse

	err := u.DB.Table("users").Where("id = ?", userId).Updates(user).First(&schema).Error

	if err != nil {
		return nil, err
	}

	return &schema, nil
}


func (u *UserRepository) DeleteUser(userId int) error {
	err := u.DB.Table("users").Where("id = ?", userId).Delete(&models.User{}).Error

	if err != nil {
		return err
	}

	return nil
}
