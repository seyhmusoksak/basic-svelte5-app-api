package models

import (
	"time"
)

type Tasks struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" binding:"required" gorm:"not null"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	CollectionID int       `json:"collection_id" binding:"required"`
}
