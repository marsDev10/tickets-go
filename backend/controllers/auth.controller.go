package controllers

import (
	"errors"
	"net/http"
	"os"

	"github.com/marsDev10/helpdesk-backend/db"
	"github.com/marsDev10/helpdesk-backend/dtos"
	"github.com/marsDev10/helpdesk-backend/models"
	"github.com/marsDev10/helpdesk-backend/utils"
	"gorm.io/gorm"
)

func Login(dto *dtos.LoginDto) (map[string]interface{}, error) {
	var user models.User
	if err := db.DB.Where("email = ?", dto.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}

	// Comparar contraseñas
	if !utils.CheckPassword(dto.Password, user.Password) {
		return nil, errors.New("contraseña inválida")
	}

	// Generar JWT
	secret := os.Getenv("JWT_SECRET_KEY")
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role, secret, 24)
	if err != nil {
		return nil, err
	}

	// Retornar datos del usuario + token
	return map[string]interface{}{
		"user": map[string]interface{}{
			"id":         user.ID,
			"first_name": user.FirstName,
			"email":      user.Email,
			"role":       user.Role,
		},
		"token": token,
	}, nil
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Lógica para manejar el registro
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("RegisterHandler no implementado"))
}
