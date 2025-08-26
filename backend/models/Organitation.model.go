package models

import (
	"time"
)

// Modelo de Organización
type Organization struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Domain    *string   `json:"domain,omitempty" db:"domain"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	Users []User `json:"users,omitempty" gorm:"foreignKey:OrganizationID"`
}
