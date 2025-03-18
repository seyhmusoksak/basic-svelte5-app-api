package models

import (
	"time"
)

type Collection struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	Tasks       []Tasks    `json:"tasks" gorm:"foreignKey:CollectionID"` // Collection'a ait foreign key
}
