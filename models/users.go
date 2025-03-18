package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	Collections []Collection `json:"collections" gorm:"foreignKey:UserID"`
}


