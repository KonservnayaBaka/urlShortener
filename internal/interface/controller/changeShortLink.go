package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"urlShortener/internal/application/service"
	"urlShortener/internal/domain/entity"
	"urlShortener/internal/utils"
)

type ChangeShortLinkRequest struct {
	ShortUrl    string `json:"shortUrl"`
	NewShortUrl string `json:"newShortUrl"`
}

func ChangeShortLink(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var urlModel = &entity.Urls{}
		var request ChangeShortLinkRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if request.ShortUrl == "" || request.NewShortUrl == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "shortUrl or newShortUrl is required"})
			return
		}

		if service.UniqCheck(request.NewShortUrl, db) == true {
			c.JSON(http.StatusBadRequest, gin.H{"error": "short url is exist"})
			return
		}

		if err := db.Where("short_url = ?", request.ShortUrl).First(urlModel).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "short url not found"})
			return
		}

		urlModel.ShortUrl = request.NewShortUrl
		urlModel.ShortUrlWhithDomain = utils.Domain() + request.NewShortUrl

		username := c.MustGet("username").(string)
		urlModel.UserLogin = username

		if err := db.Save(&urlModel).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, &urlModel)
	}
}
