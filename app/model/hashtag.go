package model

type Hashtag struct {
	ID         uint   `gorm:"primarykey"`
	Name       string `gorm:"size:50; index;"`
	PostsCount int    `gorm:"default:0;"`
}

type HashtagPost struct {
	HashtagId uint `gorm:"primaryKey; autoIncrement:false;"`
	PostId    uint `gorm:"primaryKey; autoIncrement:false;"`

	Hashtag Hashtag `gorm:"foreignKey:HashtagId; constraint:OnDelete:CASCADE;"`
	Post    Post    `gorm:"foreignKey:PostId; constraint:OnDelete:CASCADE;"`
}
