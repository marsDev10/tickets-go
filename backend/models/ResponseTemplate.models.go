package models

import (
	"time"
)

type ResponseTemplate struct {
	ID         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Subject    *string   `json:"subject,omitempty" db:"subject"`
	Content    string    `json:"content" db:"content"`
	CategoryID *int      `json:"category_id,omitempty" db:"category_id"`
	Category   *Category `json:"category,omitempty" db:"-"`
	CreatedBy  int       `json:"created_by" db:"created_by"`
	Creator    *User     `json:"creator,omitempty" db:"-"`
	IsActive   bool      `json:"is_active" db:"is_active"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
