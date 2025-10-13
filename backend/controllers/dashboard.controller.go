package controllers

import (
	"github.com/marsDev10/helpdesk-backend/db"
	"github.com/marsDev10/helpdesk-backend/dtos"
	"github.com/marsDev10/helpdesk-backend/models"
)

func GetDashboardSummary(orgID int) (*dtos.DashboardSummaryDTO, error) {

	var summary dtos.DashboardSummaryDTO

	err := db.DB.Model(&models.Ticket{}).
		Where("organization_id = ?", orgID).
		Select("COUNT(*) as total_tickets",
			"SUM(CASE WHEN status = 'open' THEN 1 ELSE 0 END) as open_tickets",
			"SUM(CASE WHEN status = 'resolved' THEN 1 ELSE 0 END) as resolved_tickets",
			"SUM(CASE WHEN status = 'closed' THEN 1 ELSE 0 END) as closed_tickets").
		Scan(&summary).Error

	if err != nil {
		return nil, err
	}

	return &summary, nil

}
