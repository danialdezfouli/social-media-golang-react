package model

import (
	"time"
)

type Post struct {
	PostId uint `gorm:"primarykey" faker:"-"`
	UserId uint `gorm:"index" faker:"-"`
	User   User `gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE;" faker:"-"`

	ParentId uint `gorm:"index" faker:"-"`

	Type    int    `gorm:"type:tinyint(1); default:1" faker:"oneof: 1,2,3,4"`
	Content string `gorm:"size:1000" faker:"sentence"`

	FavoritesCount int `gorm:"type:int(11); default:0"`
	RepliesCount   int `gorm:"type:int(11); default:0"`

	CreatedAt time.Time
}

func (post Post) GetType() string {
	switch post.Type {
	case 1:
		return "post"
	case 2:
		return "reply"
	case 3:
		return "quote"
	default:
		return "unknown-" + string(post.Type)
	}
}
