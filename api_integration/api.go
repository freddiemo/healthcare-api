package api_integration

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthcareAPI struct {
	params map[string]string
}

func NewHealthcareAPI(
	params map[string]string,
) *HealthcareAPI {
	return &HealthcareAPI{
		params: params,
	}
}

// @title        Healthcare API Integration
// @version      1.0
// @description  A Healthcare integration service API in Go using Gin
// @BasePath     /v1

// ListDiagnostics godoc
// @Summary      List diagnostics
// @Description  List dignostic data from bd.
// @Tags         diagnostics
// @Produce      application/json
// @Success      200  {array}  DiagnosticResponseWithPatient{}
// @Router       /diagnostics [get]
func (api *HealthcareAPI) FindDiagnostics(ctx *gin.Context) {
	diagnosticResp, err := getDiagnostics(api.params, "/v1/diagnostics")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, diagnosticResp)
	}
}

// CreateDiagnostics 	godoc
// @Summary      Create diagnostic
// @Description  Save diagnostic data in bd.
// @Param        diagnostic  body  DiagnosticRequest{}  true  "Create diagnostic"
// @Produce      application/json
// @Tags         diagnostics
// @Success      201  {object}  model.Diagnostic{}
// @Failure      400  {object}  Response{}
// @Router       /diagnostics [post]
func (api *HealthcareAPI) SaveDiagnostic(ctx *gin.Context) {
	diagnostic, err := parseDiagnosticRQ(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		diagnosticResp, err := postDiagnostic(api.params, "/v1/diagnostics", diagnostic)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &Response{
				Message: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusCreated, diagnosticResp)
		}
	}
}
