package feeds

import (
	"github.com/labstack/echo/v4"
	"jupiter/app/common/middleware"
	"jupiter/app/feeds/actions"
)

func Routes(e *echo.Echo) {
	r := e.Group("/")
	r.Use(middleware.AuthMiddleware(), middleware.UserMiddleware())

	r.GET("timeline", homeTimeline)
	r.GET("search", search)
	r.GET("post/:id", findPost)

	r.POST("post/:id/like", actions.LikePost)
	r.POST("post", actions.AddPost)

	r.GET("profile/:id", profile)
	r.GET("profile/:id/timeline", profileTimeline)
	r.GET("profile/:id/likes", profileLikes)
}
