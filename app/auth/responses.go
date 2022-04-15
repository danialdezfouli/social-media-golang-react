package auth

type loginResponse struct {
	AccessToken string `json:"accessToken"`
}

type meResponse struct {
	User struct {
		Name string
	}
}
