package models

import (
	"time"
)

// Modelo de Ticket
type Ticket struct {
	ID             uint   `gorm:"primaryKey"`
	TicketNumber   string `gorm:"uniqueIndex;not null"`
	Subject        string `gorm:"size:255;not null"`
	Description    string `gorm:"type:text"`
	Status         string `gorm:"size:50;not null"`
	Priority       string `gorm:"size:20"`
	RequesterID    uint   `gorm:"not null"`
	Requester      User   `gorm:"foreignKey:RequesterID"`
	AssigneeID     *uint
	Assignee       User `gorm:"foreignKey:AssigneeID"`
	CategoryID     uint
	Category       Category `gorm:"foreignKey:CategoryID"`
	OrganizationID uint
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	DueDate        *time.Time
	ResolvedAt     *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Conversations  []TicketConversation `gorm:"foreignKey:TicketID;constraint:OnDelete:CASCADE"`
	Attachments    []Attachment         `gorm:"foreignKey:TicketID;constraint:OnDelete:CASCADE"`
	Tags           []Tag                `gorm:"many2many:ticket_tags"`
}
