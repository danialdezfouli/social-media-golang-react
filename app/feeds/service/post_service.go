package service

import (
	"gorm.io/gorm"
	"jupiter/app/model"
	"regexp"
	"strings"
)

type postService struct {
	db   *gorm.DB
	post *model.Post
}

func (s postService) CreatePost(p *model.Post) postService {
	s.post = p
	s.db.Create(p)

	// TODO: find hashtags
	// s.AddHashtag("#tag")

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

func (s postService) UpdatePostCounters(post *model.Post) {
	var replies int64
	var favorites int64

	s.db.Model(&model.Post{}).Where("post_type", "reply").Where("parent_id", post.PostId).Count(&replies)
	s.db.Model(&model.Favorite{}).Where("post_id", post.PostId).Count(&favorites)

	post.RepliesCount = uint(replies)
	post.FavoritesCount = uint(favorites)
	s.db.Save(&post)
}

func NewPostService(db *gorm.DB) *postService {
	return &postService{
		db: db,
	}
}

func FindHashtags(content string) []string {
	pattern := "#[^\\s-]+"
	r := regexp.MustCompile(pattern)
	tags := r.FindAllString(content, -1)

	// Remove #
	for i, tag := range tags {
		tags[i] = strings.TrimLeft(tag, "#")
	}

	return tags
}
