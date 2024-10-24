package app

import (
	"github.com/gin-gonic/gin"
	"urlShortener/internal/domain/entity"
	"urlShortener/internal/infrastructure/database"
)

func Run() {
	r := gin.Default()
	db := database.NewDB()

	db.AutoMigrate(&entity.Urls{})
	r.Run("localhost:8085")
}
