package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/marsDev10/helpdesk-backend/controllers"
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
	claims, _ := utils.GetUserFromContext(r)

	claimsJSON, _ := json.MarshalIndent(claims, "", "  ")
	fmt.Println(string(claimsJSON))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Si cuentas con los privilegios requeridos",
	})
}
