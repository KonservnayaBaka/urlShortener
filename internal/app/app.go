package app

import (
	"github.com/gin-gonic/gin"
	"urlShortener/internal/domain/entity"
	"urlShortener/internal/infrastructure/database"
	"urlShortener/internal/interface/router"
)

func Run() {
	r := gin.Default()
	db := database.NewDB()

	db.AutoMigrate(&entity.Urls{}, &entity.User{})

	r.StaticFile("/swagger-ui", "./static/swagger-ui.html")
	r.GET("/swagger.yaml", func(c *gin.Context) {
		c.File("./swagger.yaml")
	})
	router.InitializeRoutes(r, db)
	r.Run("localhost:8085")
}
