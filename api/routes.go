package api

import (
	"github.com/gin-gonic/gin"
)

func addHealthcareRoutes(rg *gin.RouterGroup, healthcareAPI *HealthcareAPI) {
	diagnostics := rg.Group("/diagnostics")
	{
		diagnostics.POST("/", healthcareAPI.SaveDiagnostic)
		diagnostics.GET("/", healthcareAPI.FindDiagnostics)
	}
}

func getRegisterRoutes(healthcareAPI *HealthcareAPI, server *gin.Engine) {
	v1 := server.Group("/v1")
	addHealthcareRoutes(v1, healthcareAPI)
}
