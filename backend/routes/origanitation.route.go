package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/marsDev10/helpdesk-backend/controllers"
	"github.com/marsDev10/helpdesk-backend/dtos"
)

// CreateOrganizationHandler maneja la creación de organizaciones con usuario admin
func CreateOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateOrganizationDTO

	// 1️⃣ Decodificar JSON
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, fmt.Sprintf("JSON inválido: %v", err), http.StatusBadRequest)
		return
	}

	// 2️⃣ Validación de campos obligatorios
	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":  "Error de validación",
			"detail": err.Error(),
		})
		return
	}

	// 3️⃣ Llamar al servicio que crea la organización + admin
	if err := controllers.CreateOrganization(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":  "No se pudo crear la organización",
			"detail": err.Error(),
		})
		return
	}

	// 4️⃣ Respuesta exitosa
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Organización creada correctamente",
	})
}
