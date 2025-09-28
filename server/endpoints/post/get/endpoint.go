package get

import (
	"art-gallery-server/database"
	"art-gallery-server/endpoints/post/shared"
	"art-gallery-server/middleware/auth"
	"art-gallery-server/utils"
	"net/http"

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
	PostId int `json:"postId"`
}

func handler(r *Request) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post shared.ResponsePost
		database.Conn.Raw(rawSql, auth.UserId(c), r.PostId).First(&post)

		post.UpdateSavesCount()

		c.JSON(http.StatusOK, gin.H{
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
