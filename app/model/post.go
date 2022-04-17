package model

import (
	"database/sql"
	"time"
)

type Post struct {
	PostId uint `gorm:"primarykey" faker:"-"`
	UserId uint `gorm:"index" faker:"-"`
	User   User `gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE;" faker:"-"`

	ParentId       sql.NullInt32 `gorm:"index; default: null" faker:"-"`
	PostType       string        `gorm:"index; size:10; default:post" faker:"oneof: post,repost,reply,quote"`
	Content        string        `gorm:"size:1000" faker:"sentence"`
	FavoritesCount uint          `gorm:"type:int(11); default:0" faker:"-"`
	RepliesCount   uint          `gorm:"type:int(11); default:0" faker:"-"`

	CreatedAt time.Time
}

const (
	PostTypePost   = "post"
	PostTypeReply  = "reply"
	PostTypeRepost = "repost"
	PostTypeQuote  = "quote"
)
