package auth

type meResponse struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Suspended bool   `json:"suspended"`
	Image     string `json:"image"`
}
