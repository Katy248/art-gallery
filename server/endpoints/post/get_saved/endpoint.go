package get_saved

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
	var r request
	return []gin.HandlerFunc{
		auth.Middleware(),
		utils.BindRequest(&r),
		handler(&r),
	}
}

type request struct {
	UserID int `json:"userId" binding:"required,gt=0"`
	Page   int `json:"page"`
}

func handler(r *request) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var posts []*shared.ResponsePost
		var count int64
		database.Conn.Raw(pagesQuery, auth.UserId(ctx), r.UserID).Count(&count)
		database.Conn.Raw(rawQuery, auth.UserId(ctx), r.UserID).Offset(r.Page * shared.PageSize).Limit(shared.PageSize).Find(&posts)

		pages := count / shared.PageSize
		if count%shared.PageSize != 0 {
			pages++
		}
		for _, post := range posts {
			post.UpdateSavesCount()
		}
		log.Debugf("Count: %d, pages: %d", count, pages)
		ctx.JSON(http.StatusOK, gin.H{"posts": posts, "pages": pages})
	}
}

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
