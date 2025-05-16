package save

import (
	"art-gallery-server/middleware/auth"
	"art-gallery-server/middleware/validation"
	"art-gallery-server/models"
	"art-gallery-server/utils"
	"errors"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func Handlers() []gin.HandlerFunc {
	var user auth.User
	var r request
	return []gin.HandlerFunc{
		auth.Authorization(&user),
		utils.ValidateRequest(&r),
		handler(&r, &user),
	}
}

type request struct {
	PostID int `json:"postId"`
}

func (r *request) Validate() error {
	return errors.Join(
		validation.GreaterOrEqual(r.PostID, 0, "postID"),
	)
}

func handler(r *request, user *auth.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := utils.ConnectToDbOrAbort(ctx)
		var existingPostSave []models.PostSave
		res := db.Find(&existingPostSave, "post_id = ? and user_id = ?", r.PostID, user.ID)
		if res.RowsAffected != 0 {
			log.Warnf("Post with id %d is already saved by user %s", r.PostID, user.Email)
			ctx.AbortWithStatus(http.StatusOK)
			return
		}
		newSave := models.PostSave{
			PostID: r.PostID,
			UserID: user.ID,
		}
		db.Create(&newSave)
		ctx.JSON(http.StatusOK, gin.H{"saved": true})
	}
}
