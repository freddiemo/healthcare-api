package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/freddiemo/healthcare-api/internal/diagnostics/model"
	"github.com/freddiemo/healthcare-api/internal/diagnostics/service"
)

type DiagnosticsController interface {
	Save(ctx *gin.Context) (model.Diagnostic, error)
	Find() ([]model.Diagnostic, error)
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
	if err := ctx.ShouldBind(&diagnostic); err != nil {
		return model.Diagnostic{}, err
	}

	diagnostic, err := controller.diagnosticService.Save(diagnostic)
	if err != nil {
		return model.Diagnostic{}, err
	}

	return diagnostic, nil
}

func (controller *diagnosticsController) Find() ([]model.Diagnostic, error) {
	var diagnostics []model.Diagnostic
	diagnostics, err := controller.diagnosticService.Find()
	if err != nil {
		return diagnostics, err
	}

	return diagnostics, nil
}
