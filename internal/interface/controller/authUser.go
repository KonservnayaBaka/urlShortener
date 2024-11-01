package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"urlShortener/internal/application/service"
	"urlShortener/internal/domain/entity"
)

func AuthUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userModel = &entity.User{}

		if err := c.ShouldBindJSON(userModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if userModel.Login == "" || userModel.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "login or password is empty"})
			return
		}

		var storedUser entity.User
		if err := db.Where("login = ?", userModel.Login).First(&storedUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid login or password"})
			return
		}

		if !service.CheckPasswordHash(userModel.Password, storedUser.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid login or password"})
			return
		}

		storedUser.Password = ""

		c.JSON(http.StatusOK, gin.H{"status": true, "data": storedUser})
	}
}
