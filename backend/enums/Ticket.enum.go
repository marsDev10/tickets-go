package enums

import (
	"database/sql/driver"
	"fmt"
)

type TicketStatus string
type TicketPriority string

const (
	// Estados de ticket
	StatusOpen     TicketStatus = "open"
	StatusPending  TicketStatus = "pending"
	StatusResolved TicketStatus = "resolved"
	StatusClosed   TicketStatus = "closed"

	// Prioridades
	PriorityLow    TicketPriority = "low"
	PriorityMedium TicketPriority = "medium"
	PriorityHigh   TicketPriority = "high"
	PriorityUrgent TicketPriority = "urgent"
)

// Implementar interfaces para los enums (igual que antes)
func (ts *TicketStatus) Scan(value interface{}) error {
	if value == nil {
		*ts = StatusOpen
		return nil
	}
	if str, ok := value.(string); ok {
		*ts = TicketStatus(str)
		return nil
	}
	return fmt.Errorf("cannot scan %T into TicketStatus", value)
}

func (ts TicketStatus) Value() (driver.Value, error) {
	return string(ts), nil
}

func (tp *TicketPriority) Scan(value interface{}) error {
	if value == nil {
		*tp = PriorityMedium
		return nil
	}
	if str, ok := value.(string); ok {
		*tp = TicketPriority(str)
		return nil
	}
	return fmt.Errorf("cannot scan %T into TicketPriority", value)
}

func (tp TicketPriority) Value() (driver.Value, error) {
	return string(tp), nil
}
