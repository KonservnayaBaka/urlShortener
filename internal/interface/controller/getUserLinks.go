package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"urlShortener/internal/domain/entity"
)

func GetUserLinks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username query parameter is required"})
			return
		}

		var urls []entity.Urls
		if err := db.Where("user_login = ?", username).Find(&urls).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, urls)
	}
}
