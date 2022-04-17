package main

import (
	"jupiter/app"
	"jupiter/config"
	"jupiter/server"
)

func main() {
	configs := config.GetConfig()
	app.NewApp(configs)

	rest := server.NewRest()
	rest.Listen(configs.App.Url)
}
