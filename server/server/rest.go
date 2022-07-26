package server

import (
	"jupiter/app/common/validation"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Rest struct {
	echo *echo.Echo
}

func NewRest() (rest *Rest) {
	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		MaxAge: 300,
		// AllowOrigins: []string{"http://localhost:3000", "https://jupiter-client.vercel.app"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	rest = &Rest{
		echo: e,
	}

	rest.routes()

	return
}

func (r *Rest) Listen(url string) {
	r.echo.Logger.Fatal(r.echo.Start(url))
}
