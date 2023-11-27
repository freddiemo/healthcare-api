package model

import (
	"time"

	"gorm.io/gorm"

	"github.com/freddiemo/healthcare-api/internal/patients/model"
	prescriptions "github.com/freddiemo/healthcare-api/internal/prescriptions/model"
)

type Diagnostic struct {
	ID           uint                       `gorm:"primaryKey"`
	CreatedAt    time.Time                  `json:"-"`
	UpdatedAt    time.Time                  `json:"-"`
	DeletedAt    gorm.DeletedAt             `gorm:"index" json:"-"`
	DateTime     *time.Time                 `binding:"required" gorm:"not null"`
	Diagnostic   string                     `binding:"required" gorm:"not null"`
	Prescription prescriptions.Prescription `gorm:"foreignKey:DiagnosticID;references:ID"`
	PatientID    uint                       `binding:"required" gorm:"not null"`
}

type DiagnosticWithPatient struct {
	ID           uint `gorm:"primaryKey"`
	DateTime     *time.Time
	Diagnostic   string
	Prescription prescriptions.Prescription `gorm:"foreignKey:DiagnosticID;references:ID"`
	PatientID    uint                       `json:"-"`
	Patient      model.Patient
}

type DiagnosticQuery struct {
	FirstName *string    `form:"firstName" binding:"omitempty,min=3,max=20"`
	LastName  *string    `form:"lastName" binding:"omitempty,min=3,max=20"`
	Date      *time.Time `form:"date" time_format:"2006-01-02"`
}
