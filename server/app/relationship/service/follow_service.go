package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"jupiter/app"
	"jupiter/app/model"
)

type FollowService struct {
}

func NewFollowService() *FollowService {
	return &FollowService{}
}

func (s FollowService) FindUser(id uint) (*model.User, error) {

	followUser := &model.User{}
	result := app.GetDB().Where(&model.User{ID: id}).First(&followUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	return followUser, nil
}

func (s FollowService) Follow(user, friend *model.User) error {
	if user.ID == friend.ID {
		return errors.New("you cannot follow yourself")
	}

	app.GetDB().FirstOrCreate(&model.Follow{
		FollowerId:  user.ID,
		FollowingId: friend.ID,
	})

	return nil
}

func (s FollowService) Unfollow(user, friend *model.User) error {
	if user.ID == friend.ID {
		return errors.New("you cannot follow yourself")
	}

	app.GetDB().Where(&model.Follow{
		FollowerId:  user.ID,
		FollowingId: friend.ID,
	}).Limit(1).Delete(&model.Follow{})

	return nil
}

func (s FollowService) UpdateCounters(user *model.User) error {
	var following int64
	var followers int64

	db := app.GetDB()
	db.Model(&model.Follow{}).Where(model.Follow{FollowerId: user.ID}).Count(&following)
	db.Model(&model.Follow{}).Where(model.Follow{FollowingId: user.ID}).Count(&followers)

	result := db.Model(&user).Limit(1).Updates(map[string]interface{}{
		"following_count": following,
		"followers_count": followers,
	})

	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	return nil
}
