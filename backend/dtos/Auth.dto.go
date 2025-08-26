package dtos

import (
	"time"

	"github.com/marsDev10/helpdesk-backend/models"
)

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponseDto struct {
	Token        string      `json:"token"`
	RefreshToken string      `json:"refresh_token"`
	User         models.User `json:"user"`
	ExpiresAt    time.Time   `json:"expires_at"`
}
