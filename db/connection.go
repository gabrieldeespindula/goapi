package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gabrieldeespindula/goapi/models"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	err = database.AutoMigrate(&models.Record{})
	if err != nil {
			log.Fatalf("Failed to AutoMigrate: %v", err)
	}

	DB = database
}
