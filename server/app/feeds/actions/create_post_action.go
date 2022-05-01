package actions

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"jupiter/app"
	"jupiter/app/common"
	"jupiter/app/feeds/dto"
	"jupiter/app/model"
	"net/http"
)

func AddPost(c echo.Context) error {
	input := new(dto.CreatePostDTO)
	err := c.Bind(input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(input); err != nil {
		return err
	}

	user := common.GetUser(c)
	db := app.GetDB()

	postType := "post"
	parentId := sql.NullInt32{
		Int32: 0,
		Valid: false,
	}

	var parent *model.Post
	if input.Type == "reply" {
		result := db.Where("post_id", input.ReplyToID).First(&parent)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "post "+string(input.ReplyToID)+" was not found")
		}

		postType = "reply"
		parentId = sql.NullInt32{
			Int32: int32(parent.PostId),
			Valid: true,
		}
	}

	post := model.Post{
		UserId:   user.ID,
		ParentId: parentId,
		PostType: postType,
		Content:  input.Body,
	}

	db.Create(&post)

	return c.JSON(http.StatusOK, post)
}
