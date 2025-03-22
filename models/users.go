package models

import (
	"time"
)

// Migration Model
type User struct {
	ID          int          `json:"id" gorm:"primary_key"`
	Name        string       `json:"name" binding:"required"`
	CreatedAt   time.Time    `json:"created_at"`
	Collections []Collection `json:"collections" gorm:"foreignKey:UserID"`
	Tasks       []Tasks      `json:"tasks" gorm:"foreignKey:UserID"`
}

// Service Models
type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type UserUpdate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
