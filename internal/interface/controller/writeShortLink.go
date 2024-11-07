package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"urlShortener/internal/application/service"
	"urlShortener/internal/domain/entity"
	"urlShortener/internal/utils"
)

func WriteShortLink(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var urlModel = &entity.Urls{}
		if err := c.ShouldBind(&urlModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if urlModel.OriginalUrl == "" || urlModel.ShortUrl == "" {
			c.JSON(http.StatusBadRequest, gin.H{"originalUrl": "original or short url is empty"})
			return
		}

		if service.ValidateUrl(urlModel.OriginalUrl) == false {
			c.JSON(http.StatusBadRequest, gin.H{"originalUrl": "URL is not validated"})
			return
		}

		if service.UniqCheck(urlModel.ShortUrl, db) == true {
			c.JSON(http.StatusBadRequest, gin.H{"originalUrl": "short url is exist"})
			return
		}

		urlModel.ShortUrlWhithDomain = utils.Domain() + urlModel.ShortUrl

		username := c.MustGet("username").(string)
		urlModel.UserLogin = username

		if err := db.Create(&urlModel).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, &urlModel)
	}
}
