package models

import (
	"time"
)

// Modelo de Categoría
type Category struct {
	ID             uint      `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description,omitempty" db:"description"`
	Color          *string   `json:"color,omitempty" db:"color"`
	OrganizationID uint      `json:"organization_id" db:"organization_id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
