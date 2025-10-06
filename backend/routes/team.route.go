package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/marsDev10/helpdesk-backend/controllers"
	"github.com/marsDev10/helpdesk-backend/dtos"
	"github.com/marsDev10/helpdesk-backend/enums"
	"github.com/marsDev10/helpdesk-backend/utils"
)

func CreateTeamHandler(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateTeamDto

	// Decodificar body
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	// Validar DTO
	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Validación fallida", err.Error()))
		return
	}

	// Obtener claims
	claims, err := utils.GetUserFromContext(r)
	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("No autorizado", err.Error()))
		return
	}

	// Crear equipo
	team, err := controllers.CreateTeam(claims.OrganizationID, dto.Name, dto.Description)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Error al crear equipo", err.Error()))
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Equipo creado exitosamente",
		"data":    team,
	})
}

func AddMemberToTeamHandler(w http.ResponseWriter, r *http.Request) {
	var dto dtos.AddTeamMemberDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Validación fallida", err.Error()))
		return
	}

	claims, err := utils.GetUserFromContext(r)
	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("No autorizado", err.Error()))
		return
	}

	vars := mux.Vars(r)
	teamID, err := strconv.Atoi(vars["team_id"])
	if err != nil || teamID <= 0 {
		utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("ID inválido", ""))
		return
	}

	// Convertir string a enum
	role := enums.UserRole(dto.Role)

	err = controllers.AddMemberToTeam(teamID, dto.UserID, claims.OrganizationID, role)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Error al agregar miembro", err.Error()))
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Miembro agregado exitosamente",
	})
}
