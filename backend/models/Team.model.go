package models

import (
	"time"
)

// Modelo de Equipos
// Team - Equipo de trabajo
type Team struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relación con la organización
	OrganizationID uint          `gorm:"not null" json:"organization_id"`
	Organization   *Organization `gorm:"foreignKey:OrganizationID" json:"organization,omitempty"`

	// Relación many-to-many con usuarios
	Members []TeamMember `gorm:"foreignKey:TeamID" json:"members,omitempty"`
}
