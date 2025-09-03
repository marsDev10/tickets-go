package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Gender    uint   `gorm:"not null" json:"gender"`
	Email     string `gorm:"not null; unique_index" json:"email"`
	Phone     string `gorm:"not null; unique_index" json:"phone"`
	Password  string `gorm:"not null" json:"password"`
	Role      string `gorm:"not null" json:"role"`
	IsActive  bool   `gorm:"default:true" json:"is_active"`

	OrganizationID int           `gorm:"not null" json:"organization_id"`
	Organization   *Organization `gorm:"foreignKey:OrganizationID" json:"-"`
}

// Método para obtener nombre completo
func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}
