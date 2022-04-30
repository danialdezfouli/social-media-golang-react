package repository

import (
	"time"
)

type Profile struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Bio            string    `json:"bio"`
	Image          string    `json:"image,omitempty"`
	Birthday       string    `json:"birthday,omitempty"`
	Suspended      bool      `json:"suspended,omitempty"`
	Official       bool      `json:"official,omitempty"`
	Followed       bool      `json:"followed" gorm:"-"`
	FollowersCount int       `json:"followers_count"`
	FollowingCount int       `json:"following_count"`
	CreatedAt      time.Time `json:"created_at"`
}
