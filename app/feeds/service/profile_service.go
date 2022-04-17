package service

import (
	"github.com/labstack/echo/v4"
	"jupiter/app"
	"jupiter/app/feeds/dto"
	"jupiter/app/feeds/repository"
	"jupiter/app/model"
	"net/http"
)

func FindProfile(c echo.Context) (*repository.Profile, error) {
	params := new(dto.ProfileDTO)
	if err := c.Bind(params); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var profile *repository.Profile
	result := app.GetDB().Model(&model.User{}).Where("id", params.ID).First(&profile)
	if result.RowsAffected == 0 {
		return nil, echo.ErrNotFound
	}

	return profile, nil

}
