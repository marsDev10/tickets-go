package dtos

type GetUsersDto struct {
	id string `json:"id" validate:"required"`
}

type GetUsersData struct {
	first_name string
	last_name  string
}

type UserResponse struct {
	ID             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	OrganizationID int    `json:"organization_id"`
}
