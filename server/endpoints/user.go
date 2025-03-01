package endpoints

import (
	"errors"
	"fmt"
	"net/http"

	"art-gallery-server/database"
	m "art-gallery-server/models"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *createUserRequest) Validate() error {
	return errors.Join(
		ValidateNotEmpty(r.Name),
		ValidateNotEmpty(r.Email),
		ValidateNotEmpty(r.Password),
	)
}

func CreateUserHandlers() []gin.HandlerFunc {
	request := &createUserRequest{}
	handlers := []gin.HandlerFunc{
		ValidateRequest(request), createUser(request),
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
		db, err := database.ConnectToDb()
		if err != nil {
			log.Errorf("Failed to connect to database: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
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
		ValidateMoreThan(r.Id, 0, "id"),
	)
}

type getUserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUserHandlers() []gin.HandlerFunc {
	request := &getUserRequest{}
	handlers := []gin.HandlerFunc{
		ValidateRequest(request), getUser(request),
	}
	return handlers
}
func getUser(r *getUserRequest) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db, _ := database.ConnectToDb()
		var user m.User
		query := fmt.Sprintf("id = %d", r.Id)
		db.Model(&m.User{}).First(&user, query)
		ctx.JSON(http.StatusOK, newGetUserResponse(user))
	}
}

func newGetUserResponse(u m.User) *getUserResponse {
	return &getUserResponse{
		Id:   u.ID,
		Name: u.Name,
	}
}
