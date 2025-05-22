package storage

import (
	// "log"
	// "user/models"

	// "gorm.io/driver/sqlite"
	"log"
	"user/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// this shit is laced  :(
var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate schema
	DB.AutoMigrate(&models.User{})
}

func main() {
	InitDB()
	log.Println("Database initialized and User.db created successfully.")
}
