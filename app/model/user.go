package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID uint `gorm:"primarykey" faker:"-"`

	Name           string `gorm:"size:255" faker:"name"`
	Username       string `json:"username" gorm:"size:255; index; not null" faker:"username"`
	Email          string `json:"email" gorm:"size:255; index; not null" faker:"email"`
	Password       string `gorm:"size:255; not null" faker:"password"`
	Bio            string `gorm:"size:255" faker:"sentence"`
	Image          string `gorm:"size:255" faker:"-"`
	Birthday       string `gorm:"type: date; default: null" faker:"-"`
	Suspended      bool   `gorm:"size:1; default:0" faker:"-"`
	Official       bool   `gorm:"size:1; default:0" faker:"-"`
	FollowersCount int
	FollowingCount int

	EmailVerifiedAt *time.Time     `faker:"-"`
	CreatedAt       time.Time      `faker:"-"`
	UpdatedAt       time.Time      `faker:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index" faker:"-"`
}
