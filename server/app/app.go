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
	instance.createDatabaseConnection(config)
	instance.Migrate()

	return instance
}

func (app *App) createDatabaseConnection(config *config.Config) {

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	var logProvider logger.Interface
	if !config.App.Production {
		logProvider = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{
		Logger: logProvider,
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
