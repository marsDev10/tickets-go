package models

import "time"

type TicketConversation struct {
	ID          uint   `gorm:"primaryKey"`
	TicketID    uint   `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
	User        User   `gorm:"foreignKey:UserID"`
	Message     string `gorm:"type:text;not null"`
	IsPublic    bool   `gorm:"default:true"`
	CreatedAt   time.Time
	Attachments []Attachment `gorm:"foreignKey:ConversationID;constraint:OnDelete:CASCADE"`
}
