package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/seyhmusoksak/to-do-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDB() (error) {
	env := godotenv.Load()
	if env != nil {
		return env
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.Collection{}, &models.Todo{})
	DB = db
	return nil
}
