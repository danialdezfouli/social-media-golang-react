package repository

type SearchProfile struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	Image          string `json:"image"`
	Official       bool   `json:"official,omitempty"`
	FollowersCount int    `json:"followers_count"`
}
