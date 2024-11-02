package entity

import "time"

type Urls struct {
	ID                  uint      `json:"id"`
	OriginalUrl         string    `json:"originalUrl"`
	ShortUrl            string    `json:"shortUrl"`
	ShortUrlWhithDomain string    `json:"shortDomainUrl"`
	CreatedAt           time.Time `json:"created_at"`
	UserLogin           string    `json:"user_login"`
}
