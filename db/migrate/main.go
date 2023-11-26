package main

import (
	"log"

	"github.com/freddiemo/healthcare-api/config"
	"github.com/freddiemo/healthcare-api/db"

	patientsSeeder "github.com/freddiemo/healthcare-api/db/seeders"
	patients "github.com/freddiemo/healthcare-api/internal/patients/model"
)

func main() {
	params := config.Init()
	db := db.Init(params)

	db.AutoMigrate(&patients.Patient{})
	var countPatients int64
	if result := db.Model(&patients.Patient{}).Count(&countPatients); result.Error != nil {
		log.Fatal(result.Error)
	} else if countPatients == 0 {
		patientSeeder := patientsSeeder.NewPatientsSeeder(db)
		if err := patientSeeder.Populate(); err != nil {
			log.Fatal(err)
		}
	}
}
