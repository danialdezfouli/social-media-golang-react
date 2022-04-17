package model

type Follow struct {
	FollowerId  uint `gorm:"primaryKey; autoIncrement:false;"`
	FollowingId uint `gorm:"primaryKey; autoIncrement:false;"`

	Follower  User `gorm:"foreignKey:FollowerId; constraint:OnDelete:CASCADE;"`
	Following User `gorm:"foreignKey:FollowingId; constraint:OnDelete:CASCADE;"`
}
