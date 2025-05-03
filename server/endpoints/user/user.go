package user

import (
	"errors"
	"fmt"
	"net/http"

	"art-gallery-server/middleware/auth"
	m "art-gallery-server/models"
	"art-gallery-server/utils"

	u "art-gallery-server/utils"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	gravatar "github.com/katy248/gravatar/pkg/url"
)

type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *createUserRequest) Validate() error {
	return errors.Join(
		u.ValidateNotEmptyNamed(r.Name, "name"),
		u.ValidateNotEmptyNamed(r.Email, "email"),
		u.ValidateNotEmptyNamed(r.Password, "password"),
	)
}

func CreateUserHandlers() []gin.HandlerFunc {
	request := &createUserRequest{}
	handlers := []gin.HandlerFunc{
		u.ValidateRequest(request), createUser(request),
	}
	return handlers
}

func createUser(r *createUserRequest) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := m.NewUser(r.Name, r.Email, r.Password)
		if err != nil {
			log.Errorf("Error while creating new user: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		db := u.ConnectToDbOrAbort(ctx)
		if err := user.Save(db); err != nil {
			log.Errorf("Failed to save user: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.AbortWithStatus(http.StatusOK)

	}
}

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
	AvatarUrl string `json:"avatar.url"`
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

func newGetUserResponse(u m.User) *getUserResponse {
	return &getUserResponse{
		Id:        u.ID,
		Name:      u.Name,
		AvatarUrl: gravatar.NewAvatarUrl(u.Email, gravatar.DefaultImage(gravatar.DefaultRetro)),
	}
}

type editRequest struct {
	Name string `json:"name"`
}

func (r *editRequest) Validate() error {
	return errors.Join(u.ValidateNotEmptyNamed(r.Name, "name"))
}

func EditUserHandlers() []gin.HandlerFunc {
	var user auth.AuthUser
	var request editRequest
	handlers := []gin.HandlerFunc{
		auth.Authorization(&user),
		u.ValidateRequest(&request),
		editUser(&request, &user),
	}
	return handlers
}
func editUser(r *editRequest, user *auth.AuthUser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := u.ConnectToDbOrAbort(ctx)

		var dbUser m.User
		query := fmt.Sprintf("id = %d", user.ID)
		result := db.Model(&m.User{}).First(&dbUser, query)

		if result.Error != nil {
			log.Errorf("Failed get user (id = %d) from database: %s", user.ID, result.Error)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		db.Model(&dbUser).Update("Name", r.Name)
		ctx.AbortWithStatus(http.StatusOK)
	}
}
