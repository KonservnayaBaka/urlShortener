package tests

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"urlShortener/internal/domain/entity"
)

func setupTestDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=1234 dbname=testurlshortner port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.User{}, &entity.Urls{})

	return db, nil
}
