package models

import (
	"time"
)

// Modelo de Equipos
type Team struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`

	// Miembros del equipo
	Members []User `json:"members,omitempty" db:"-"`
}
