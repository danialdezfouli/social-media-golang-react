package feeds

import (
	"github.com/labstack/echo/v4"
	"jupiter/app"
	"jupiter/app/feeds/repository"
	"jupiter/app/model"
	"net/http"
)

const (
	timelinePostLimit = 30
)

func timeline(c echo.Context) error {
	params := new(timelineDTO)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

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
		Limit(timelinePostLimit).
		Offset(int(params.Offset)).
		Find(posts)

	return c.JSON(http.StatusOK, posts)
}

func profile(c echo.Context) error {
	params := new(profileDTO)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user *model.User

	result := app.GetDB().Where(model.User{ID: params.ID}).First(&user)

	if result.RowsAffected == 0 {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, repository.ProfileResponse{
		ID:             user.ID,
		Name:           user.Name,
		Username:       user.Username,
		Bio:            user.Bio,
		Image:          user.Image,
		Birthday:       user.Birthday,
		Suspended:      user.Suspended,
		Official:       user.Official,
		FollowersCount: user.FollowersCount,
		FollowingCount: user.FollowingCount,
		CreatedAt:      user.CreatedAt,
	})

}
