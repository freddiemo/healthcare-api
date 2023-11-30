package repository

import (
	"strings"

	"gorm.io/gorm"

	"github.com/freddiemo/healthcare-api/internal/diagnostics/model"
)

type DiagnosticsRepository interface {
	Save(*model.Diagnostic) error
	Find(*model.DiagnosticQuery) ([]model.DiagnosticWithPatient, error)
	Delete(id uint) error
	DeleteAll() error
}

type diagnosticsRepo struct {
	db *gorm.DB
}

func NewDiagnosticsRepository(db *gorm.DB) DiagnosticsRepository {
	return &diagnosticsRepo{
		db: db,
	}
}

func (diagnosticRepo *diagnosticsRepo) Save(diagnostic *model.Diagnostic) error {
	if result := diagnosticRepo.db.Save(diagnostic); result.Error != nil {
		return result.Error
	}

	return nil
}

func (diagnosticRepo *diagnosticsRepo) Find(diagnosticQuery *model.DiagnosticQuery) ([]model.DiagnosticWithPatient, error) {
	queriesPatient := diagnosticRepo.db.Where("TRUE")
	if diagnosticQuery.FirstName != nil {
		queriesPatient = queriesPatient.Where("UPPER(first_name) = ?", strings.ToUpper(*diagnosticQuery.FirstName))
	}
	if diagnosticQuery.LastName != nil {
		queriesPatient = queriesPatient.Where("UPPER(last_name) = ?", strings.ToUpper(*diagnosticQuery.LastName))
	}
	queryDiagnostic := diagnosticRepo.db.Where("TRUE")
	if diagnosticQuery.Date != nil {
		queryDiagnostic = queryDiagnostic.Where("date_time BETWEEN ? AND ? ", diagnosticQuery.Date, diagnosticQuery.Date.AddDate(0, 0, 1))
	}

	var diagnostics []model.DiagnosticWithPatient
	if result := diagnosticRepo.db.Table("diagnostics").Preload("Prescription").InnerJoins("Patient", queriesPatient).Where(queryDiagnostic).Find(&diagnostics); result.Error != nil {
		return diagnostics, result.Error
	}

	return diagnostics, nil
}

func (diagnosticRepo *diagnosticsRepo) Delete(id uint) error {
	var diagnostic model.Diagnostic
	if result := diagnosticRepo.db.Delete(&diagnostic, id); result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (diagnosticRepo *diagnosticsRepo) DeleteAll() error {
	if result := diagnosticRepo.db.Exec("DELETE FROM diagnostics"); result.Error != nil {
		return result.Error
	}

	return nil
}
