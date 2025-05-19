package unsave

import (
	"art-gallery-server/database"
	"art-gallery-server/middleware/auth"
	"art-gallery-server/models/posts"
	"art-gallery-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handlers() []gin.HandlerFunc {
	var u auth.User
	var r request
	return []gin.HandlerFunc{
		auth.Authorization(&u),
		utils.BindRequest(&r),
		handler(&r, &u),
	}
}

type request struct {
	PostID int `json:"postId" binding:"required,gte=0"`
}

func handler(r *request, u *auth.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var existingSave posts.PostSave
		database.Conn.Unscoped().First(&existingSave, "post_id = ? and user_id = ?", r.PostID, u.ID)
		if existingSave.ID != 0 {
			database.Conn.Unscoped().Delete(&existingSave)
		}

		ctx.JSON(http.StatusOK, gin.H{"saved": false})
	}
}
