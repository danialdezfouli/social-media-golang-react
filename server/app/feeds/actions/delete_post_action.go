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

func DeletePost(c echo.Context) error {
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

	if user.ID != post.UserId {
		return echo.ErrUnauthorized
	}

	db.Where(model.Post{
		PostId: post.PostId,
	}).Limit(1).Delete(&model.Post{})

	return c.JSON(http.StatusOK, "deleted")
}
