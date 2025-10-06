package models

import (
	"time"

	"github.com/marsDev10/helpdesk-backend/enums"
)

type TeamMember struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	TeamID    uint           `gorm:"not null;index:idx_user_team,unique" json:"team_id"`
	UserID    uint           `gorm:"not null;index:idx_user_team,unique" json:"user_id"`
	Role      enums.UserRole `gorm:"type:varchar(50);not null;default:'member'" json:"role"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`

	// Relaciones
	Team *Team `gorm:"foreignKey:TeamID" json:"team,omitempty"`
	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// Métodos helper para TeamMember
func (tm *TeamMember) IsManager() bool {
	return tm.Role == enums.Manager
}

func (tm *TeamMember) IsSupervisor() bool {
	return tm.Role == enums.Supervisor
}

func (tm *TeamMember) CanManageTickets() bool {
	return tm.Role == enums.Manager || tm.Role == enums.Supervisor
}

func (tm *TeamMember) IsAdmin() bool {
	return tm.Role == enums.Admin
}
