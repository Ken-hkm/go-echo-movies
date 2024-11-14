package db

import (
	"go-echo-movies/internal/model"
	"gorm.io/gorm"
	"log"
)

// MigrateDB migrates the database by creating tables for all models
func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&model.Movie{})
	if err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}
	log.Println("Database migration completed!")
}
