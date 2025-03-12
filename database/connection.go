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
		fmt.Println("Error loading .env file")
		return env
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return err
	}
	err = DB.AutoMigrate(&models.Collection{}, &models.Todo{})
	if err != nil {
		fmt.Println("Failed to migrate database:", err)
	}
	return nil
}
