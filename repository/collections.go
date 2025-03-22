package repository

import (
	"github.com/seyhmusoksak/to-do-api/models"
	"gorm.io/gorm"
)

type CollectionsRepository struct
{
	db *gorm.DB
}

func NewCollectionsRepository(connection *gorm.DB) *CollectionsRepository {
	return &CollectionsRepository{
		db: connection,
	}
}


func (r *CollectionsRepository) GetAllCollections() (*[]models.Collection, error) {
	var collections []models.Collection
	err := r.db.Table("collections").Find(&collections).Error
	if err != nil {
		return nil, err
	}
	return &collections, nil
}

func (r *CollectionsRepository) GetUserCollectionsByID(userID int) (*[]models.Collection, error) {
	var collections []models.Collection
	err := r.db.Table("collections").Where("user_id = ?", userID).Find(&collections).Error
	if err != nil {
		return nil, err
	}
	return &collections, nil
}
