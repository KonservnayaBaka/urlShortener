package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=urlshortener port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return db
}
