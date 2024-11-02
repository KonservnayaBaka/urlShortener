package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"urlShortener/internal/domain/entity"
)

func GetUserLinks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.MustGet("username").(string)

		var urls []entity.Urls
		if err := db.Where("user_login = ?", username).Find(&urls).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, urls)
	}
}
