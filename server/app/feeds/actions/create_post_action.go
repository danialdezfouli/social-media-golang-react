package actions

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"jupiter/app"
	"jupiter/app/common"
	"jupiter/app/feeds/dto"
	"jupiter/app/feeds/repository"
	"jupiter/app/feeds/service"
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
	postService := service.NewPostService(db, c)
	postType := model.PostTypePost
	parentId := sql.NullInt32{
		Int32: 0,
		Valid: false,
	}

	var parent model.Post
	if input.Type == model.PostTypeReply {
		result := db.Where("post_id", input.ReplyToID).First(&parent)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "reply to post was not found")
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

	if postType == model.PostTypeReply {
		postService.UpdatePostCounters(&model.Post{PostId: parent.PostId})
	}

	response := repository.Post{
		PostId:          post.PostId,
		UserId:          post.UserId,
		ProfileName:     user.Name,
		ProfileImage:    user.Image,
		ProfileUsername: user.Username,
		ParentId:        parent.PostId,
		PostType:        post.PostType,
		Content:         post.Content,
		CreatedAt:       post.CreatedAt,
	}

	return c.JSON(http.StatusCreated, response)
}
