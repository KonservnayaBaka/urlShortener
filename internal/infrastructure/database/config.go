package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=1234 dbname=dbname port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}
