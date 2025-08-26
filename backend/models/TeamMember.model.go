package models

import (
	"time"

	"github.com/marsDev10/helpdesk-backend/enums"
)

type TeamMember struct {
	ID        int            `json:"id" db:"id"`
	TeamID    int            `json:"team_id" db:"team_id"`
	UserID    int            `json:"user_id" db:"user_id"`
	Role      enums.UserRole `json:"role" db:"role"` // e.g., "admin", "member"
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
}
