package service

import (
	"gorm.io/gorm"
	"jupiter/app"
	"jupiter/app/model"
)

const timelinePostLimit = 30

func QueryTimeline(offset int) *gorm.DB {
	return app.GetDB().Model(&model.Post{}).
		Select("posts.*, users.name profile_name, users.username profile_username, users.image as profile_image").
		Joins("inner join users on users.id = user_id").
		Group("post_id").
		Order("created_at desc").
		Limit(timelinePostLimit).
		Offset(offset)

}
