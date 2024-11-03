package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"urlShortener/internal/application/service"
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

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := service.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
