package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"urlShortener/internal/application/service"
	"urlShortener/internal/domain/entity"
	"urlShortener/internal/utils"
)

func MakeShortLink(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var urlModel = &entity.Urls{}
		if err := c.ShouldBind(&urlModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if urlModel.OriginalUrl == "" {
			c.JSON(http.StatusBadRequest, gin.H{"originalUrl": "String is empty"})
			return
		}

		if service.ValidateUrl(urlModel.OriginalUrl) == false {
			c.JSON(http.StatusBadRequest, gin.H{"originalUrl": "URL is not validated"})
			return
		}

		shortUrl := service.GenerateShortURL(urlModel.OriginalUrl, db)
		urlModel.ShortUrl = shortUrl
		urlModel.ShortUrlWhithDomain = utils.Domain() + shortUrl

		username := c.MustGet("username").(string)
		urlModel.UserLogin = username

		if err := db.Create(&urlModel).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, urlModel)
	}
}
