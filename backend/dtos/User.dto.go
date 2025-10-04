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

type CreateUserDto struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Role      string `json:"role" validate:"required"`

	OrganizationID int `json:"organization_id" validate:"required"`
}
