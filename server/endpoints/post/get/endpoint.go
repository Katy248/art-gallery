package get

import (
	"art-gallery-server/middleware/auth"
	"art-gallery-server/middleware/validation"
	"art-gallery-server/utils"

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
	PostId int `json:"postId"`
}

func (r *Request) Validate() error {
	return validation.GreaterOrEqual(r.PostId, 0, "page")
}

type responsePost struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	PublisherID int    `json:"publisherId"`
	Saved       bool   `json:"saved"`
	CreatedAt   string `json:"createdAt"`
}

func handler(r *Request, user *auth.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := utils.ConnectToDbOrAbort(c)
		var post responsePost
		db.Raw(rawSql, user.ID, r.PostId).First(&post)
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
		, u.name
		, CAST(CASE WHEN ps.user_id IS NULL THEN 0 ELSE 1 END AS BOOLEAN) as saved
	FROM 
		posts p 
			LEFT JOIN users u ON p.publisher_id = u.id
			LEFT JOIN post_saves ps ON p.id = ps.post_id and ps.user_id = ?

	WHERE p.id = ?
`
