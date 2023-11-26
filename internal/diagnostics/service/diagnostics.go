package service

import (
	"github.com/freddiemo/healthcare-api/internal/diagnostics/model"
	"github.com/freddiemo/healthcare-api/internal/diagnostics/repository"
)

type DiagnosticsService interface {
	Save(model.Diagnostic) (model.Diagnostic, error)
	Find() ([]model.Diagnostic, error)
}

type diagnosticsServ struct {
	diagnosticsRepository repository.DiagnosticsRepository
}

func NewDiagnosticsService(
	repository repository.DiagnosticsRepository,
) DiagnosticsService {
	return &diagnosticsServ{
		diagnosticsRepository: repository,
	}
}

func (service *diagnosticsServ) Save(diagnostic model.Diagnostic) (model.Diagnostic, error) {
	var diagnosticSaved model.Diagnostic
	if diagnosticSaved, err := service.diagnosticsRepository.Save(diagnostic); err != nil {
		return diagnosticSaved, err
	}

	return diagnosticSaved, nil
}

func (service *diagnosticsServ) Find() ([]model.Diagnostic, error) {
	var diagnostics []model.Diagnostic
	diagnostics, err := service.diagnosticsRepository.Find()
	if err != nil {
		return diagnostics, err
	}

	return diagnostics, nil
}
