package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"urlShortener/internal/interface/controller"
)

func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	linkGroup := r.Group("/link")
	{
		linkGroup.POST("/shorten", JWTAuthMiddleware(), controller.MakeShortLink(db))
		linkGroup.GET("/user", JWTAuthMiddleware(), controller.GetUserLinks(db))
		linkGroup.POST("/upload", JWTAuthMiddleware(), controller.UploadCSV(db)) // Новый маршрут
	}
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", controller.RegUser(db))
		authGroup.POST("/signin", controller.AuthUser(db))
	}
	r.GET("/:short_url", controller.FollowShortLink(db))
}
