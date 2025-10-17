package routes

import (
	"net/http"

	"github.com/marsDev10/helpdesk-backend/controllers"
	"github.com/marsDev10/helpdesk-backend/dtos"
	"github.com/marsDev10/helpdesk-backend/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var dto dtos.LoginDto
	if err := utils.ParseJSON(r, &dto); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	data, err := controllers.Login(&dto)
	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("Credenciales incorrectas", err.Error()))
		return
	}

	utils.JSONResponse(w, http.StatusOK, utils.SuccessResponse("Hola, ¡Bienvenido! 😁👌", data))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Lógica para manejar el registro
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("RegisterHandler no implementado"))
}
