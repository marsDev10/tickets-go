package models

import (
	"time"

	"github.com/marsDev10/helpdesk-backend/enums"
)

type SLAPolicy struct {
	ID                  int                  `json:"id" db:"id"`
	Name                string               `json:"name" db:"name"`
	Priority            enums.TicketPriority `json:"priority" db:"priority"`
	ResponseTimeHours   int                  `json:"response_time_hours" db:"response_time_hours"`
	ResolutionTimeHours int                  `json:"resolution_time_hours" db:"resolution_time_hours"`
	IsActive            bool                 `json:"is_active" db:"is_active"`
	CreatedAt           time.Time            `json:"created_at" db:"created_at"`

	OrganizationID int `json:"organization_id" db:"organization_id"`
}
