package service

import (
	"gorm.io/gorm"
	"jupiter/app/model"
)

type postService struct {
	db   *gorm.DB
	post *model.Post
}

func (s postService) CreatePost(p *model.Post) postService {
	s.post = p
	s.db.Create(p)

	return s
}

func (s postService) AddHashtag(name string) postService {
	hashtag := &model.Hashtag{
		Name: name,
	}
	s.db.Where(hashtag).FirstOrCreate(&hashtag)

	hashtagPost := &model.HashtagPost{
		Post:    *s.post,
		Hashtag: *hashtag,
	}
	s.db.Create(hashtagPost)

	// update posts count
	var count int64
	s.db.Model(model.HashtagPost{}).Where("hashtag_id", hashtag.ID).Count(&count)
	s.db.Model(&hashtag).Update("posts_count", count)

	return s
}

func NewPostService(db *gorm.DB) *postService {
	return &postService{
		db: db,
	}
}
