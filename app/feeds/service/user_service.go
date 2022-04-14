package service

import (
	"gorm.io/gorm"
	"jupiter/app/model"
	"log"
)

type userService struct {
	db   *gorm.DB
	post *model.Post
}

func (s userService) Follow(user2 *model.User, user *model.User) userService {
	follow := &model.Follow{
		Follower:  *user2,
		Following: *user,
	}

	s.db.Create(follow)

	return s
}

func (s userService) UnFollow(user2 *model.User, user *model.User) userService {
	log.Fatal("unfollow is not implemented")

	return s
}

func NewUserService(db *gorm.DB) *userService {
	return &userService{db: db}
}
