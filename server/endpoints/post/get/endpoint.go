package get

import (
	"art-gallery-server/database"
	"art-gallery-server/endpoints/post/shared"
	"art-gallery-server/middleware/auth"
	"art-gallery-server/utils"

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
	PostId int `json:"postId"`
}

func handler(r *Request, user *auth.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post shared.ResponsePost
		database.Conn.Raw(rawSql, user.ID, r.PostId).First(&post)
		// db.Where("publisher_id = ?", user.ID).Offset(r.Page * 20).Limit(20).Find(&posts)
		c.JSON(200, gin.H{
			"success": true,
			"post":    post,
		})
	}
}

const rawSql = `
	SELECT 
		p.id
		, p.description
		, p.image_url
		, p.publisher_id
		, p.created_at
		, p.warning_message
		, u.name
		, CAST(CASE WHEN ps.user_id IS NULL THEN 0 ELSE 1 END AS BOOLEAN) as saved
	FROM 
		posts p 
			LEFT JOIN users u ON p.publisher_id = u.id
			LEFT JOIN post_saves ps ON p.id = ps.post_id and ps.user_id = ?

	WHERE p.id = ?
`
