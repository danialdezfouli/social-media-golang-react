package service

import (
	"gorm.io/gorm"
	"jupiter/app"
	"jupiter/app/model"
)

const timelinePostLimit = 30
const profileSearchPostLimit = 10

func QueryTimelineBasic(user *model.User) *gorm.DB {
	return app.GetDB().Model(&model.Post{}).
		Select("posts.*, users.name profile_name, users.username profile_username, users.image as profile_image, "+
			"case when user_posts.post_type='repost' then TRUE ELSE FALSE END as reposted, "+
			"(CASE WHEN EXISTS ( SELECT * FROM favorites WHERE favorites.post_id = posts.post_id and favorites.user_id = ? ) THEN TRUE ELSE FALSE END) as liked", user.ID,
		).
		Joins("inner join users on users.id = user_id").
		Joins("left join posts user_posts on user_posts.parent_id = posts.post_id and user_posts.post_type = ? and user_posts.user_id = ?", "repost", user.ID)
	//Joins("left join favorites on favorites.post_id = posts.post_id and favorites.user_id = ?", user.Username)

}
func QueryTimeline(offset int, user *model.User) *gorm.DB {
	return QueryTimelineBasic(user).
		//Group("posts.post_id").
		Order("created_at desc").
		Limit(timelinePostLimit).
		Offset(offset)
}

func QuerySuggestedProfiles(user *model.User) *gorm.DB {
	return app.GetDB().Model(&model.User{}).
		Where("users.id not in (select following_id from follows where follower_id=?)", user.ID).
		Where("users.id != ?", user.ID).
		Order("rand()").
		Limit(3)
}

func SearchUsers(query string) *gorm.DB {
	return app.GetDB().Model(&model.User{}).
		Where("name like ?", "%"+query+"%").
		Or("username like ?", query+"%").
		Order("followers_count desc, created_at desc").
		Limit(profileSearchPostLimit)

}
