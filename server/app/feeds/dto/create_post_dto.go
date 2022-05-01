package dto

type CreatePostDTO struct {
	Body      string `json:"body" validate:"required,max=300"`
	Type      string `json:"type" validate:"oneof=post reply"`
	ReplyToID uint   `json:"reply_to,omitempty"`
}
