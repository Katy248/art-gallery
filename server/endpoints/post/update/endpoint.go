package update

import (
	"art-gallery-server/database/users"
	"art-gallery-server/middleware/auth"
	"art-gallery-server/models"
	"art-gallery-server/utils"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func Handlers() []gin.HandlerFunc {
	var u auth.User
	var r Request
	return []gin.HandlerFunc{
		auth.Authorization(&u),
		utils.BindRequest(&r),
		handler(&r, &u),
	}
}

type Request struct {
	PostId         int    `json:"postId" binding:"required,gte=0"`
	Description    string `json:"description"`
	WarningMessage string `json:"warningMessage"`
}

func handler(r *Request, user *auth.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := utils.ConnectToDbOrAbort(c)
		var post models.Post
		result := db.Raw(rawSql, r.PostId).First(&post)
		if !utils.WrapDbResult(c, result) {
			log.Errorf("Failed get post with id %d: %v", r.PostId, result.Error)
			return
		}

		if post.PublisherID == user.ID || users.IsAdmin(db, user.ID) {
		} else {
			log.Errorf("User %d is not the owner of post %d, or is not an admin", user.ID, r.PostId)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not the owner of the post", "success": false})
			return

		}

		result = db.Model(&post).Update("description", r.Description)
		if !utils.WrapDbResult(c, result) {
			log.Errorf("Failed update post with id %d: %v", r.PostId, result.Error)
			return
		}

		if users.IsAdmin(db, user.ID) && r.WarningMessage != "" {
			result = db.Model(&post).Update("warning_message", r.WarningMessage)
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
