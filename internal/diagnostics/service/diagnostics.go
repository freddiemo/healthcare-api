package service

import (
	"github.com/freddiemo/healthcare-api/internal/diagnostics/model"
	"github.com/freddiemo/healthcare-api/internal/diagnostics/repository"
)

type DiagnosticsService interface {
	Save(model.Diagnostic) (model.Diagnostic, error)
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
	diagnostic, err := service.diagnosticsRepository.Save(diagnostic)
	if err != nil {
		return model.Diagnostic{}, err
	}

	return diagnostic, nil
}
