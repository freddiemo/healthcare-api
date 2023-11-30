package api

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/freddiemo/healthcare-api/api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	diagnosticsCtrl "github.com/freddiemo/healthcare-api/internal/diagnostics/controller"
	diagnosticsRepo "github.com/freddiemo/healthcare-api/internal/diagnostics/repository"
	diagnosticsServ "github.com/freddiemo/healthcare-api/internal/diagnostics/service"
)

func Init(params map[string]string, db *gorm.DB) {
	setupLogOutput(params["APP_NAME"])
	server := GetServer(db)

	server.GET("/swagger/*any", func(ctx *gin.Context) {
		docs.SwaggerInfo.Host = ctx.Request.Host
		ginSwagger.WrapHandler(swaggerFiles.Handler)(ctx)
	})

	server.GET("/health", func(c *gin.Context) {
		c.String(200, "Healthy")
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

func GetServer(db *gorm.DB) *gin.Engine {
	var diagnosticsRepository diagnosticsRepo.DiagnosticsRepository = diagnosticsRepo.NewDiagnosticsRepository(db)
	var diagnosticsService diagnosticsServ.DiagnosticsService = diagnosticsServ.NewDiagnosticsService(diagnosticsRepository)
	var diagnosticsController diagnosticsCtrl.DiagnosticsController = diagnosticsCtrl.NewDiagnosticsController(diagnosticsService)

	server := gin.Default()
	healthcareAPI := NewHealthcareAPI(
		diagnosticsController,
	)
	getRegisterRoutes(healthcareAPI, server)

	return server
}
