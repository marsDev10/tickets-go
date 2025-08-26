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
	AssigneeID  *int                  `json:"assignee_id,omitempty"`
	DueDate     *time.Time            `json:"due_date,omitempty"`
	Tags        []int                 `json:"tags,omitempty"` // Array de IDs de tags
}

type UpdateTicketDto struct {
	Subject     *string               `json:"subject,omitempty"`
	Description *string               `json:"description,omitempty"`
	Status      *enums.TicketStatus   `json:"status,omitempty"`
	Priority    *enums.TicketPriority `json:"priority,omitempty"`
	CategoryID  *int                  `json:"category_id,omitempty"`
	AssigneeID  *int                  `json:"assignee_id,omitempty"`
	DueDate     *time.Time            `json:"due_date,omitempty"`
}
