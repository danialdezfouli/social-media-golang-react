package main

import (
	"jupiter/app"
	"jupiter/config"
	"jupiter/server"
)

func main() {
	configs := config.GetConfig()
	jupiter := app.NewApp(configs)
	rest := server.NewRest(jupiter)

	rest.Listen(configs.App.Url)

}
