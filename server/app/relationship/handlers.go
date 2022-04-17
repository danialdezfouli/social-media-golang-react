package relationship

import (
	"github.com/labstack/echo/v4"
	"jupiter/app/model"
	"jupiter/app/relationship/dto"
	"jupiter/app/relationship/service"
	"net/http"
)

func follow(c echo.Context) error {
	params := new(dto.FollowDTO)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*model.User)
	f := service.FollowService{}
	err := f.Follow(user, params.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "followed"})
}

func unfollow(c echo.Context) error {
	params := new(dto.FollowDTO)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*model.User)
	f := service.FollowService{}
	err := f.Unfollow(user, params.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "unfollowed"})
}
