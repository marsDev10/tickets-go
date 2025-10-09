package controllers

import (
    "errors"
    "fmt"

    "github.com/marsDev10/helpdesk-backend/db"
    "github.com/marsDev10/helpdesk-backend/dtos"
    "github.com/marsDev10/helpdesk-backend/enums"
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
		Preload("Members.User").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&teams).Error

	return teams, total, err
}

func GetTeamsByOrganization(orgID int) ([]dtos.TeamMembersByOrganizationResponse, error) {
	var teams []models.Team

	err := db.DB.
		Model(&models.Team{}).
		Where("organization_id = ?", orgID).
		Preload("Members.User").
		Find(&teams).Error

	if err != nil {
		return nil, err
	}

	var result []dtos.TeamMembersByOrganizationResponse
	for _, t := range teams {
		var members []dtos.MemberBasic
		for _, m := range t.Members {
			members = append(members, dtos.MemberBasic{
				Role:      string(m.Role),
				ID:        uint(m.User.ID),
				FirstName: m.User.FirstName,
				LastName:  m.User.LastName,
				Email:     m.User.Email,
			})
		}

		result = append(result, dtos.TeamMembersByOrganizationResponse{
			ID:      t.ID,
			Name:    t.Name,
			Members: members,
		})
	}

	return result, nil
}

// Create Team
func CreateTeam(orgID int, name, description string) (*models.Team, error) {

	var existing models.Team

	err := db.DB.
		Where("organization_id = ? AND name = ?", orgID, name).First(&existing).Error

	if err == nil {
		return nil, fmt.Errorf("ya existe un equipo con el nombre '%s' en esta organización", name)
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

// UpdateTeam actualiza campos del equipo (nombre, descripción) dentro de la misma organización
func UpdateTeam(orgID, teamID int, dto dtos.UpdateTeamDto) (*models.Team, error) {
    var team models.Team

    if err := db.DB.Where("id = ? AND organization_id = ?", teamID, orgID).First(&team).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("equipo no encontrado")
        }
        return nil, err
    }

    // Validar unicidad del nombre si se envía uno nuevo
    if dto.Name != nil {
        var exists models.Team
        if err := db.DB.
            Where("organization_id = ? AND LOWER(name) = LOWER(?) AND id <> ?", orgID, *dto.Name, teamID).
            First(&exists).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, err
        } else if err == nil {
            return nil, fmt.Errorf("ya existe un equipo con el nombre '%s' en esta organización", *dto.Name)
        }
        team.Name = *dto.Name
    }

    if dto.Description != nil {
        // Permite setear a vacío si viene "" o a nil si se requiere omitir
        desc := *dto.Description
        team.Description = &desc
        if desc == "" {
            // Si prefieres que vacío signifique NULL, descomenta:
            // team.Description = nil
        }
    }

    if err := db.DB.Save(&team).Error; err != nil {
        return nil, err
    }

    return &team, nil
}

// AddMemberToTeam agrega un usuario a un equipo con un rol específico
func AddMemberToTeam(teamID, userID, orgID int, role enums.UserRole) error {
	// Validar que el rol sea válido para equipos
	if !role.IsTeamRole() {
		return errors.New("rol inválido para equipo")
	}

	// Verificar que el equipo existe y pertenece a la organización
	var team models.Team
	if err := db.DB.Where("id = ? AND organization_id = ?", teamID, orgID).First(&team).Error; err != nil {
		return errors.New("equipo no encontrado")
	}

	// Verificar que el usuario existe y pertenece a la misma organización
	var user models.User
	if err := db.DB.Where("id = ? AND organization_id = ?", userID, orgID).First(&user).Error; err != nil {
		return errors.New("usuario no encontrado en la organización")
	}

	// Verificar que el usuario no esté ya en el equipo
	var existingMember models.TeamMember
	err := db.DB.Where("team_id = ? AND user_id = ?", teamID, userID).First(&existingMember).Error

	if err == nil {
		return errors.New("el usuario ya es miembro del equipo")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Crear la membresía
	member := models.TeamMember{
		TeamID: uint(teamID),
		UserID: uint(userID),
		Role:   role,
	}

	return db.DB.Create(&member).Error
}

// DeleteMemberToTeam
func RemoveMemberFromTeam(teamID, userID, orgID int) error {
    // Verificar que el equipo existe y pertenece a la organización
    var team models.Team
    if err := db.DB.Where("id = ? AND organization_id = ?", teamID, orgID).First(&team).Error; err != nil {
        return errors.New("equipo no encontrado")
    }

    // Verificar que el usuario existe y pertenece a la organización
    var user models.User
    if err := db.DB.Where("id = ? AND organization_id = ?", userID, orgID).First(&user).Error; err != nil {
        return errors.New("usuario no encontrado en la organización")
    }

    // Eliminar la membresía específica del equipo
    result := db.DB.Where("team_id = ? AND user_id = ?", teamID, userID).
        Delete(&models.TeamMember{})

    if result.Error != nil {
        return fmt.Errorf("error al eliminar membresía del equipo: %v", result.Error)
    }

    if result.RowsAffected == 0 {
        return errors.New("el usuario no pertenece a este equipo")
    }

    return nil
}
