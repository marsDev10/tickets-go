package routes

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-playground/validator/v10"
    "github.com/gorilla/mux"
    "github.com/marsDev10/helpdesk-backend/controllers"
    "github.com/marsDev10/helpdesk-backend/dtos"
    "github.com/marsDev10/helpdesk-backend/utils"
)

func CreateTicketHandler(w http.ResponseWriter, r *http.Request) {
    var dto dtos.CreateTicketDto
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

    ticket, err := controllers.CreateTicket(claims.OrganizationID, claims.UserID, dto)
    if err != nil {
        utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Error al crear ticket", err.Error()))
        return
    }

    utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
        "success": true,
        "message": "Ticket creado exitosamente",
        "data":    ticket,
    })
}

func GetTicketHandler(w http.ResponseWriter, r *http.Request) {
    claims, err := utils.GetUserFromContext(r)
    if err != nil {
        utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("No autorizado", err.Error()))
        return
    }

    idStr := mux.Vars(r)["ticket_id"]
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("ticket_id inválido", ""))
        return
    }

    ticket, err := controllers.GetTicketByID(claims.OrganizationID, id)
    if err != nil {
        utils.JSONResponse(w, http.StatusNotFound, utils.ErrorResponse("No encontrado", err.Error()))
        return
    }

    utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
        "success": true,
        "message": "Ticket obtenido",
        "data":    ticket,
    })
}

func ListTicketsHandler(w http.ResponseWriter, r *http.Request) {
    claims, err := utils.GetUserFromContext(r)
    if err != nil {
        utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("No autorizado", err.Error()))
        return
    }

    q := r.URL.Query()
    page, _ := strconv.Atoi(q.Get("page"))
    limit, _ := strconv.Atoi(q.Get("limit"))
    status := q.Get("status")
    search := q.Get("search")

    var assigneeID, requesterID, categoryID *int
    if v := q.Get("assignee_id"); v != "" {
        if i, err := strconv.Atoi(v); err == nil {
            assigneeID = &i
        }
    }
    if v := q.Get("requester_id"); v != "" {
        if i, err := strconv.Atoi(v); err == nil {
            requesterID = &i
        }
    }
    if v := q.Get("category_id"); v != "" {
        if i, err := strconv.Atoi(v); err == nil {
            categoryID = &i
        }
    }

    tickets, total, err := controllers.ListTickets(claims.OrganizationID, page, limit, status, search, assigneeID, requesterID, categoryID)
    if err != nil {
        utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Error al listar tickets", err.Error()))
        return
    }

    utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
        "success": true,
        "message": "Tickets obtenidos",
        "data": map[string]interface{}{
            "items": tickets,
            "total": total,
            "page":  page,
            "limit": limit,
        },
    })
}

func UpdateTicketHandler(w http.ResponseWriter, r *http.Request) {
    claims, err := utils.GetUserFromContext(r)
    if err != nil {
        utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("No autorizado", err.Error()))
        return
    }

    idStr := mux.Vars(r)["ticket_id"]
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("ticket_id inválido", ""))
        return
    }

    var dto dtos.UpdateTicketDto
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Datos inválidos", err.Error()))
        return
    }
    validate := validator.New()
    if err := validate.Struct(dto); err != nil {
        utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Validación fallida", err.Error()))
        return
    }

    ticket, err := controllers.UpdateTicket(claims.OrganizationID, id, dto)
    if err != nil {
        utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Error al actualizar ticket", err.Error()))
        return
    }

    utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
        "success": true,
        "message": "Ticket actualizado",
        "data":    ticket,
    })
}

func AssignTicketHandler(w http.ResponseWriter, r *http.Request) {
    claims, err := utils.GetUserFromContext(r)
    if err != nil {
        utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("No autorizado", err.Error()))
        return
    }

    idStr := mux.Vars(r)["ticket_id"]
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("ticket_id inválido", ""))
        return
    }

    var dto dtos.AssignTicketDto
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Datos inválidos", err.Error()))
        return
    }
    validate := validator.New()
    if err := validate.Struct(dto); err != nil {
        utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Validación fallida", err.Error()))
        return
    }

    ticket, err := controllers.AssignTicket(claims.OrganizationID, id, claims.UserID, dto)
    if err != nil {
        utils.JSONResponse(w, http.StatusBadRequest, utils.ErrorResponse("Error al asignar ticket", err.Error()))
        return
    }

    utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
        "success": true,
        "message": "Ticket asignado",
        "data":    ticket,
    })
}
