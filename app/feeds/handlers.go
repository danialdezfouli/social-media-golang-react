package feeds

import (
	"github.com/labstack/echo/v4"
	"jupiter/app"
	"jupiter/app/model"
	"net/http"
)

func feeds(c echo.Context) error {
	var users []model.User
	db := app.GetDB()
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
	//return c.String(http.StatusOK, "Home")
}
