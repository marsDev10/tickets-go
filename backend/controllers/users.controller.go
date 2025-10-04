package controllers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/marsDev10/helpdesk-backend/db"
	"github.com/marsDev10/helpdesk-backend/dtos"
	"github.com/marsDev10/helpdesk-backend/models"
	"github.com/marsDev10/helpdesk-backend/utils"
	"gorm.io/gorm"
)

// Función auxiliar para obtener usuarios por organización (sin HTTP)
func GetUsersByOrganization(orgID int, page, limit int, search, role string) ([]models.User, int64, error) {
	// Construir la query base
	query := db.DB.Where("organization_id = ?", strconv.Itoa(orgID))

	//query := db.DB.Model(&models.User{}).Find(&dtos.GetUsersData{})

	// Aplicar filtro de búsqueda si existe
	if search != "" {
		searchPattern := "%" + strings.ToLower(search) + "%"
		query = query.Where(
			"LOWER(first_name) LIKE ? OR LOWER(last_name) LIKE ? OR LOWER(email) LIKE ?",
			searchPattern, searchPattern, searchPattern,
		)
	}

	// Aplicar filtro de rol si existe
	if role != "" {
		query = query.Where("role = ?", role)
	}

	// Contar el total
	var total int64
	if err := query.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Obtener usuarios con paginación
	var users []models.User
	offset := (page - 1) * limit

	err := query.
		Select("id, first_name, last_name, email, role, organization_id").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&users).Error

	return users, total, err
}

func GetUserByOrganization(orgID int, userID int) (*dtos.UserResponse, error) {

	if orgID <= 0 {
		return nil, errors.New("ID de organización inválido")
	}
	if userID <= 0 {
		return nil, errors.New("ID de usuario inválido")
	}

	var userResponse dtos.UserResponse

	err := db.DB.Model(&models.User{}).
		Select("id, first_name, last_name, email, role, organization_id").
		Where("organization_id = ? AND id = ?", orgID, userID).
		First(&userResponse).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuario no encontrado en la organización especificada")
		}
		return nil, errors.New("error al buscar usuario: " + err.Error())
	}

	return &userResponse, nil
}

func CreateUser(dto *dtos.CreateUserDto, orgID int) (*models.User, error) {
	var existingUser models.User

	// Verificar si ya existe el email
	err := db.DB.Where("email = ?", dto.Email).First(&existingUser).Error

	if err == nil {
		return nil, fmt.Errorf("el correo %s ya está registrado", dto.Email)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Hashear contraseña
	hashedPass, err := utils.HashPassword(dto.Password)
	if err != nil {
		return nil, fmt.Errorf("error al hashear contraseña: %v", err)
	}

	genderMap := map[string]uint{
		"male":   1,
		"female": 2,
		"other":  3,
	}

	genderVal := genderMap[dto.Gender]

	// Crear la entidad con el modelo
	user := models.User{
		FirstName:      dto.FirstName,
		LastName:       dto.LastName,
		Gender:         genderVal,
		Email:          dto.Email,
		Phone:          dto.Phone,
		Password:       hashedPass,
		Role:           dto.Role,
		IsActive:       true, // por defecto
		OrganizationID: orgID,
	}

	// Guardar en DB
	if err := db.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
