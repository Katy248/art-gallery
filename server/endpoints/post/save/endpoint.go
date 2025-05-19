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
	var r request
	return []gin.HandlerFunc{
		auth.Middleware(),
		utils.BindRequest(&r),
		handler(&r),
	}
}

type request struct {
	PostID int `json:"postId" binding:"required,gte=0"`
}

func handler(r *request) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var existingPostSave []posts.PostSave
		res := database.Conn.Find(&existingPostSave, "post_id = ? and user_id = ?", r.PostID, auth.UserId(ctx))
		if res.RowsAffected != 0 {
			log.Warnf("Post with id %d is already saved by user %d", r.PostID, auth.UserId(ctx))
			ctx.AbortWithStatus(http.StatusOK)
			return
		}
		newSave := posts.PostSave{
			PostID: r.PostID,
			UserID: auth.UserId(ctx),
		}
		database.Conn.Create(&newSave)
		ctx.JSON(http.StatusOK, gin.H{"saved": true})
	}
}
