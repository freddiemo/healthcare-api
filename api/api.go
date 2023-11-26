package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	diagnostics "github.com/freddiemo/healthcare-api/internal/diagnostics/controller"
)

type HealthcareAPI struct {
	diagnosticsController diagnostics.DiagnosticsController
}

func NewHealthcareAPI(
	diagnosticsController diagnostics.DiagnosticsController,
) *HealthcareAPI {
	return &HealthcareAPI{
		diagnosticsController: diagnosticsController,
	}
}

func (api *HealthcareAPI) SaveDiagnostic(ctx *gin.Context) {
	diagnostic, err := api.diagnosticsController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, diagnostic)
	}
}

func (api *HealthcareAPI) FindDiagnostics(ctx *gin.Context) {
	diagnostics, err := api.diagnosticsController.Find()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, diagnostics)
	}
}
