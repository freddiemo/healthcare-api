package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/healthcare-api/internal/patients/model"
)

type PatientRepository interface {
	FindById(int) (model.Patient, error)
}

type patientRepo struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepo{
		db: db,
	}
}

func (patientRepo *patientRepo) FindById(id int64) (model.Patient, error) {
	var patient model.Patient
	if result := patientRepo.db.First(&patient, id); result.Error != nil {
		return patient, result.Error
	}

	return patient, nil
}
