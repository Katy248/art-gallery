package delete

import (
	"art-gallery-server/database"
	"art-gallery-server/middleware/auth"
	"art-gallery-server/models/posts"
	"art-gallery-server/models/users"
	"art-gallery-server/utils"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type Request struct {
	PostId int `json:"postId" binding:""`
}

func Handlers() []gin.HandlerFunc {
	var user auth.User
	var r Request
	return []gin.HandlerFunc{
		auth.Authorization(&user),
		utils.BindRequest(&r),
		handler(&r, &user),
	}
}

func handler(request *Request, user *auth.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var post posts.Post
		result := database.Conn.Raw(rawGetPostQuery, request.PostId).First(&post)
		if result.Error != nil {
			log.Errorf("Failed to get post: %s", result.Error)
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Failed to get post", "success": false})
			return
		}

		if post.PublisherID == user.ID || users.IsAdmin(user.ID) {
		} else {
			log.Warnf("Unauthorized try to delete post %d by user %d", post.ID, request.PostId)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to delete this post", "success": false})
			return

		}

		result = database.Conn.Unscoped().Delete(&post)
		if result.Error != nil {
			log.Errorf("Failed to delete post: %s", result.Error)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post", "success": false})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"success": true})
	}
}

var rawGetPostQuery = `
	SELECT * FROM posts WHERE id = ?
`
