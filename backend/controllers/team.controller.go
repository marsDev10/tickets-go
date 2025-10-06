package controllers

import (
	"errors"
	"fmt"

	"github.com/marsDev10/helpdesk-backend/db"
	"github.com/marsDev10/helpdesk-backend/models"
	"gorm.io/gorm"
)

func GetTeamByOrganization(orgID int, page, limit int, search string) ([]models.Team, int64, error) {

	query := db.DB.Where("organization_id = ?", orgID)

	// Search Filter
	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("LOWER(name) LIKE ?", searchPattern)
	}

	// Total couunt
	var total int64

	if err := query.Model(&models.Team{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get teams with pagination

	var teams []models.Team

	offset := (page - 1) * limit

	err := query.
		Preload("Members").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&teams).Error

	return teams, total, err
}

// GeTeamByID get team by id
func GetTeamByID(orgID int, teamID int) (*models.Team, error) {
	var team models.Team

	err := db.DB.
		Preload("Members").
		Where("organization_id = ? AND id = ?", orgID, teamID).
		First(&team).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("equipo no encontrado")
		}
		return nil, err
	}

	return &team, nil
}

// Create Team
func CreateTeam(teamID, orgID int, name, description string) (*models.Team, error) {

	var existing models.Team

	err := db.DB.
		Where("organization_id = ? AND name = ?", orgID, name).First(&existing).Error

	if err == nil {
		return nil, fmt.Errorf("Ya existe un equipo con el nombre '%s' en esta organización", name)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Create team
	desc := &description
	team := models.Team{
		Name:           name,
		Description:    desc,
		OrganizationID: uint(orgID),
	}

	if err := db.DB.Create(&team).Error; err != nil {
		return nil, err
	}

	return &team, nil
}
