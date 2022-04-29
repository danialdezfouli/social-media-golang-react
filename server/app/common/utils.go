package common

import (
	"github.com/labstack/echo/v4"
	"jupiter/app/model"
)

func GetUser(context echo.Context) *model.User {
	return context.Get("user").(*model.User)
}

func Contains[T comparable](arr []T, elm T) bool {
	for _, x := range arr {
		if x == elm {
			return true
		}
	}
	return false

}
