package api_integration

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/freddiemo/healthcare-api/api_integration/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(params map[string]string) {
	setupLogOutput(params["APP_NAME"])
	server := GetServer(params)

	server.GET("/swagger/*any", func(ctx *gin.Context) {
		docs.SwaggerInfo.Host = ctx.Request.Host
		ginSwagger.WrapHandler(swaggerFiles.Handler)(ctx)
	})

	server.Run(fmt.Sprintf(":%s", params["APP_INTEGRATION_PORT"]))
}

func setupLogOutput(app_name string) {
	file, err := os.Create(fmt.Sprintf("%s_integration.log", app_name))
	if err != nil {
		log.Fatal("Cannot create log file")
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func GetServer(params map[string]string) *gin.Engine {
	server := gin.Default()
	healthcareAPI := NewHealthcareAPI(
		params,
	)
	getRegisterRoutes(healthcareAPI, server)

	return server
}
