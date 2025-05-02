package get_users

import (
	"art-gallery-server/middleware/validation"
	"art-gallery-server/models"
	"art-gallery-server/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsersPostHandlers() []gin.HandlerFunc {
	var r request
	return []gin.HandlerFunc{
		utils.ValidateRequest(&r),
		handler(&r),
	}
}

const PageLimit = 20

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
		db.Model(&models.Post{}).
			Offset(PageLimit*r.Page).
			Limit(PageLimit).
			Where("publisher_id = ?", r.UserID).
			Find(&posts)

		ctx.JSON(http.StatusOK, posts)
	}
}
