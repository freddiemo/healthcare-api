package seeders

import (
	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"

	"github.com/freddiemo/healthcare-api/internal/patients/model"
)

const defaultSize uint8 = 10

type PatientSeeder interface {
	Populate() error
}

type patientSeeder struct {
	db   *gorm.DB
	size uint8
}

func NewPatientsSeeder(db *gorm.DB) PatientSeeder {
	return &patientSeeder{
		db: db,
	}
}

func (patientSeeder *patientSeeder) Populate() error {
	patients := [defaultSize]model.Patient{}
	for i := 0; i < int(defaultSize); i++ {
		patient := model.Patient{}
		patient.FirstName = faker.FirstName()
		patient.LastName = faker.LastName()
		patient.Email = faker.Email()
		patient.DocumentID = faker.CCNumber()
		patient.Phone = faker.E164PhoneNumber()

		patients[i] = patient
	}

	if result := patientSeeder.db.Model(&model.Patient{}).Create(&patients); result.Error != nil {
		return result.Error
	}

	return nil
}
