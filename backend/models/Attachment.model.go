package models

import (
	"time"
)

// Modelo de Archivos Adjuntos
type Attachment struct {
	ID             uint   `gorm:"primaryKey"`
	TicketID       uint   `gorm:"index"`
	ConversationID uint   `gorm:"index"`
	Filename       string `gorm:"size:255;not null"`
	FilePath       string `gorm:"size:255;not null"`
	FileSize       int64  `gorm:"not null"`
	MimeType       string `gorm:"size:100"`
	UploaderID     uint   `gorm:"not null"`
	Uploader       User   `gorm:"foreignKey:UploaderID;references:ID"` // Changed to non-pointer
	CreatedAt      time.Time
}
