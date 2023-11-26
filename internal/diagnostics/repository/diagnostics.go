package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/healthcare-api/internal/diagnostics/model"
)

type DiagnosticsRepository interface {
	Save(diagnostic model.Diagnostic) (model.Diagnostic, error)
}

type diagnosticsRepo struct {
	db *gorm.DB
}

func NewDiagnosticsRepository(db *gorm.DB) DiagnosticsRepository {
	return &diagnosticsRepo{
		db: db,
	}
}

func (diagnosticRepo *diagnosticsRepo) Save(diagnostic model.Diagnostic) (model.Diagnostic, error) {
	if result := diagnosticRepo.db.Save(&diagnostic); result.Error != nil {
		return model.Diagnostic{}, result.Error
	}

	return diagnostic, nil
}
