package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"urlShortener/internal/application/service"
	"urlShortener/internal/domain/entity"
)

func RegUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userModel = &entity.User{}
		if err := c.ShouldBindJSON(&userModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if userModel.Login == "" || userModel.Password == "" || userModel.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "login or password or name is empty"})
			return
		}

		hashedPassword, err := service.HashPassword(userModel.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userModel.Password = hashedPassword

		if err := db.Create(&userModel).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true})
	}
}
