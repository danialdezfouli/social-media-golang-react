package app

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jupiter/app/model"
	"jupiter/config"
	"log"
)

type App struct {
	DB *gorm.DB
}

var instance *App

func GetInstance() *App {
	return instance
}

func NewApp(config *config.Config) *App {
	instance = new(App)
	instance.createDatabaseConnection(config.DB)

	return instance
}

func (app *App) createDatabaseConnection(config *config.DBConfig) {

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset)

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	app.DB = db

}

func (app *App) Migrate() {
	app.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Activity{},
		&model.Hashtag{},
		&model.HashtagPost{},
		&model.Favorite{},
		&model.Follow{},
	)
}
