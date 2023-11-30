package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/freddiemo/healthcare-api/internal/diagnostics/model"
	"github.com/freddiemo/healthcare-api/internal/diagnostics/service"
)

type DiagnosticsController interface {
	Save(ctx *gin.Context) (model.Diagnostic, error)
	Find(ctx *gin.Context) ([]model.DiagnosticWithPatient, error)
}

type diagnosticsController struct {
	diagnosticService service.DiagnosticsService
}

func NewDiagnosticsController(diagnosticService service.DiagnosticsService) DiagnosticsController {
	return &diagnosticsController{
		diagnosticService: diagnosticService,
	}
}

func (controller *diagnosticsController) Save(ctx *gin.Context) (model.Diagnostic, error) {
	var diagnostic model.Diagnostic
	if err := ctx.ShouldBindJSON(&diagnostic); err != nil {
		return model.Diagnostic{}, err
	}

	if err := controller.diagnosticService.Save(&diagnostic); err != nil {
		return model.Diagnostic{}, err
	}

	return diagnostic, nil
}

func (controller *diagnosticsController) Find(ctx *gin.Context) ([]model.DiagnosticWithPatient, error) {
	queryParams := model.DiagnosticQuery{}
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		return nil, err
	}

	var diagnostics []model.DiagnosticWithPatient
	diagnostics, err := controller.diagnosticService.Find(&queryParams)
	if err != nil {
		return diagnostics, err
	}

	return diagnostics, nil
}
