package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	App "jupiter/app"
	"jupiter/app/feeds/service"
	"jupiter/app/model"
	"jupiter/config"
	"log"
)

func freshDB(app *App.App, configs *config.Config) *gorm.DB {
	con := configs.DB
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		con.Username,
		con.Password,
		con.Host,
		con.Port,
		con.Name,
		con.Charset)

	var err error
	db := app.DB
	db.Exec(fmt.Sprintf("drop database %s", "jupiter"))
	db.Exec(fmt.Sprintf("create database %s CHARACTER SET utf8 COLLATE utf8_general_ci", "jupiter"))
	db, err = gorm.Open(mysql.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to reconnect database")
	}
	app.DB = db
	return db

}

func main() {
	configs := config.GetConfig()
	app := App.NewApp(configs)

	db := freshDB(app, configs)
	app.Migrate()

	// default data

	user := &model.User{
		Username: "Danial",
		Email:    "danial@gmail.com",
	}
	db.Create(user)

	user2 := &model.User{
		Username: "Ali",
		Email:    "ali@gmail.com",
	}
	db.Create(user2)

	service.NewUserService(db).Follow(user2, user)
	//service.NewUserService(db).UnFollow(user2, user)

	post := &model.Post{
		User:    *user,
		Content: "پست نمونه #توییتر",
	}
	service.NewPostService(db).CreatePost(post).AddHashtag("توییتر")

	post2 := &model.Post{
		User:    *user,
		Content: "پست نمونه #توییت2",
	}
	service.NewPostService(db).CreatePost(post2).AddHashtag("توییتر").AddHashtag("example")

	service.NewFavoriteService(db).AddFavorite(post, user)
	service.NewFavoriteService(db).AddFavorite(post, user2)

}
