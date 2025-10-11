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
	if dto.CategoryID != nil {
		var category models.Category
		if err := db.DB.Where("id = ? AND organization_id = ?", *dto.CategoryID, orgID).First(&category).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("categoria no pertenece a la organizacion")
			}
			return nil, err
		}
		cid := category.ID
		ticket.CategoryID = &cid
	}

	// En creacion forzamos ticket sin asignado. Asignacion se hace en endpoint dedicado.
	// Validar opcionalmente Team sugerido
	if dto.TeamID != nil {
		var team models.Team
		if err := db.DB.Where("id = ? AND organization_id = ?", *dto.TeamID, orgID).Preload("Category").First(&team).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("equipo no pertenece a la organizacion")
			}
			return nil, err
		}
		tid := team.ID
		ticket.TeamID = &tid

		if ticket.CategoryID != nil && team.CategoryID != nil && *ticket.CategoryID != *team.CategoryID {
			return nil, errors.New("la categoria seleccionada no coincide con la configurada para el equipo")
		}
		if ticket.CategoryID == nil && team.CategoryID != nil {
			cid := *team.CategoryID
			ticket.CategoryID = &cid
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
		CategoryID:   ticket.CategoryID,
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
		Select("tickets.id, tickets.ticket_number, tickets.subject, tickets.description, tickets.status, tickets.priority, tickets.requester_id, tickets.created_by_id, tickets.assignee_id, tickets.team_id, tickets.category_id").
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
	// Para asignacion usar endpoint dedicado
	if dto.CategoryID != nil {
		if *dto.CategoryID == 0 {
			ticket.CategoryID = nil
		} else {
			var category models.Category
			if err := db.DB.Where("id = ? AND organization_id = ?", *dto.CategoryID, orgID).First(&category).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errors.New("categoria no pertenece a la organizacion")
				}
				return nil, err
			}
			cid := category.ID
			ticket.CategoryID = &cid
		}

		if ticket.TeamID != nil && ticket.CategoryID != nil {
			var team models.Team
			if err := db.DB.Where("id = ?", *ticket.TeamID).Preload("Category").First(&team).Error; err == nil {
				if team.CategoryID != nil && *team.CategoryID != *ticket.CategoryID {
					return nil, errors.New("la categoria seleccionada no coincide con la del equipo asignado")
				}
			}
		}
	}
	if dto.DueDate != nil {
		ticket.DueDate = dto.DueDate
	}

	if err := db.DB.Save(&ticket).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

func DeleteTicket(orgID, ticketID int) error {
	// Iniciamos una transacción
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var ticket models.Ticket

		// Usamos tx en lugar de db.DB para todas las operaciones
		if err := tx.Where("id = ? AND organization_id = ?", ticketID, orgID).First(&ticket).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("ticket no encontrado")
			}
			return fmt.Errorf("error al buscar ticket: %v", err)
		}

		if ticket.Status != string(enums.StatusClosed) {
			return errors.New("solo se pueden eliminar tickets cerrados")
		}

		// Si hay tablas relacionadas, primero las eliminamos
		/* if err := tx.Where("ticket_id = ?", ticketID).Delete(&models.Comment{}).Error; err != nil {
			return fmt.Errorf("error al eliminar comentarios: %v", err)
		} */

		// Finalmente eliminamos el ticket
		if err := tx.Delete(&ticket).Error; err != nil {
			return fmt.Errorf("error al eliminar ticket: %v", err)
		}

		return nil
	})
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

	// 2) Validar equipo pertenece a la organizacion
	var team models.Team
	if err := db.DB.Where("id = ? AND organization_id = ?", dto.TeamID, orgID).Preload("Category").First(&team).Error; err != nil {
		return nil, errors.New("equipo no pertenece a la organizacion")
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

	if team.CategoryID != nil {
		if ticket.CategoryID != nil && *ticket.CategoryID != *team.CategoryID {
			return nil, errors.New("la categoria del ticket no coincide con la del equipo")
		}
		if ticket.CategoryID == nil {
			cid := *team.CategoryID
			ticket.CategoryID = &cid
		}
	}

	if err := db.DB.Save(&ticket).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}
