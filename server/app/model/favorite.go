package model

import "time"

type Favorite struct {
	UserId uint `gorm:"primaryKey; autoIncrement:false;"`
	PostId uint `gorm:"primaryKey; autoIncrement:false;"`

	User User `gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE;"`
	Post Post `gorm:"foreignKey:PostId; constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
}
