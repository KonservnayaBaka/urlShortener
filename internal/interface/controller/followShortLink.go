package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"urlShortener/internal/domain/entity"
)

func FollowShortLink(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortURL := c.Param("short_url")

		if shortURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "shortUrl is required"})
			return
		}

		var urlModel = &entity.Urls{}
		if err := db.Where("short_url = ?", shortURL).First(&urlModel).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
			return
		}

		c.Redirect(http.StatusMovedPermanently, urlModel.OriginalUrl)
	}
}
