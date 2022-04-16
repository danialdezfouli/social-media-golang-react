package main

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
	"jupiter/app"
	"jupiter/app/model"
	"jupiter/config"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	configs := config.GetConfig()
	db := app.NewApp(configs).DB

	if count := db.Find(&[]model.User{}).RowsAffected; count < 10 {
		for i := 0; i < 10-int(count); i++ {
			createUser(db)
		}
	}

	var users []model.User
	db.Find(&users)

	for i := 0; i < 100; i++ {
		user := users[rand.Intn(len(users))]
		createPost(db, &user)
	}

}

func createUser(db *gorm.DB) {
	user := new(model.User)
	err := faker.FakeData(user)
	user.Email = strings.ToLower(user.Email)

	if err != nil {
		log.Fatal(err)
	}

	db.Create(user)
}

func createPost(db *gorm.DB, user *model.User) {
	post := &model.Post{
		User: *user,
	}

	err := faker.FakeData(post)
	post.CreatedAt = time.Now()

	if err != nil {
		log.Fatal(err)
	}

	db.Create(post)

}
