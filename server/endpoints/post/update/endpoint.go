package update

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
	var r Request
	return []gin.HandlerFunc{
		auth.Middleware(),
		utils.BindRequest(&r),
		handler(&r),
	}
}

type Request struct {
	PostId         int    `json:"postId" binding:"required,gte=0"`
	Description    string `json:"description"`
	WarningMessage string `json:"warningMessage"`
}

func handler(r *Request) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post posts.Post
		result := database.Conn.Raw(rawSql, r.PostId).First(&post)
		if !utils.WrapDbResult(c, result) {
			log.Errorf("Failed get post with id %d: %v", r.PostId, result.Error)
			return
		}

		if post.PublisherID == auth.UserId(c) || auth.UserIsAdmin(c) {
		} else {
			log.Errorf("User %d is not the owner of post %d, or is not an admin", auth.UserId(c), r.PostId)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not the owner of the post", "success": false})
			return

		}

		result = database.Conn.Model(&post).Update("description", r.Description)
		if !utils.WrapDbResult(c, result) {
			log.Errorf("Failed update post with id %d: %v", r.PostId, result.Error)
			return
		}

		if auth.UserIsAdmin(c) && r.WarningMessage != "" {
			result = database.Conn.Model(&post).Update("warning_message", r.WarningMessage)
			if !utils.WrapDbResult(c, result) {
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

const rawSql = `
	SELECT 
		p.id
		, p.description
	FROM 
		posts p 

	WHERE p.id = ?
`
