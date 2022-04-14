package auth

type loginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type meResponse struct {
	User struct {
		Name string
	}
}
