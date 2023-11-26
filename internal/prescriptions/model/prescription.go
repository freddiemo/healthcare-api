package model

import (
	"time"

	"gorm.io/gorm"
)

type Prescription struct {
	gorm.Model
	Description  string    `gorm:"not null"`
	DueDate      time.Time `gorm:"not null"`
	DiagnosticID uint
}
