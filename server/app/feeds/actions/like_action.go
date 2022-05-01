package actions

import (
	"github.com/labstack/echo/v4"
	"jupiter/app"
	"jupiter/app/common"
	"jupiter/app/feeds/dto"
	"jupiter/app/feeds/service"
	"jupiter/app/model"
	"net/http"
)

func LikePost(c echo.Context) error {
	params := new(dto.FindPostDTO)
	err := c.Bind(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := common.GetUser(c)

	db := app.GetDB()
	postService := service.NewPostService(db, c)
	post, findPostErr := postService.FindPost(params.ID)
	if findPostErr != nil {
		return echo.ErrNotFound
	}

	liked := postService.ToggleLike(post, user)
	postService.UpdatePostCounters(&model.Post{PostId: post.PostId})

	return c.JSON(http.StatusOK, echo.Map{
		"liked": liked,
	})

}
