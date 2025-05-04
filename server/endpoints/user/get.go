package user

import (
	"art-gallery-server/middleware/auth"
	m "art-gallery-server/models"
	"art-gallery-server/utils"
	u "art-gallery-server/utils"
	"errors"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
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
}

func newGetUserResponse(u m.User) *getUserResponse {
	return &getUserResponse{
		Id:        u.ID,
		Name:      u.Name,
		AvatarUrl: gravatar.NewAvatarUrl(u.Email, gravatar.DefaultImage(gravatar.DefaultRetro), gravatar.Size(512)),
	}
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
		userId := r.Id
		if userId <= 0 {
			if user, err := auth.GetUser(ctx); err != nil {
				log.Warn("Invalid user id with non-authorized")
			} else {
				userId = user.ID
			}
		}
		var user m.User
		query := fmt.Sprintf("id = %d", userId)
		db.Model(&m.User{}).First(&user, query)
		ctx.JSON(http.StatusOK, newGetUserResponse(user))
	}
}
