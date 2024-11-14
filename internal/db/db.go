package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

// ConnectDB initializes a PostgreSQL connection using GORM
func ConnectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "gomovie.", // Prefixes every table with the specified schema
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}
	log.Println("Connected to PostgreSQL database!")
	return db, nil
}
