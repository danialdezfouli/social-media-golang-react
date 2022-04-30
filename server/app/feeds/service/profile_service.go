package service

import (
	"github.com/labstack/echo/v4"
	"jupiter/app"
	"jupiter/app/common"
	"jupiter/app/feeds/dto"
	"jupiter/app/feeds/repository"
	"jupiter/app/model"
	"net/http"
)

func FindProfile(c echo.Context) (*repository.Profile, error) {
	params := new(dto.ProfileDTO)
	user := common.GetUser(c)

	if err := c.Bind(params); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var profile *repository.Profile
	result := app.GetDB().Model(&model.User{}).Where("username", params.Username).First(&profile)

	if result.RowsAffected == 0 {
		return nil, echo.ErrNotFound
	}

	if user.ID != profile.ID {
		profile.Followed = IsFollowing(user.ID, profile.ID)
	}

	return profile, nil

}

func IsFollowing(followerId, followingId uint) bool {
	result := app.GetDB().Where(model.Follow{
		FollowerId:  followerId,
		FollowingId: followingId,
	}).First(&model.Follow{})

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
