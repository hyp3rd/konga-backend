package konga

import (
	"github.com/google/uuid"
)

// Konga struct
type Konga struct {
	// gorm.Model
	ID uuid.UUID `json:"uuid,omitempty" gorm:"primary_key"`
}

// Repository describes the persistence on konga model
type Repository interface {
}
