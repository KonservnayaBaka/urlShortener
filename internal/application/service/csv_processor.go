package service

import (
	"encoding/csv"
	"gorm.io/gorm"
	"io"
	"log"
	"sync"
	"urlShortener/internal/domain/entity"
)

func ProcessCSV(file io.Reader, db *gorm.DB, username string) error {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	processRecord := func(record []string) {
		defer wg.Done()
		originalUrl := record[0]
		shortUrl := GenerateShortURL(originalUrl, db)
		urlModel := &entity.Urls{
			OriginalUrl: originalUrl,
			ShortUrl:    shortUrl,
			UserLogin:   username,
		}

		mu.Lock()
		defer mu.Unlock()

		tx := db.Begin()
		if err := tx.Error; err != nil {
			log.Printf("Failed to begin transaction: %v", err)
			return
		}

		if err := db.Create(&urlModel).Error; err != nil {
			tx.Rollback()
			log.Printf("Failed to create URL: %v", err)
			return
		}

		tx.Commit()
	}

	for _, record := range records {
		wg.Add(1)
		go processRecord(record)
	}

	wg.Wait()

	return nil
}
