package get_users

import (
	"art-gallery-server/middleware/validation"
	"art-gallery-server/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsersPostHandlers() []gin.HandlerFunc {
	var r request
	return []gin.HandlerFunc{
		utils.ValidateRequest(&r),
		handler(&r),
	}
}

const PageLimit = 20

type request struct {
	UserID int `json:"userId"`
	Page   int `json:"page"`
}

func (r *request) Validate() error {
	return errors.Join(
		validation.GreaterThan(r.UserID, 0, "userID"),
		validation.GreaterOrEqual(r.Page, 0, "page"),
	)
}

type responsePost struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	PublisherID int    `json:"publisherId"`
	Name        string `json:"name"`
	Saved       bool   `json:"saved"`
}

func handler(r *request) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := utils.ConnectToDbOrAbort(ctx)
		var posts []responsePost
		db.Raw(rawSql, r.UserID).
			Offset(PageLimit * r.Page).
			Limit(PageLimit).
			Find(&posts)

		ctx.JSON(http.StatusOK, posts)
	}
}

const rawSql = `
	SELECT 
		p.id
		, p.description
		, p.image_url
		, p.publisher_id
		, u.name
		, CAST(CASE WHEN ps.user_id IS NULL THEN 0 ELSE 1 END AS BOOLEAN) as saved
	FROM 
		posts p 
			LEFT JOIN users u ON p.publisher_id = u.id
			LEFT JOIN post_saves ps ON p.id = ps.post_id and ps.user_id = u.id

	WHERE publisher_id = ?
	
	ORDER BY 
		p.created_at DESC
`
