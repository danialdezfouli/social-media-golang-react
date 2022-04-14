package service

import (
	"errors"
	"gorm.io/gorm"
	"jupiter/app"
	"jupiter/app/auth/dto"
	"jupiter/app/common/bcrypt"
	"jupiter/app/model"
)

type LoginService struct {
}

func (s LoginService) Attempt(input *dto.LoginInput) (*model.User, error) {
	db := app.GetInstance().DB

	var user *model.User

	result := db.Model(model.User{}).Where("email=? or username=?", input.Username, input.Username).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || user == nil {
		return nil, errors.New("user not found")
	}

	if !bcrypt.Compare(user.Password, input.Password) {
		return nil, errors.New("password is invalid")
	}

	if user.Suspended {
		return nil, errors.New("your account is suspended")
	}

	return user, nil

}

func (s LoginService) LoginUsingId(id int) (*model.User, error) {
	return nil, nil
}
