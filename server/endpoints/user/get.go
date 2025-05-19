package user

import (
	"art-gallery-server/database"
	"art-gallery-server/middleware/auth"
	"art-gallery-server/models/users"
	"art-gallery-server/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	gravatar "github.com/katy248/gravatar/pkg/url"
)

type getUserRequest struct {
	Id int `json:"id" binding:"gte=0"`
}

type getUserResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	AvatarUrl   string `json:"avatarUrl"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

func newGetUserResponse(u users.User, authorized bool) *getUserResponse {
	resp := &getUserResponse{
		Id:          u.ID,
		Name:        u.Name,
		AvatarUrl:   gravatar.NewAvatarUrl(u.Email, gravatar.DefaultImage(gravatar.DefaultRetro), gravatar.Size(512)),
		Description: u.Description,
	}
	if authorized {
		resp.Email = u.Email
	}
	return resp
}

func GetUserHandlers() []gin.HandlerFunc {
	request := &getUserRequest{}
	handlers := []gin.HandlerFunc{
		utils.BindRequest(request),
		getUser(request),
	}
	return handlers
}
func getUser(r *getUserRequest) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorized := false
		userId := r.Id
		if user, err := auth.GetUser(ctx); err == nil {
			if userId == 0 {
				userId = user.ID
			}
			authorized = userId == user.ID
		}
		var user users.User
		query := fmt.Sprintf("id = %d", userId)
		result := database.Conn.First(&user, query)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   true,
				"message": "user not found",
				"code":    404,
			})
			return
		}
		ctx.JSON(http.StatusOK, newGetUserResponse(user, authorized))
	}
}
