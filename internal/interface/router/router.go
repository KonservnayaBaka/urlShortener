package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"urlShortener/internal/interface/controller"
)

func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	urlsGroup := r.Group("/urls")
	{
		urlsGroup.POST("/shorten", controller.MakeShortLink(db))
		urlsGroup.GET("/short_url", controller.FollowShortLink(db))
	}
}
