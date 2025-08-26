package models

import (
	"time"
)

// Modelo de Categoría
type Category struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description,omitempty" db:"description"`
	Color       *string   `json:"color,omitempty" db:"color"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
