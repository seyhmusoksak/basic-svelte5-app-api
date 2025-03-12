package models

import (
	"time"
)

type Todo struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	CollectionID int       `json:"collection_id"`
	Collection   Collection `gorm:"foreignKey:CollectionID"`
}
