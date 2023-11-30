package service

import (
	"github.com/freddiemo/healthcare-api/internal/patients/model"
	"github.com/freddiemo/healthcare-api/internal/patients/repository"
)

type PatientService interface {
	FindById(id uint) (model.Patient, error)
}

type patientServ struct {
	patientRepository repository.PatientRepository
}

func NewPatientService(
	repository repository.PatientRepository,
) PatientService {
	return &patientServ{
		patientRepository: repository,
	}
}

func (service *patientServ) FindById(id uint) (model.Patient, error) {
	var patient model.Patient
	patient, err := service.patientRepository.FindById(id)
	if err != nil {
		return model.Patient{}, err
	}

	return patient, nil
}
