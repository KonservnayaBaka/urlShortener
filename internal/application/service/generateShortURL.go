package service

import (
	"crypto/sha256"
	"encoding/hex"
	"gorm.io/gorm"
	"urlShortener/internal/domain/entity"
)

func GenerateShortURL(originalUrl string, db *gorm.DB) string {
	hash := sha256.Sum256([]byte(originalUrl))
	hashString := hex.EncodeToString(hash[:])

	shortUrl := hashString[:8]

	if UniqCheck(shortUrl, db) {
		shortUrl = GenerateShortURL(shortUrl, db)
	}

	return shortUrl
}

func UniqCheck(shortUrl string, db *gorm.DB) bool {
	check := db.Where("short_url = ?", shortUrl).First(&entity.Urls{})
	if check.Error != nil {
		return false
	}

	return true
}
