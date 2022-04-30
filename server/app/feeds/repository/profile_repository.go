package repository

import (
	"time"
)

type Profile struct {
	ID             uint      `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Username       string    `json:"username,omitempty"`
	Bio            string    `json:"bio,omitempty"`
	Image          string    `json:"image,omitempty"`
	Birthday       string    `json:"birthday,omitempty"`
	Suspended      bool      `json:"suspended,omitempty"`
	Official       bool      `json:"official,omitempty"`
	Followed       bool      `json:"followed" gorm:"-"`
	FollowersCount int       `json:"followers_count,omitempty"`
	FollowingCount int       `json:"following_count,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}
