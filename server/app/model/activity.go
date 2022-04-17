package model

type Activity struct {
	Id     uint   `gorm:"primarykey"`
	UserId uint   `gorm:"index"`
	PostId uint   `gorm:"index"`
	Type   string `gorm:"size:40"`

	User User `gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE;"`
	Post User `gorm:"foreignKey:PostId; constraint:OnDelete:CASCADE;"`
}
