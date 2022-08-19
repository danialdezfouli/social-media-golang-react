package dto

type UpdateProfileInput struct {
	Name string `json:"name" validate:"required"`
	Bio  string `json:"bio"`
}
