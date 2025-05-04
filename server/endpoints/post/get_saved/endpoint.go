package get_saved

import (
	"art-gallery-server/middleware/validation"
	"art-gallery-server/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handlers() []gin.HandlerFunc {
	var r request
	return []gin.HandlerFunc{
		utils.ValidateRequest(&r),
		handler(&r),
	}
}

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

func handler(r *request) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := utils.ConnectToDbOrAbort(ctx)

		var posts []responsePost
		db.Raw(rawQuery, r.UserID).Offset(r.Page * PageSize).Limit(PageSize).Find(&posts)
		ctx.JSON(http.StatusOK, posts)
	}
}

const PageSize = 20

type responsePost struct {
	ID            int    `json:"id"`
	CreatedAt     string `json:"createdAt"`
	Description   string `json:"description"`
	ImageUrl      string `json:"imageUrl"`
	PublisherID   int    `json:"publisherId"`
	Saved         bool   `json:"saved"`
	PublisherName string `json:"publisherName"`
	// Name        string `json:"name"`
}

const rawQuery = `
	SELECT 
		p.id
		, p.created_at
		, p.description
		, p.image_url
		, p.publisher_id
		
		, u.name as publisher_name
		
		, CAST(CASE WHEN ps.user_id IS NULL THEN 0 ELSE 1 END AS BOOLEAN) as saved

	FROM 
		posts p
			LEFT JOIN post_saves ps ON ps.post_id = p.id
		LEFT JOIN users u ON u.id = p.publisher_id
	WHERE 
		saved = true and ps.user_id = ?
`
