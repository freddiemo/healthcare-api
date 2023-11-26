package model

import (
	"time"

	"gorm.io/gorm"

	prescriptions "github.com/freddiemo/healthcare-api/internal/prescriptions/model"
)

type Diagnostic struct {
	gorm.Model
	DateTime     *time.Time `binding:"required" gorm:"not null"`
	Diagnostic   string     `binding:"required" gorm:"not null"`
	Prescription prescriptions.Prescription
	PatientID    uint `binding:"required" gorm:"not null"`
}
