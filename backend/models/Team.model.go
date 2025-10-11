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

	// Relacion con la organizacion
	OrganizationID uint          `gorm:"not null" json:"organization_id"`
	Organization   *Organization `gorm:"foreignKey:OrganizationID" json:"organization,omitempty"`

	// Clasificacion del equipo
	CategoryID *uint     `gorm:"index" json:"category_id,omitempty"`
	Category   *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`

	// Relacion many-to-many con usuarios
	Members []TeamMember `gorm:"foreignKey:TeamID" json:"members,omitempty"`
}
