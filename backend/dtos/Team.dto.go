package dtos

type CreateTeamDto struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=500"`
}

type UpdateTeamDto struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=3,max=100"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
}

type TeamResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	OrganizationID uint   `json:"organization_id"`
	MemberCount    int    `json:"member_count,omitempty"`
}

// ====================================
// TEAM MEMBER DTOs
// ====================================

type AddTeamMemberDto struct {
	UserID int    `json:"user_id" validate:"required,gt=0"`
	Role   string `json:"role" validate:"required,oneof=manager supervisor agent member viewer"`
}

type UpdateMemberRoleDto struct {
	Role string `json:"role" validate:"required,oneof=manager supervisor agent member viewer"`
}

type TeamMemberResponse struct {
	ID       uint         `json:"id"`
	TeamID   uint         `json:"team_id"`
	UserID   uint         `json:"user_id"`
	Role     string       `json:"role"`
	User     *UserSummary `json:"user,omitempty"`
	JoinedAt string       `json:"joined_at"`
}

type UserSummary struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

//Team Response

type TeamMembersByOrganizationResponse struct {
	ID      uint          `json:"id"`
	Name    string        `json:"name"`
	Members []MemberBasic `json:"members"`
}

type MemberBasic struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}
