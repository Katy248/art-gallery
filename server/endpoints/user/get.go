package user

import (
	"art-gallery-server/middleware/auth"
	m "art-gallery-server/models"
	"art-gallery-server/utils"
	u "art-gallery-server/utils"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	gravatar "github.com/katy248/gravatar/pkg/url"
)

type getUserRequest struct {
	Id int `json:"id"`
}

func (r *getUserRequest) Validate() error {
	return errors.Join(
	// u.ValidateMoreThan(r.Id, 0, "id"),
	)
}

type getUserResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
	Email     string `json:"email"`
}

func newGetUserResponse(u m.User, authorized bool) *getUserResponse {
	resp := &getUserResponse{
		Id:        u.ID,
		Name:      u.Name,
		AvatarUrl: gravatar.NewAvatarUrl(u.Email, gravatar.DefaultImage(gravatar.DefaultRetro), gravatar.Size(512)),
	}
	if authorized {
		resp.Email = u.Email
	}
	return resp
}

func GetUserHandlers() []gin.HandlerFunc {
	request := &getUserRequest{}
	handlers := []gin.HandlerFunc{
		utils.ValidateRequest(request),
		getUser(request),
	}
	return handlers
}
func getUser(r *getUserRequest) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := u.ConnectToDbOrAbort(ctx)
		authorized := false
		userId := r.Id
		if user, err := auth.GetUser(ctx); err == nil {
			if userId == 0 {
				userId = user.ID
			}
			authorized = userId == user.ID
		}
		var user m.User
		query := fmt.Sprintf("id = %d", userId)
		result := db.Model(&m.User{}).First(&user, query)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   true,
				"message": "user not found",
				"code":    404,
			})
		}
		ctx.JSON(http.StatusOK, newGetUserResponse(user, authorized))
	}
}
