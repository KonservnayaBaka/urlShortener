package app

import "github.com/gin-gonic/gin"

func Run() {
	r := gin.Default()

	r.Run("localhost:8085")
}
