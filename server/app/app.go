package app

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

func GetDB() *gorm.DB {
	return instance.DB
}

func NewApp(config *config.Config) *App {
	instance = new(App)
	instance.createDatabaseConnection(config.DB)
	instance.Migrate()

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

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

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