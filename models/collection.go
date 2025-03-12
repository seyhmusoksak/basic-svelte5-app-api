package models

import (
	"time"
	"gorm.io/gorm"
)

type Collection struct {
	ID		  int    `json:"id" gorm:"primary_key"`
	Name	  string `json:"name" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	Description string    `json:"description"`
    IsCompleted bool      `json:"is_completed"`
}

func GetAllCollections(db *gorm.DB) ([]Collection, error) {
	var collections []Collection
	if err := db.Find(&collections).Error; err != nil {
		return nil, err
	}
	return collections, nil
}

func GetCollectionByID(db *gorm.DB, id string) (Collection, error) {
	var collection Collection
	if err := db.First(&collection, id).Error; err != nil {
		return Collection{}, err
	}
	return collection, nil
}

func CreateCollection(db *gorm.DB, collection Collection) (Collection, error) {
	if err := db.Create(&collection).Error; err != nil {
		return Collection{}, err
	}
	return collection, nil
}

func UpdateCollection(db *gorm.DB, id string, collection Collection) (Collection, error) {
	if err := db.Model(&Collection{}).Where("id = ?", id).Updates(collection).Error; err != nil {
		return Collection{}, err
	}
	return collection, nil
}

func DeleteCollection(db *gorm.DB, id string) error {
	if err := db.Delete(&Collection{}, id).Error; err != nil {
		return err
	}
	return nil
}
