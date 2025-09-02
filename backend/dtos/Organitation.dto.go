package dtos

type CreateOrganizationDTO struct {
	Name      string `json:"name" binding:"required"`
	Domain    string `json:"domain" binding:"required"`
	AdminUser struct {
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=6"`
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
	} `json:"admin_User" binding:"required"`
}
