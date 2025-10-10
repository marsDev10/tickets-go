package dtos

import (
	"time"

	"github.com/marsDev10/helpdesk-backend/enums"
)

type CreateTicketDto struct {
	Subject     string                `json:"subject" validate:"required,min=5,max=255"`
	Description string                `json:"description" validate:"required,min=10"`
	Priority    *enums.TicketPriority `json:"priority,omitempty"`
	CategoryID  *int                  `json:"category_id,omitempty"`
	// Deprecated en creación: los tickets se crean sin asignado
	AssigneeID *int `json:"assignee_id,omitempty"`
	// Permite crear en nombre de otra persona (si no se envía, usa el usuario autenticado)
	RequesterID *int `json:"requester_id,omitempty"`
	// Opcional: sugerir equipo destino (se valida en asignación)
	TeamID  *int       `json:"team_id,omitempty"`
	DueDate *time.Time `json:"due_date,omitempty"`
	Tags    []int      `json:"tags,omitempty"` // Array de IDs de tags
}

type UpdateTicketDto struct {
	Subject     *string               `json:"subject,omitempty"`
	Description *string               `json:"description,omitempty"`
	Status      *enums.TicketStatus   `json:"status,omitempty"`
	Priority    *enums.TicketPriority `json:"priority,omitempty"`
	CategoryID  *int                  `json:"category_id,omitempty"`
	// Para asignación usar AssignTicketDto y endpoint dedicado
	AssigneeID *int       `json:"assignee_id,omitempty"`
	DueDate    *time.Time `json:"due_date,omitempty"`
}

// Payload para asignar un ticket por un manager/supervisor del equipo
type AssignTicketDto struct {
	TeamID     int `json:"team_id" validate:"required,gt=0"`
	AssigneeID int `json:"assignee_id" validate:"required,gt=0"`
}

type TicketResponseDto struct {
	ID           uint   `json:"id"`
	TicketNumber string `json:"ticket_number"`
	Subject      string `json:"subject"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	Priority     string `json:"priority"`
	RequesterID  uint   `json:"requester_id"`
	CreatedByID  uint   `json:"created_by_id"`
	AssigneeID   *uint  `json:"assignee_id,omitempty"`
	TeamID       *uint  `json:"team_id,omitempty"`
}
