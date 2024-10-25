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

	if uniqCheck(shortUrl, db) == true {
		shortUrl = hashString[:8+1]
	}
	return shortUrl
}

func uniqCheck(shortUrl string, db *gorm.DB) bool {
	check := db.Where("short_url = ?", shortUrl).First(&entity.Urls{})
	if check.Error != nil {
		return false
	}

	return true
}
