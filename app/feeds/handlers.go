package feeds

import (
	"github.com/labstack/echo/v4"
	"jupiter/app"
	"jupiter/app/feeds/repository"
	"jupiter/app/model"
	"net/http"
)

func timeline(c echo.Context) error {
	user := c.Get("user").(*model.User)
	var posts = &[]repository.Post{}

	db := app.GetDB()

	db.Model(&model.Post{}).
		Select("posts.*, users.name profile_name, users.username profile_username, users.image as profile_image").
		Joins("right join follows on follows.follower_id = ?", user.ID).
		Joins("right join users on users.id = user_id").
		Group("post_id").
		Where("posts.user_id = follows.following_id").
		Order("created_at desc").
		Find(posts)

	//for i, post := range *posts {
	//	post.TypeText = translate(post.Type)
	//}

	return c.JSON(http.StatusOK, posts)
}
