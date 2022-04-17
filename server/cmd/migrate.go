package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	App "jupiter/app"
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

	freshDB(app, configs)
	app.Migrate()

}
