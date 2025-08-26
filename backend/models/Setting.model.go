package models

import (
	"time"
)

type Setting struct {
	ID          int       `json:"id" db:"id"`
	Key         string    `json:"key" db:"key"`
	Value       string    `json:"value" db:"value"`
	Description *string   `json:"description,omitempty" db:"description"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
