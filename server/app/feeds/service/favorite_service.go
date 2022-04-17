package service

import (
	"gorm.io/gorm"
	"jupiter/app/model"
)

type favoriteService struct {
	db   *gorm.DB
	post *model.Post
}

func (s favoriteService) AddFavorite(post *model.Post, user *model.User) favoriteService {
	s.db.Create(&model.Favorite{
		UserId: user.ID,
		PostId: post.PostId,
	})

	return s

}

func NewFavoriteService(db *gorm.DB) *favoriteService {
	return &favoriteService{db: db}
}
