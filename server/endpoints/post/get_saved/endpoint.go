package get_saved

import (
	"art-gallery-server/endpoints/post/shared"
	"art-gallery-server/middleware/auth"
	"art-gallery-server/middleware/validation"
	"art-gallery-server/utils"
	"errors"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func Handlers() []gin.HandlerFunc {
	var u auth.User
	var r request
	return []gin.HandlerFunc{
		auth.Authorization(&u),
		utils.ValidateRequest(&r),
		handler(&r, &u),
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

func handler(r *request, user *auth.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := utils.ConnectToDbOrAbort(ctx)

		var posts []shared.ResponsePost
		var count int64
		db.Raw(pagesQuery, user.ID, r.UserID).Count(&count)
		db.Raw(rawQuery, user.ID, r.UserID).Offset(r.Page * PageSize).Limit(PageSize).Find(&posts)

		pages := count / PageSize
		if count%PageSize != 0 {
			pages++
		}
		log.Debugf("Count: %d, pages: %d", count, pages)
		ctx.JSON(http.StatusOK, gin.H{"posts": posts, "pages": pages})
	}
}

const PageSize = 20

const rawQuery = `
	SELECT 
		p.id
		, p.description
		, p.image_url
		, p.publisher_id
		, p.warning_message
		, p.created_at
		, u.name as publisher_name
		, CAST(CASE WHEN ps.user_id IS NULL THEN 0 ELSE 1 END AS BOOLEAN) as saved
	FROM 
		posts p 
			LEFT JOIN users u ON p.publisher_id = u.id
			LEFT JOIN post_saves ps ON 
				p.id = ps.post_id 
				and ps.user_id = ?

	WHERE (SELECT COUNT(*) FROM post_saves WHERE post_id = p.id and user_id = ?) <> 0
	
	ORDER BY 
		p.created_at DESC
`
const pagesQuery = `
	SELECT 
		COUNT(*)
	FROM 
		posts p 
			LEFT JOIN users u ON p.publisher_id = u.id
			LEFT JOIN post_saves ps ON 
				p.id = ps.post_id 
				and ps.user_id = ?

	WHERE (SELECT COUNT(*) FROM post_saves WHERE post_id = p.id and user_id = ?) <> 0
	
	ORDER BY 
		p.created_at DESC
`
