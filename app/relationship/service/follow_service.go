package service

import (
	"errors"
	"gorm.io/gorm"
	"jupiter/app"
	"jupiter/app/model"
)

type FollowService struct {
}

func findUser(user *model.User, id uint) (*model.User, error) {
	if user.ID == id {
		return nil, errors.New("you cannot follow yourself")
	}

	followUser := &model.User{}
	result := app.GetDB().Where(&model.User{ID: id}).First(&followUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	return followUser, nil
}

func (s FollowService) Follow(user *model.User, id uint) error {
	friend, err := findUser(user, id)

	if err != nil {
		return err
	}

	app.GetDB().FirstOrCreate(&model.Follow{
		FollowerId:  user.ID,
		FollowingId: friend.ID,
	})

	return nil
}

func (s FollowService) Unfollow(user *model.User, id uint) error {
	friend, err := findUser(user, id)

	if err != nil {
		return err
	}

	app.GetDB().Where(&model.Follow{
		FollowerId:  user.ID,
		FollowingId: friend.ID,
	}).Delete(&model.Follow{})

	return nil
}
