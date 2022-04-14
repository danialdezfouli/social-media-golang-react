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
		User: *user,
		Post: *post,
	})

	return s

}

func NewFavoriteService(db *gorm.DB) *favoriteService {
	return &favoriteService{db: db}
}
