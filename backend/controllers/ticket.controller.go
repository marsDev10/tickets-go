package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/marsDev10/helpdesk-backend/db"
	"github.com/marsDev10/helpdesk-backend/dtos"
	"github.com/marsDev10/helpdesk-backend/enums"
	"github.com/marsDev10/helpdesk-backend/models"
	"gorm.io/gorm"
)

func generateTicketNumber(orgID int) string {
	// Simple unique-ish number: ORG-<org>-<unix-nano>
	return fmt.Sprintf("ORG-%d-%d", orgID, time.Now().UnixNano())
}

func CreateTicket(orgID int, createdByID uint, dto dtos.CreateTicketDto) (*dtos.TicketResponseDto, error) {
	// Determinar requester (por defecto el creador)
	requesterID := createdByID
	if dto.RequesterID != nil {
		var requester models.User
		if err := db.DB.Where("id = ? AND organization_id = ?", *dto.RequesterID, orgID).First(&requester).Error; err != nil {
			return nil, errors.New("requester no pertenece a la organización")
		}
		requesterID = uint(*dto.RequesterID)
	}

	ticket := models.Ticket{
		TicketNumber:   generateTicketNumber(orgID),
		Subject:        dto.Subject,
		Description:    dto.Description,
		Status:         string(enums.StatusOpen),
		Priority:       string(enums.PriorityMedium),
		RequesterID:    requesterID,
		CreatedByID:    createdByID,
		OrganizationID: uint(orgID),
	}

	if dto.Priority != nil {
		ticket.Priority = string(*dto.Priority)
	}
	/* if dto.CategoryID != nil {
	    ticket.CategoryID = uint(*dto.CategoryID)
	} */

	// En creación forzamos ticket sin asignado. Asignación se hace en endpoint dedicado.
	// Validar opcionalmente Team sugerido
	if dto.TeamID != nil {
		var team models.Team
		if err := db.DB.Where("id = ? AND organization_id = ?", *dto.TeamID, orgID).First(&team).Error; err == nil {
			tid := uint(*dto.TeamID)
			ticket.TeamID = &tid
		}
	}
	if dto.DueDate != nil {
		ticket.DueDate = dto.DueDate
	}

	// Crear el ticket
	if err := db.DB.Create(&ticket).Error; err != nil {
		return nil, err
	}

	// Convertir a DTO de respuesta
	response := dtos.TicketResponseDto{
		ID:           ticket.ID,
		TicketNumber: ticket.TicketNumber,
		Subject:      ticket.Subject,
		Description:  ticket.Description,
		Status:       ticket.Status,
		Priority:     ticket.Priority,
		RequesterID:  ticket.RequesterID,
		CreatedByID:  ticket.CreatedByID,
		AssigneeID:   ticket.AssigneeID,
		TeamID:       ticket.TeamID,
	}

	return &response, nil
}

func GetTicketByID(orgID, ticketID int) (*models.Ticket, error) {
	var ticket models.Ticket
	err := db.DB.Where("id = ? AND organization_id = ?", ticketID, orgID).
		Preload("Requester").Preload("Assignee").Preload("Category").
		First(&ticket).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("ticket no encontrado")
		}
		return nil, err
	}
	return &ticket, nil
}

func ListTickets(orgID, page, limit int, status, search string, assigneeID, requesterID, categoryID *int) ([]dtos.TicketResponseDto, int64, error) {
	query := db.DB.Model(&models.Ticket{}).Where("organization_id = ?", orgID)

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if search != "" {
		like := "%" + search + "%"
		query = query.Where("LOWER(subject) LIKE LOWER(?)", like)
	}
	if assigneeID != nil {
		query = query.Where("assignee_id = ?", *assigneeID)
	}
	if requesterID != nil {
		query = query.Where("requester_id = ?", *requesterID)
	}
	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	offset := (page - 1) * limit

	var tickets []dtos.TicketResponseDto
	if err := query.
		Select("tickets.id, tickets.ticket_number, tickets.subject, tickets.description, tickets.status, tickets.priority, tickets.requester_id, tickets.created_by_id, tickets.assignee_id, tickets.team_id").
		Order("tickets.created_at DESC").Offset(offset).Limit(limit).Scan(&tickets).Error; err != nil {
		return nil, 0, err
	}

	return tickets, total, nil
}

func UpdateTicket(orgID, ticketID int, dto dtos.UpdateTicketDto) (*models.Ticket, error) {
	var ticket models.Ticket
	if err := db.DB.Where("id = ? AND organization_id = ?", ticketID, orgID).First(&ticket).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("ticket no encontrado")
		}
		return nil, err
	}

	if dto.Subject != nil {
		ticket.Subject = *dto.Subject
	}
	if dto.Description != nil {
		ticket.Description = *dto.Description
	}
	if dto.Priority != nil {
		ticket.Priority = string(*dto.Priority)
	}
	if dto.Status != nil {
		prev := ticket.Status
		ticket.Status = string(*dto.Status)
		if (prev != string(enums.StatusResolved) && *dto.Status == enums.StatusResolved) || (prev != string(enums.StatusClosed) && *dto.Status == enums.StatusClosed) {
			now := time.Now()
			ticket.ResolvedAt = &now
		}
	}
	// Para asignación usar endpoint dedicado
	/* if dto.CategoryID != nil {
		ticket.CategoryID = uint(*dto.CategoryID)
	} */
	if dto.DueDate != nil {
		ticket.DueDate = dto.DueDate
	}

	if err := db.DB.Save(&ticket).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

// AssignTicket asigna un ticket a un agente dentro de un equipo, validando que el asignador pueda gestionar el equipo
func AssignTicket(orgID, ticketID int, assignerID uint, dto dtos.AssignTicketDto) (*models.Ticket, error) {
	// 1) Traer ticket de la organización
	var ticket models.Ticket
	if err := db.DB.Where("id = ? AND organization_id = ?", ticketID, orgID).First(&ticket).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("ticket no encontrado")
		}
		return nil, err
	}

	// 2) Validar equipo pertenece a la organización
	var team models.Team
	if err := db.DB.Where("id = ? AND organization_id = ?", dto.TeamID, orgID).First(&team).Error; err != nil {
		return nil, errors.New("equipo no pertenece a la organización")
	}

	// 3) Validar que asignador sea manager/supervisor del equipo
	var assignerMembership models.TeamMember
	if err := db.DB.Where("team_id = ? AND user_id = ?", dto.TeamID, assignerID).First(&assignerMembership).Error; err != nil {
		return nil, errors.New("no perteneces a este equipo")
	}
	if !assignerMembership.CanManageTickets() {
		return nil, errors.New("no tienes permisos para asignar en este equipo")
	}

	// 4) Validar que el asignado pertenece a la organización y al equipo
	var assignee models.User
	if err := db.DB.Where("id = ? AND organization_id = ?", dto.AssigneeID, orgID).First(&assignee).Error; err != nil {
		return nil, errors.New("asignado no pertenece a la organización")
	}
	var assigneeMembership models.TeamMember
	if err := db.DB.Where("team_id = ? AND user_id = ?", dto.TeamID, dto.AssigneeID).First(&assigneeMembership).Error; err != nil {
		return nil, errors.New("asignado no pertenece a este equipo")
	}

	// 5) Aplicar asignación
	tid := uint(dto.TeamID)
	aid := uint(dto.AssigneeID)
	ticket.TeamID = &tid
	ticket.AssigneeID = &aid

	if err := db.DB.Save(&ticket).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}
