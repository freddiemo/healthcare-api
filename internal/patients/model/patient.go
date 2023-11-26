package model

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	FirstName  string `gorm:"type:varchar(32);not null"`
	LastName   string `gorm:"type:varchar(32);not null"`
	Email      string `gorm:"type:varchar(32);not null;UNIQUE" validate:"is-email"`
	DocumentID string `gorm:"not null;UNIQUE"`
	Phone      string `gorm:"not null"`
}
