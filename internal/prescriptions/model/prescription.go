package model

import (
	"time"

	"gorm.io/gorm"
)

type Prescription struct {
	ID           uint           `gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Description  string         `binding:"omitempty" gorm:"not null"`
	DueDate      time.Time      `binding:"omitempty" gorm:"not null"`
	DiagnosticID uint
}
