package dtos

type AddConversationDto struct {
	Message  string `json:"message" validate:"required,min=1"`
	IsPublic *bool  `json:"is_public,omitempty"`
}
