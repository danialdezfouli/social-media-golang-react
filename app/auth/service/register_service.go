package service

import (
	"github.com/labstack/echo/v4"
	"jupiter/app"
	"jupiter/app/auth/dto"
	"jupiter/app/model"
	"net/http"
	"strings"
)

type RegisterService struct {
}

func (s RegisterService) Register(c echo.Context, input *dto.RegisterInput) error {
	db := app.GetInstance().DB
	var foundUser *model.User
	result := db.Model(&model.User{}).
		Where("LOWER(email)=? or LOWER(username)=?", strings.ToLower(input.Email), strings.ToLower(input.Username)).
		First(&foundUser)

	if result.RowsAffected > 0 {
		if strings.ToLower(input.Email) == strings.ToLower(foundUser.Email) {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "user already exists with this email")
		} else {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "choose a unique username")
		}
	}

	createdUser := &model.User{
		Email:    input.Email,
		Name:     input.Name,
		Username: input.Username,
		Password: input.Password,
	}

	creatingResult := db.Create(createdUser)

	if creatingResult.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Please try later")
	}

	// TODO: return access token

	return c.JSON(http.StatusOK, "registered successfully")
}
