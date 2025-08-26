package models

import (
	"time"
)

// Modelo de Etiquetas
type Tag struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Color     *string   `json:"color,omitempty" db:"color"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
