package service

import (
	"github.com/freddiemo/healthcare-api/internal/diagnostics/model"
	"github.com/freddiemo/healthcare-api/internal/diagnostics/repository"
)

type DiagnosticsService interface {
	Save(*model.Diagnostic) error
	Find(*model.DiagnosticQuery) ([]model.DiagnosticWithPatient, error)
	Delete(id uint) error
	DeleteAll() error
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

func (service *diagnosticsServ) Save(diagnostic *model.Diagnostic) error {
	if err := service.diagnosticsRepository.Save(diagnostic); err != nil {
		return err
	}

	return nil
}

func (service *diagnosticsServ) Find(diagnosticQuery *model.DiagnosticQuery) ([]model.DiagnosticWithPatient, error) {
	var diagnostics []model.DiagnosticWithPatient
	diagnostics, err := service.diagnosticsRepository.Find(diagnosticQuery)
	if err != nil {
		return diagnostics, err
	}

	return diagnostics, nil
}

func (service *diagnosticsServ) Delete(id uint) error {
	if err := service.diagnosticsRepository.Delete(id); err != nil {
		return err
	}

	return nil
}

func (service *diagnosticsServ) DeleteAll() error {
	if err := service.diagnosticsRepository.DeleteAll(); err != nil {
		return err
	}

	return nil
}
