package main

import (
	"database/sql"
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
	"jupiter/app"
	"jupiter/app/common/bcrypt"
	"jupiter/app/feeds/service"
	"jupiter/app/model"
	relationshipService "jupiter/app/relationship/service"
	"jupiter/config"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	configs := config.GetConfig()
	db := app.NewApp(configs).DB

	password, _ := bcrypt.Hash("123456")
	admin := &model.User{
		Username: "danial",
		Name:     "Danial",
		Email:    "danial@gmail.com",
		Password: password,
	}

	result := db.FirstOrCreate(admin)

	createUsers(db)
	createPosts(db)

	if result.RowsAffected > 0 {
		feedUserWithPosts(admin)
	}
	feedUsersWithFollowing()

	// TODO: add replies to posts
	// TODO: add favorite to posts
	//service.NewFavoriteService(db).AddFavorite(post, user)
}

func feedUsersWithFollowing() {
	db := app.GetDB()
	for i := 0; i < 15; i++ {
		var following *model.User
		db.Model(&model.User{}).Order("rand()").First(&following)

		var follower *model.User
		db.Model(&model.User{}).Order("rand()").First(&follower)
		relationshipService.NewFollowService().Follow(follower, following.ID)
	}
}

func feedUserWithPosts(user *model.User) {
	db := app.GetDB()

	post1 := &model.Post{
		User:     *user,
		ParentId: sql.NullInt32{},
		Content:  "این اولین پست من است #توییترـفارسی",
	}

	service.NewPostService(db).CreatePost(post1)

	post2 := &model.Post{
		User: *user,
		ParentId: sql.NullInt32{
			Int32: int32(post1.PostId),
			Valid: true,
		},
		PostType: "reply",
		Content:  "‌سلام به همه این دومین #توییت من است",
	}

	service.NewPostService(db).CreatePost(post2)
	service.NewPostService(db).UpdatePostCounters(post1)

}

func createUsers(db *gorm.DB) {
	if count := db.Find(&[]model.User{}).RowsAffected; count < 10 {
		for i := 0; i < 10-int(count); i++ {

			user := new(model.User)
			err := faker.FakeData(user)
			user.Email = strings.ToLower(user.Email)

			if err != nil {
				log.Fatal(err)
			}

			db.Create(user)
		}
	}

}

func createPosts(db *gorm.DB) {
	var users []model.User
	db.Find(&users)

	for i := 0; i < 30; i++ {
		user := users[rand.Intn(len(users))]

		post := &model.Post{
			User: user,
		}

		err := faker.FakeData(post)
		post.CreatedAt = time.Now()

		if err != nil {
			log.Fatal(err)
		}

		db.Create(post)
	}

}
