package feeds

import (
	"github.com/labstack/echo/v4"
	"jupiter/app"
	"jupiter/app/common"
	"jupiter/app/feeds/dto"
	"jupiter/app/feeds/repository"
	"jupiter/app/feeds/service"
	"net/http"
)

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

	postService := service.NewPostService(app.GetDB(), c)

	user := common.GetUser(c)
	profile, err := service.FindProfile(c)
	if err != nil {
		return err
	}

	var posts = &[]repository.Post{}

	service.QueryTimeline(int(params.Offset), user).
		Where("posts.user_id = ?", profile.ID).
		Find(posts)

	parents := postService.FindParentsForTimeline(posts)
	parentsMap := postService.KeyByPostId(parents)

	return c.JSON(http.StatusOK, echo.Map{
		"posts":   posts,
		"parents": parentsMap,
	})
}

func profileLikes(c echo.Context) error {
	params := new(dto.TimelineDTO)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := common.GetUser(c)
	profile, err := service.FindProfile(c)
	if err != nil {
		return err
	}

	var posts = &[]repository.Post{}

	service.QueryTimeline(int(params.Offset), user).
		Joins("inner join favorites on favorites.post_id = posts.post_id").
		Where("favorites.user_id", profile.ID).
		Find(posts)

	return c.JSON(http.StatusOK, posts)
}
