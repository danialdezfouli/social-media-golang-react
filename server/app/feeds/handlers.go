package feeds

import (
	"github.com/labstack/echo/v4"
	"jupiter/app"
	"jupiter/app/common"
	"jupiter/app/feeds/dto"
	"jupiter/app/feeds/repository"
	"jupiter/app/feeds/service"
	"jupiter/app/model"
	"net/http"
)

func homeTimeline(c echo.Context) error {
	params := new(dto.TimelineDTO)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	postService := service.NewPostService(app.GetDB(), c)

	user := common.GetUser(c)
	var posts = &[]repository.Post{}
	var profiles = &[]repository.SearchProfile{}

	service.QueryTimeline(int(params.Offset), user).
		Joins("inner join follows on follows.follower_id = ?", user.ID).
		Where("posts.user_id = follows.following_id").
		Find(posts)

	service.QuerySuggestedProfiles(user).Find(profiles)

	parents := postService.FindParentsForTimeline(posts)
	parentsMap := postService.KeyByPostId(parents)

	return c.JSON(http.StatusOK, echo.Map{
		"parents":            parentsMap,
		"suggested_profiles": profiles,
		"posts":              posts,
	})
}

func search(c echo.Context) error {
	params := new(dto.SearchDTO)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var users = &[]repository.SearchProfile{}

	// TODO: add tags search
	service.SearchUsers(params.Query).Find(users)

	return c.JSON(http.StatusOK, echo.Map{
		"profiles": users,
	})
}

func findPost(c echo.Context) error {
	params := new(dto.PostDTO)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db := app.GetDB()
	postService := service.NewPostService(db, c)

	post, err := postService.FindPost(params.ID)
	replies := postService.FindReplies(post)
	parents := postService.FindPostParentsHirarchy(post)

	if err != nil {
		return echo.ErrNotFound
	}

	var parentsMap = map[uint]repository.Post{}

	for _, parent := range parents {
		parentsMap[parent.PostId] = parent
	}

	return c.JSON(http.StatusOK, echo.Map{
		"post":    post,
		"replies": replies,
		"parents": parentsMap,
	})
}

func likePost(c echo.Context) error {
	params := new(dto.PostDTO)
	user := common.GetUser(c)
	err := c.Bind(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db := app.GetDB()
	postService := service.NewPostService(db, c)
	post, err := postService.FindPost(params.ID)
	liked := postService.ToggleLike(post, user)
	postService.UpdatePostCounters(&model.Post{PostId: post.PostId})

	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, echo.Map{
		"post":  post,
		"liked": liked,
	})

}
