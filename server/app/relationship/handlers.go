package relationship

import (
	"github.com/labstack/echo/v4"
	"jupiter/app/common"
	"jupiter/app/relationship/dto"
	"jupiter/app/relationship/service"
	"net/http"
)

func follow(c echo.Context) error {
	params := new(dto.FollowDTO)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	f := service.FollowService{}
	user := common.GetUser(c)
	friend, err := f.FindUser(params.ID)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	err = f.Follow(user, friend)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	f.UpdateCounters(user)
	f.UpdateCounters(friend)

	return c.JSON(http.StatusOK, echo.Map{"message": "followed"})
}

func unfollow(c echo.Context) error {
	params := new(dto.FollowDTO)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	f := service.FollowService{}
	user := common.GetUser(c)
	friend, err := f.FindUser(params.ID)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	err = f.Unfollow(user, friend)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	f.UpdateCounters(user)
	f.UpdateCounters(friend)

	return c.JSON(http.StatusOK, echo.Map{"message": "unfollowed"})
}
