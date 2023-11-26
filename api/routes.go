package api

import (
	"github.com/gin-gonic/gin"
)

func addHealthcareRoutes(rg *gin.RouterGroup, healthcareAPI *HealthcareAPI) {
	clients := rg.Group("/diagnostics")
	{
		clients.POST("/", healthcareAPI.SaveDiagnostic)
	}
}

func getRegisterRoutes(healthcareAPI *HealthcareAPI, server *gin.Engine) {
	v1 := server.Group("/v1")
	addHealthcareRoutes(v1, healthcareAPI)
}
