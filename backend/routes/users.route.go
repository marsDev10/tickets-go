package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/marsDev10/helpdesk-backend/controllers"
	"github.com/marsDev10/helpdesk-backend/dtos"
	"github.com/marsDev10/helpdesk-backend/utils"
)

func GetOrganizationUsersHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := utils.GetUserFromContext(r)

	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("No autorizado", err.Error()))
		return
	}

	claimsJSON, _ := json.MarshalIndent(claims, "", "  ")
	fmt.Println(string(claimsJSON))

	// 3️⃣ Obtener parámetros opcionales para filtrado/paginación
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	search := r.URL.Query().Get("search")
	role := r.URL.Query().Get("role") // Filtrar por rol si es necesario

	// 4️⃣ Convertir parámetros de paginación
	pageInt := 1
	limitInt := 10

	if page != "" {
		if p, err := strconv.Atoi(page); err == nil && p > 0 {
			pageInt = p
		}
	}

	if limit != "" {
		if l, err := strconv.Atoi(limit); err == nil && l > 0 && l <= 100 {
			limitInt = l
		}
	}

	// 5️⃣ Verificar que la organización existe
	orgExists, err := controllers.OrganizationExists(claims.OrganizationID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Error al verificar la organización",
		})
		return
	}

	if !orgExists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Organización no encontrada",
		})
		return
	}

	// 6️⃣ Obtener usuarios de la organización
	users, total, err := controllers.GetUsersByOrganization(claims.OrganizationID, pageInt, limitInt, search, role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Error al obtener usuarios de la organización",
		})
		return
	}

	// 7️⃣ Respuesta exitosa
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    users,
		"pagination": map[string]interface{}{
			"page":        pageInt,
			"limit":       limitInt,
			"total":       total,
			"total_pages": (total + int64(limitInt) - 1) / int64(limitInt),
		},
	})
}

func GetOrganizationUserHandler(w http.ResponseWriter, r *http.Request) {

	claims, err := utils.GetUserFromContext(r)

	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("No autorizado", err.Error()))
		return
	}

	vars := mux.Vars(r)
	userIDStr := vars["idUser"] // o vars["org_id"] dependiendo de tu ruta

	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "ID de organización inválido, debe ser un número positivo",
		})
		return
	}

	// 5️⃣ Verificar que la organización existe
	userExist, err := controllers.GetUserByOrganization(claims.OrganizationID, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Error al verificar el usuario",
		})
		return
	}

	if userExist == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Usuario no encontrado en la organización",
		})
		return
	}

	// 7️⃣ Respuesta exitosa
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"user":    userExist,
	})

}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateUserDto

	// Decodificar el body en dto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":  "Datos inválidos",
			"detail": err.Error(),
		})
		return
	}

	// ✅ Validación con validator
	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":  "Validación fallida",
			"detail": err.Error(),
		})
		return
	}

	// Obtener claims (usuario autenticado)
	claims, _ := utils.GetUserFromContext(r)

	claimsJSON, _ := json.MarshalIndent(claims, "", "  ")
	fmt.Println(string(claimsJSON))

	// Pasar dto + orgID al servicio
	user, err := controllers.CreateUser(&dto, claims.OrganizationID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":  "No se pudo crear el usuario",
			"detail": err.Error(),
		})
		return
	}

	// Respuesta correcta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Usuario creado correctamente",
		"user":    user, // o bien devolver el modelo `User` que guardaste
	})
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var dto dtos.UpdateUserDto

	// Imprimir la data de manera formateada
	dtoJSON, _ := json.MarshalIndent(&dto, "", "    ")
	log.Printf("📝 Datos recibidos para actualización:\n%s", string(dtoJSON))

	// Decodificar el body en dto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":  "Datos inválidos",
			"detail": err.Error(),
		})
		return
	}

	// ✅ Validación con validator
	validate := validator.New()
	if err := validate.Struct(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":  "Validación fallida",
			"detail": err.Error(),
		})
		return
	}

	// Obtener claims (usuario autenticado)
	claims, err := utils.GetUserFromContext(r)
	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("No autorizado", err.Error()))
		return
	}

	// Pasar dto + orgID al servicio
	user, err := controllers.UpdateUser(&dto, claims.OrganizationID)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Error al actualizar", err.Error()))
		return
	}

	// Respuesta correcta
	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Usuario actualizado correctamente",
		"user":    user,
	})
}

func ToggleUserStatusHandler(w http.ResponseWriter, r *http.Request) {
	var dto dtos.ToggleStatusUser

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	claims, err := utils.GetUserFromContext(r)
	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("No autorizado", err.Error()))
		return
	}

	if err := controllers.ToggleUserStatus(dto.ID, claims.OrganizationID); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Error al cambiar estado", err.Error()))
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Estado del usuario actualizado correctamente",
	})
}
