package controllers

import (
	"errors"
	"fmt"

	"github.com/marsDev10/helpdesk-backend/db"
	"github.com/marsDev10/helpdesk-backend/dtos"
	"github.com/marsDev10/helpdesk-backend/models"
	"github.com/marsDev10/helpdesk-backend/utils"
	"gorm.io/gorm"
)

// CreateOrganization crea una organización y su usuario admin
func CreateOrganization(dto *dtos.CreateOrganizationDTO) error {
	var existingOrg models.Organization

	// Verificar si el dominio ya existe
	err := db.DB.Where("domain = ?", dto.Domain).First(&existingOrg).Error
	if err == nil {
		return errors.New("dominio no disponible")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Crear organización
	org := models.Organization{
		Name:   dto.Name,
		Domain: &dto.Domain, // *string en el modelo
	}
	if err := db.DB.Create(&org).Error; err != nil {
		return err
	}

	// Hashear contraseña del admin
	hashedPass, err := utils.HashPassword(dto.AdminUser.Password)
	if err != nil {
		return fmt.Errorf("error al hashear contraseña: %v", err)
	}

	// Crear usuario administrador
	adminUser := models.User{
		Email:          dto.AdminUser.Email,
		Password:       hashedPass,
		FirstName:      dto.AdminUser.FirstName,
		LastName:       dto.AdminUser.LastName,
		OrganizationID: org.ID,
	}

	if err := db.DB.Create(&adminUser).Error; err != nil {
		return err
	}

	return nil
}
