package unsave

import (
	"art-gallery-server/middleware/auth"
	"art-gallery-server/middleware/validation"
	"art-gallery-server/models"
	"art-gallery-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handlers() []gin.HandlerFunc {
	var u auth.AuthUser
	var r request
	return []gin.HandlerFunc{
		auth.Authorization(&u),
		utils.ValidateRequest(&r),
		handler(&r, &u),
	}
}

type request struct {
	PostID int `json:"postId"`
}

func (r *request) Validate() error {
	return validation.GreaterOrEqual(r.PostID, 0, "postID")
}

func handler(r *request, u *auth.AuthUser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := utils.ConnectToDbOrAbort(ctx)
		var existingSave models.PostSave
		db.First(&existingSave, "post_id = ? and user_id = ?", r.PostID, u.ID)
		if existingSave.ID != 0 {
			db.Delete(&existingSave)
		}

		ctx.JSON(http.StatusOK, gin.H{"saved": false})
	}
}
