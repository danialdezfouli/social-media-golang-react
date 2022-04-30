package service

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"jupiter/app/common"
	"jupiter/app/feeds/repository"
	"jupiter/app/model"
	"regexp"
	"strings"
)

type postService struct {
	db   *gorm.DB
	post *model.Post
	c    echo.Context
}

func NewPostService(db *gorm.DB, context echo.Context) *postService {
	return &postService{
		db: db,
		c:  context,
	}
}

func (s postService) FindPost(id uint) (*repository.Post, error) {
	var post *repository.Post
	user := common.GetUser(s.c)

	result := QueryTimelineBasic(user).Where(model.Post{PostId: id}).First(&post)

	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}

func (s postService) ToggleLike(post *repository.Post, user *model.User) bool {
	fav := model.Favorite{}
	result := s.db.Where(&model.Favorite{UserId: user.ID, PostId: post.PostId}).First(&fav)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fav := &model.Favorite{
			UserId: user.ID,
			PostId: post.PostId,
		}
		s.db.Create(fav)

		return true
	}

	s.db.Where(&model.Favorite{UserId: user.ID, PostId: post.PostId}).Limit(1).Delete(model.Favorite{})

	return false
}

func (s postService) FindReplies(post *repository.Post) []repository.Post {
	var replies []repository.Post
	user := common.GetUser(s.c)

	result := QueryTimelineBasic(user).Where("posts.parent_id", post.PostId).Where("posts.post_type", model.PostTypeReply).Find(&replies)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return []repository.Post{}
	}

	return replies
}

func getParent(user *model.User, post *repository.Post) *repository.Post {
	var parent *repository.Post
	result := QueryTimelineBasic(user).Where("posts.post_id", post.ParentId).First(&parent)

	if result.Error != nil {
		return nil
	}

	return parent
}

func (s postService) FindPostParentsHirarchy(post *repository.Post) []repository.Post {
	var parents []repository.Post
	user := common.GetUser(s.c)

	n := 0
	parent := post

	for {
		n++
		if n > 10 {
			break
		}

		fmt.Println("post.ParentId", post.ParentId)
		if post.ParentId == 0 {
			break
		}
		parent = getParent(user, parent)
		fmt.Println(parent)

		if parent == nil {
			break
		}
		parents = append(parents, *parent)
	}

	return parents
}

func (s postService) CreatePost(p *model.Post) postService {
	s.post = p
	s.db.Create(p)

	tags := FindHashtags(p.Content)
	for _, tag := range tags {
		s.AddHashtag(p, tag)
	}

	return s
}

func (s postService) AddHashtag(p *model.Post, name string) postService {
	hashtag := &model.Hashtag{
		Name: name,
	}
	s.db.Where(hashtag).FirstOrCreate(&hashtag)

	hashtagPost := &model.HashtagPost{
		PostId:    p.PostId,
		HashtagId: hashtag.ID,
	}
	s.db.FirstOrCreate(hashtagPost)

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

	result := s.db.Model(&post).Limit(1).Updates(map[string]interface{}{
		"replies_count":   replies,
		"favorites_count": favorites,
	})

	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}
}

func (s postService) FindParentsForTimeline(posts *[]repository.Post) *[]repository.Post {
	var parents = &[]repository.Post{}
	var parentIds []uint

	user := common.GetUser(s.c)
	for _, post := range *posts {
		if post.ParentId != 0 && !common.Contains(parentIds, post.ParentId) {
			parentIds = append(parentIds, post.ParentId)
		}
	}

	QueryTimelineBasic(user).Where("posts.post_id in ?", parentIds).Find(parents)
	return parents
}

func (s postService) KeyByPostId(posts *[]repository.Post) map[uint]repository.Post {
	var postsMap = map[uint]repository.Post{}
	for _, post := range *posts {
		postsMap[post.PostId] = post
	}
	return postsMap
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
