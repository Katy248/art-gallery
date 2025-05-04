package all

import (
	"art-gallery-server/middleware/auth"
	"art-gallery-server/middleware/validation"
	"art-gallery-server/utils"

	"github.com/gin-gonic/gin"
)

func Handlers() []gin.HandlerFunc {
	var u auth.AuthUser
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

func handler(r *Request, u *auth.AuthUser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := utils.ConnectToDbOrAbort(ctx)
		var posts []responsePost
		db.Raw(rawSql, u.ID).Offset(r.Page * 20).Limit(20).Find(&posts)

		ctx.JSON(200, gin.H{
			"posts": posts,
		})
	}
}

type responsePost struct {
	ID            int    `json:"id"`
	Description   string `json:"description"`
	ImageUrl      string `json:"imageUrl"`
	PublisherID   int    `json:"publisherId"`
	PublisherName string `json:"publisherName"`
	Saved         bool   `json:"saved"`
}

const rawSql = `
	SELECT 
		p.id
		, p.description
		, p.image_url
		, p.publisher_id
		, u.name as publisher_name
		, CAST(CASE WHEN ps.user_id IS NULL THEN 0 ELSE 1 END AS BOOLEAN) as saved
	FROM 
		posts p 
			LEFT JOIN users u ON p.publisher_id = u.id
			LEFT JOIN post_saves ps ON p.id = ps.post_id and ps.user_id = ?

	ORDER BY 
		p.created_at DESC
`
