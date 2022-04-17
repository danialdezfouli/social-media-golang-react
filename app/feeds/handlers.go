package feeds

import (
	"github.com/labstack/echo/v4"
	"jupiter/app/feeds/dto"
	"jupiter/app/feeds/repository"
	"jupiter/app/feeds/service"
	"jupiter/app/model"
	"net/http"
)

func homeTimeline(c echo.Context) error {
	params := new(dto.TimelineDTO)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*model.User)
	var posts = &[]repository.Post{}

	service.QueryTimeline(int(params.Offset)).
		Joins("inner join follows on follows.follower_id = ?", user.ID).
		Where("posts.user_id = follows.following_id").
		Find(posts)

	return c.JSON(http.StatusOK, posts)
}

func profile(c echo.Context) error {
	profile, err := service.FindProfile(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, profile)
}

func profileTimeline(c echo.Context) error {
	params := new(dto.TimelineDTO)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	profile, err := service.FindProfile(c)
	if err != nil {
		return err
	}

	var posts = &[]repository.Post{}

	service.QueryTimeline(int(params.Offset)).
		Where("posts.user_id = ?", profile.ID).
		Find(posts)

	return c.JSON(http.StatusOK, posts)
}

func profileLikes(c echo.Context) error {
	params := new(dto.TimelineDTO)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	profile, err := service.FindProfile(c)
	if err != nil {
		return err
	}

	var posts = &[]repository.Post{}

	service.QueryTimeline(int(params.Offset)).
		Joins("inner join favorites on favorites.post_id = posts.post_id").
		Where("favorites.user_id", profile.ID).
		Find(posts)

	return c.JSON(http.StatusOK, posts)
}
