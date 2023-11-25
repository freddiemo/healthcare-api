package main

import (
	"github.com/gin-gonic/gin"

	"github.com/freddiemo/healthcare-api/config"
)

func main() {
	params := config.Init()
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!",
		})
	})

	server.Run(":" + params["APP_PORT"])
}
