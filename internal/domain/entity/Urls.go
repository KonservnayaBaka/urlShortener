package entity

import "time"

type Urls struct {
	ID          uint      `json:"id"`
	OriginalUrl string    `json:"originalUrl"`
	ShortUrl    string    `json:"shortUrl"`
	CreatedAt   time.Time `json:"created_at"`
}
