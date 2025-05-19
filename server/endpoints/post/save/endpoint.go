package save

import (
	"art-gallery-server/database"
	"art-gallery-server/middleware/auth"
	"art-gallery-server/models/posts"
	"art-gallery-server/utils"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func Handlers() []gin.HandlerFunc {
	var user auth.User
	var r request
	return []gin.HandlerFunc{
		auth.Authorization(&user),
		utils.BindRequest(&r),
		handler(&r, &user),
	}
}

type request struct {
	PostID int `json:"postId" binding:"required,gte=0"`
}

func handler(r *request, user *auth.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var existingPostSave []posts.PostSave
		res := database.Conn.Find(&existingPostSave, "post_id = ? and user_id = ?", r.PostID, user.ID)
		if res.RowsAffected != 0 {
			log.Warnf("Post with id %d is already saved by user %s", r.PostID, user.Email)
			ctx.AbortWithStatus(http.StatusOK)
			return
		}
		newSave := posts.PostSave{
			PostID: r.PostID,
			UserID: user.ID,
		}
		database.Conn.Create(&newSave)
		ctx.JSON(http.StatusOK, gin.H{"saved": true})
	}
}
