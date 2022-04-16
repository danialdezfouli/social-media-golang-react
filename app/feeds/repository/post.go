package repository

import (
	"time"
)

type Post struct {
	PostId          uint      `json:"post_id"`
	UserId          uint      `json:"user_id"`
	ProfileName     string    `json:"profile_name"`
	ProfileImage    string    `json:"profile_image"`
	ProfileUsername string    `json:"profile_username"`
	ParentId        uint      `json:"parent_id"`
	PostType        string    `json:"post_type"`
	Content         string    `json:"content"`
	FavoritesCount  int       `json:"favorites_count"`
	RepliesCount    int       `json:"replies_count"`
	CreatedAt       time.Time `json:"created_at"`
}
