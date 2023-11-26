package api

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Init(params map[string]string) {
	setupLogOutput(params["APP_NAME"])
	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!",
		})
	})

	server.Run(fmt.Sprintf(":%s", params["APP_PORT"]))
}

func setupLogOutput(app_name string) {
	file, err := os.Create(fmt.Sprintf("%s.log", app_name))
	if err != nil {
		log.Fatal("Cannot create log file")
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}
