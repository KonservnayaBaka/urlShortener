package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"urlShortener/internal/interface/controller"
)

func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/shorten", controller.MakeShortLink(db))
	r.GET("/:short_url", controller.FollowShortLink(db))
}
