package db

import (
	"log"
	"task/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=task_db user=postgres password=qwerty dbname=tasks port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Автоматическая миграция моделей
	database.AutoMigrate(&models.Task{})

	DB = database
}
