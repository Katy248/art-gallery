package all

import (
	"art-gallery-server/endpoints/post/shared"
	"art-gallery-server/middleware/auth"
	"art-gallery-server/middleware/validation"
	"art-gallery-server/utils"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func Handlers() []gin.HandlerFunc {
	var u auth.User
	var r Request
	return []gin.HandlerFunc{
		auth.Authorization(&u),
		utils.ValidateRequest(&r),
		handler(&r, &u),
	}
}

type Request struct {
	Page int `json:"page"`
}

func (r *Request) Validate() error {
	return validation.GreaterOrEqual(r.Page, 0, "page")
}

const PageLimit = 20

func handler(r *Request, u *auth.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := utils.ConnectToDbOrAbort(ctx)
		var posts []shared.ResponsePost
		var postsCount int64
		db.Raw(rawSql, u.ID, PageLimit, r.Page*PageLimit).Find(&posts)
		db.Raw(rawPagesQuery).Count(&postsCount)

		pages := postsCount / PageLimit
		if pages%PageLimit != 0 {
			pages++
		}
		log.Debugf("Posts: %d, pages: %d", postsCount, pages)

		ctx.JSON(200, gin.H{
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
