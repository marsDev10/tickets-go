package models

import (
	"time"
)

// Modelo de relación Ticket-Tag
type TicketTag struct {
	TicketID  int       `json:"ticket_id" db:"ticket_id"`
	TagID     int       `json:"tag_id" db:"tag_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
