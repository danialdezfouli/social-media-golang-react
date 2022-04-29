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

	db.FirstOrCreate(admin)
	db.Where("1=1").Delete(model.Post{})

	feedUserWithPosts(admin)
	createUsers(db)
	createPosts(db)

	feedUsersWithFollowing()
	createReplies(db)

	createFavorites(db)

}

func createFavorites(db *gorm.DB) {
	for i := 0; i < 50; i++ {

		var user *model.User
		db.Order("rand()").First(&user)

		var post *model.Post
		db.Order("rand()").First(&post)

		db.FirstOrCreate(&model.Favorite{
			UserId: user.ID,
			PostId: post.PostId,
		})
		service.NewPostService(db, nil).UpdatePostCounters(post)
	}
}

func createReplies(db *gorm.DB) {
	var users []model.User
	db.Find(&users)

	var posts []model.Post
	db.Where("parent_id is null").Find(&posts)

	for i := 0; i < 20; i++ {
		user := users[rand.Intn(len(users))]
		parent := posts[rand.Intn(len(posts))]

		post := &model.Post{
			User: user,
			ParentId: sql.NullInt32{
				Int32: int32(parent.PostId),
				Valid: true,
			},
		}

		faker.FakeData(post)

		if post.PostType == model.PostTypeRepost {
			post.Content = ""
			post.FavoritesCount = 0
			post.RepliesCount = 0
		} else if post.PostType == model.PostTypePost {
			post.ParentId = sql.NullInt32{
				Int32: 0,
				Valid: false,
			}
		}

		service.NewPostService(db, nil).CreatePost(post)
		service.NewPostService(db, nil).UpdatePostCounters(&parent)
	}
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
	var count int64
	db.Model(&model.Post{}).Where("user_id", user.ID).Count(&count)

	if count > 2 {
		return
	}

	post1 := &model.Post{
		User:     *user,
		ParentId: sql.NullInt32{},
		Content:  "این اولین پست من است #توییتر_فارسی",
	}

	service.NewPostService(db, nil).CreatePost(post1)

	post2 := &model.Post{
		User: *user,
		ParentId: sql.NullInt32{
			Int32: int32(post1.PostId),
			Valid: true,
		},
		PostType: "reply",
		Content:  "‌سلام به همه این دومین #توییت من است",
	}

	service.NewPostService(db, nil).CreatePost(post2)
	service.NewPostService(db, nil).UpdatePostCounters(post1)

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

	for i := 0; i < 15; i++ {
		user := users[rand.Intn(len(users))]

		post := &model.Post{
			User: user,
		}

		err := faker.FakeData(post)
		post.PostType = model.PostTypePost
		post.CreatedAt = time.Now()

		if err != nil {
			log.Fatal(err)
		}

		db.Create(post)
	}

}
