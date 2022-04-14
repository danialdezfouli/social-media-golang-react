package main

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
	"jupiter/app"
	"jupiter/app/model"
	"jupiter/config"
	"strings"
	"time"
)

func main() {
	configs := config.GetConfig()
	db := app.NewApp(configs).DB

	//db.Unscoped().Where("1=1").Delete(model.User{})
	for i := 0; i < 2; i++ {
		user := new(model.User)
		err := faker.FakeData(user)
		user.Email = strings.ToLower(user.Email)

		if err == nil {
			db.Create(user)
			createPosts(db, user)

		}

	}
}

func createPosts(db *gorm.DB, user *model.User) {
	for i := 0; i < 5; i++ {
		post := &model.Post{
			User: *user,
		}
		err := faker.FakeData(post)

		post.CreatedAt = time.Now()

		if err == nil {
			db.Create(post)
		}
	}
}
