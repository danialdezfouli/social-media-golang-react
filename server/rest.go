package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"jupiter/app"
	"jupiter/app/common/validation"
)

type Rest struct {
	app  *app.App
	echo *echo.Echo
}

func NewRest(app *app.App) (rest *Rest) {
	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}

	rest = &Rest{
		app:  app,
		echo: e,
	}

	rest.routes()

	return
}

func (r *Rest) Listen(url string) {
	r.echo.Logger.Fatal(r.echo.Start(url))
}
