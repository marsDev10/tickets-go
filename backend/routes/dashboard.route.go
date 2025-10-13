package routes

import (
	"net/http"

	"github.com/marsDev10/helpdesk-backend/controllers"
	"github.com/marsDev10/helpdesk-backend/utils"
)

func GetDashboardSummaryHandler(w http.ResponseWriter, r *http.Request) {

	claims, err := utils.GetUserFromContext(r)

	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("No autorizado", err.Error()))
		return
	}

	summary, err := controllers.GetDashboardSummary(claims.OrganizationID)

	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, utils.ErrorResponse("Error al obtener resumen", err.Error()))
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Resumen obtenido exitosamente",
		"data":    summary,
	})
}
