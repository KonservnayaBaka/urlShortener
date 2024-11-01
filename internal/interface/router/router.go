package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"urlShortener/internal/interface/controller"
)

func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	linkGroup := r.Group("/link")
	{
		linkGroup.POST("/shorten", controller.MakeShortLink(db))
		linkGroup.GET("/:short_url", controller.FollowShortLink(db))
	}
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", controller.RegUser(db))
		authGroup.GET("/signin", controller.AuthUser(db))
	}
}
