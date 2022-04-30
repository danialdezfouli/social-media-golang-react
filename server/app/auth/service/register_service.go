package service

import (
	"github.com/labstack/echo/v4"
	"jupiter/app"
	"jupiter/app/auth/dto"
	"jupiter/app/common/bcrypt"
	"jupiter/app/model"
	"net/http"
	"strings"
)

type RegisterService struct {
}

func (s RegisterService) Register(input *dto.RegisterInput) (*model.User, error) {
	db := app.GetDB()
	var foundUser *model.User
	result := db.Model(&model.User{}).
		Where("LOWER(email)=? or LOWER(username)=?", strings.ToLower(input.Email), strings.ToLower(input.Username)).
		First(&foundUser)

	if result.RowsAffected > 0 {
		if strings.ToLower(input.Email) == strings.ToLower(foundUser.Email) {
			return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "email-exists")
		} else {
			return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "username-exists")
		}

	}

	password, err := bcrypt.Hash(input.Password)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "something was wrong with the password")
	}

	createdUser := &model.User{
		Email:    input.Email,
		Name:     input.Name,
		Username: input.Username,
		Password: password,
	}

	creatingResult := db.Create(createdUser)

	if creatingResult.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Please try later")
	}

	return createdUser, nil
}
