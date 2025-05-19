package get_users

import (
	"art-gallery-server/database"
	"art-gallery-server/endpoints/post/shared"
	"art-gallery-server/middleware/auth"
	"art-gallery-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsersPostHandlers() []gin.HandlerFunc {
	var r request
	return []gin.HandlerFunc{
		auth.Middleware(),
		utils.BindRequest(&r),
		handler(&r),
	}
}

type request struct {
	UserID int `json:"userId" binding:"required,gt=0"`
	Page   int `json:"page" binding:"gte=0"`
}

func handler(r *request) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var posts []shared.ResponsePost
		database.Conn.Raw(rawSql, auth.UserId(ctx), r.UserID).
			Offset(shared.PageSize * r.Page).
			Limit(shared.PageSize).
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
		, p.created_at
		, p.warning_message
		, u.name
		, CAST(CASE WHEN ps.user_id IS NULL THEN 0 ELSE 1 END AS BOOLEAN) as saved
	FROM 
		posts p 
			LEFT JOIN users u ON p.publisher_id = u.id
			LEFT JOIN post_saves ps ON 
				p.id = ps.post_id 
				and ps.user_id = ?

	WHERE publisher_id = ?
	
	ORDER BY 
		p.created_at DESC
`
