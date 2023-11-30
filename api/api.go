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

// @title        Healthcare API
// @version      1.0
// @description  A Healthcare service API in Go using Gin
// @BasePath     /v1

// ListDiagnostics godoc
// @Summary      List diagnostics
// @Description  List dignostic data from bd.
// @Tags         diagnostics
// @Produce      application/json
// @Param        firstName  query    string  false  "Filter diagnostics by firstName"
// @Param        lastName   query    string  false  "Filter diagnostics by lastName"
// @Param        date       query    string  false  "Filter diagnostics by date"
// @Success      200        {array}  model.DiagnosticWithPatient{}
// @Router       /diagnostics [get]
func (api *HealthcareAPI) FindDiagnostics(ctx *gin.Context) {
	diagnostics, err := api.diagnosticsController.Find(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, diagnostics)
	}
}

// CreateDiagnostics 	godoc
// @Summary      Create diagnostic
// @Description  Save diagnostic data in bd.
// @Param        diagnostic  body  model.DiagnosticRequest{}  true  "Create diagnostic"
// @Produce      application/json
// @Tags         diagnostics
// @Success      201  {object}  model.Diagnostic{}
// @Failure      400  {object}  Response{}
// @Router       /diagnostics [post]
func (api *HealthcareAPI) SaveDiagnostic(ctx *gin.Context) {
	diagnostic, err := api.diagnosticsController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, diagnostic)
	}
}
