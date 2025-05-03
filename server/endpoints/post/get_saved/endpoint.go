package get_saved

import (
	"art-gallery-server/middleware/validation"
	"art-gallery-server/models"
	"art-gallery-server/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handlers() []gin.HandlerFunc {
	var r request
	return []gin.HandlerFunc{
		utils.ValidateRequest(&r),
		handler(&r),
	}
}

type request struct {
	UserID int `json:"userId"`
	Page   int `json:"page"`
}

func (r *request) Validate() error {
	return errors.Join(
		validation.GreaterThan(r.UserID, 0, "userID"),
		validation.GreaterOrEqual(r.Page, 0, "page"),
	)
}

func handler(r *request) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := utils.ConnectToDbOrAbort(ctx)

		var posts []models.Post
		db.Raw(rawQuery).Offset(r.Page * PageSize).Limit(PageSize).Find(&posts)
		ctx.JSON(http.StatusOK, posts)
	}
}

const PageSize = 20

const rawQuery = `
	SELECT 
		posts.id
		, posts.created_at
		, posts.description
		, posts.image_url
		, posts.publisher_id
		
		, users.id
		, users.name
	FROM 
		posts
		, post_saves ON post_saves.post_id = posts.id
		, users ON users.id = posts.publisher_id
	WHERE 
		post_saves.user_id = 1
`
