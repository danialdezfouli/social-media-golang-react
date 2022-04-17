package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"jupiter/app/common/validation"
)

type Rest struct {
	echo *echo.Echo
}

func NewRest() (rest *Rest) {
	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}

	rest = &Rest{
		echo: e,
	}

	rest.routes()

	return
}

func (r *Rest) Listen(url string) {
	r.echo.Logger.Fatal(r.echo.Start(url))
}
