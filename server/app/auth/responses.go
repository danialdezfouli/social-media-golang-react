package auth

type meResponse struct {
	ID        uint   `json:"id"`
	Bio       string `json:"bio"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Suspended bool   `json:"suspended"`
	Image     string `json:"image"`
}
