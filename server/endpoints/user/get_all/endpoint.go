package get_all

import (
	"art-gallery-server/database"
	"art-gallery-server/middleware/auth"
	"art-gallery-server/models"
	"art-gallery-server/models/users"
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
	Page     int `json:"page" binding:"gte=0"`
	PageSize int `json:"pageSize" binding:"gte=10,lte=50"`
}

type UserResponse struct {
	*models.BaseModel
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	IsAdmin     bool   `json:"isAdmin"`
}

func handler(request *Request) gin.HandlerFunc {
	return func(c *gin.Context) {
		if request.PageSize == 0 {
			request.PageSize = 10
		}
		dbUser, err := users.GetUser(auth.UserId(c))
		if err != nil {
			log.Errorf("Failed get user with id `%d` from database: %s", auth.UserId(c), err)
		}

		if !dbUser.IsAdmin {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed", "success": false})
		}

		var users []UserResponse
		database.Conn.Raw(rawQuery, request.PageSize, request.Page*request.PageSize).Find(&users)
		c.JSON(http.StatusOK, gin.H{"users": users, "success": true})
	}
}

const rawQuery = `
	SELECT * 
	FROM users u
	ORDER BY u.created_at DESC
	LIMIT ? 
	OFFSET ?
`
