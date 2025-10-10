package models

import (
	"github.com/marsDev10/helpdesk-backend/enums"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Gender    uint   `gorm:"not null" json:"gender"`
	Email     string `gorm:"not null; unique_index" json:"email"`
	Phone     string `gorm:"not null; unique_index" json:"phone"`
	Password  string `gorm:"not null" json:"password"`
	Role      string `gorm:"not null" json:"role"`
	IsActive  bool   `gorm:"default:true" json:"is_active"`

	// Rol a nivel organizacional (para permisos globales)
	GlobalRole enums.UserRole `gorm:"type:varchar(50);default:'member'" json:"global_role"`

	OrganizationID int           `gorm:"not null" json:"organization_id"`
	Organization   *Organization `gorm:"foreignKey:OrganizationID" json:"-"`

	// Relación many-to-many con equipos a través de TeamMember
	TeamMemberships []TeamMember `gorm:"foreignKey:UserID" json:"team_memberships,omitempty"`

	// Relación one-to-many con tickets asignados
	AssignedTickets []Ticket `gorm:"foreignKey:AssigneeID" json:"assigned_tickets,omitempty"`

	// Relación one-to-many con tickets creados (solicitados)
	RequestedTickets []Ticket `gorm:"foreignKey:RequesterID" json:"requested_tickets,omitempty"`

	// Relación one-to-many con tickets creados (creados por)
	CreatedTickets []Ticket `gorm:"foreignKey:CreatedByID" json:"created_tickets,omitempty"`
}

// Método para obtener nombre completo
func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}
