package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marsDev10/helpdesk-backend/db"
	"github.com/marsDev10/helpdesk-backend/dtos"
	"github.com/marsDev10/helpdesk-backend/enums"
	"github.com/marsDev10/helpdesk-backend/models"
	"github.com/marsDev10/helpdesk-backend/utils"
)

func CreateOrganization(c *gin.Context) {
	var dto dtos.CreateOrganizationDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Iniciar transacción
	tx := db.DB.Begin()

	// 1. Crear organización
	org := models.Organization{
		Name:   dto.Name,
		Domain: &dto.Domain,
	}

	if err := tx.Create(&org).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Error creating organization"})
		return
	}

	// 2. Crear usuario administrador
	hashedPassword, err := utils.HashPassword(dto.AdminUser.Password)
	if err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Error hashing password"})
		return
	}

	admin := models.User{
		Email:          dto.AdminUser.Email,
		Password:       hashedPassword,
		FirstName:      dto.AdminUser.FirstName,
		LastName:       dto.AdminUser.LastName,
		Role:           string(enums.RoleAdmin),
		OrganizationID: org.ID,
	}

	if err := tx.Create(&admin).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Error creating admin user"})
		return
	}

	// 3. Crear categorías por defecto

	// 4. Crear configuración SLA por defecto
	defaultSLA := models.SLAPolicy{
		Name:           "Default SLA",
		Priority:       "Normal",
		OrganizationID: org.ID,
	}

	if err := tx.Create(&defaultSLA).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Error creating default SLA"})
		return
	}

	// Confirmar transacción
	tx.Commit()

	c.JSON(201, gin.H{
		"message":      "Organization created successfully",
		"organization": org,
	})
}
