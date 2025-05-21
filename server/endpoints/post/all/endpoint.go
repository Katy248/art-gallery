package all

import (
	"art-gallery-server/database"
	"art-gallery-server/endpoints/post/shared"
	"art-gallery-server/middleware/auth"
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
	Page int `json:"page" binding:"gte=0"`
}

func handler(r *Request) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var posts []shared.ResponsePost
		var postsCount int64
		database.Conn.Raw(rawSql, auth.UserId(ctx), shared.PageSize, r.Page*shared.PageSize).Find(&posts)
		database.Conn.Raw(rawPagesQuery).Count(&postsCount)

		pages := postsCount / shared.PageSize
		if pages%shared.PageSize != 0 {
			pages++
		}
		log.Debugf("Posts: %d, pages: %d", postsCount, pages)

		ctx.JSON(http.StatusOK, gin.H{
			"posts": posts,
			"pages": pages,
		})
	}
}

const rawSql = `
	SELECT 
		p.id
		, p.description
		, p.created_at
		, p.image_url
		, p.publisher_id
		, p.warning_message
		, u.name as publisher_name
		, CAST(CASE WHEN ps.user_id IS NULL THEN 0 ELSE 1 END AS BOOLEAN) as saved
	FROM 
		posts p 
			LEFT JOIN users u ON p.publisher_id = u.id
			LEFT JOIN post_saves ps ON p.id = ps.post_id and ps.user_id = ?

	ORDER BY 
		p.created_at DESC
	LIMIT ?
	OFFSET ?
`
const rawPagesQuery = `
	SELECT count(*)
	FROM posts
`
