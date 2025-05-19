package user

import (
	"art-gallery-server/middleware/auth"
	"art-gallery-server/models/users"
	"art-gallery-server/utils"
	"net/http"

	"github.com/charmbracelet/log"
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
		auth.Middleware(),
		utils.BindRequest(request),
		getUser(request),
	}
	return handlers
}
func getUser(r *getUserRequest) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorized := false
		userId := r.Id

		requestedUserId := auth.UserId(ctx)
		if userId == 0 {
			userId = requestedUserId
		}
		authorized = userId == requestedUserId

		user, err := users.GetUser(userId)
		if err != nil {
			log.Errorf("Failed get user with id '%d': %s", userId, err)
			ctx.JSON(http.StatusNotFound, gin.H{
				"statusCode": http.StatusNotFound,
				"success":    false,
			})
			return
		}
		ctx.JSON(http.StatusOK, newGetUserResponse(user, authorized))
	}
}
