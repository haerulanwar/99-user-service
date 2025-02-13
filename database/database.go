package database

import (
	"fmt"
	"log"

	"99-user-service/config"
	"99-user-service/internal/app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dbName := config.GetEnv("DB_NAME", "database.db")

	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connected successfully")

	DB.AutoMigrate(&models.User{})
}
