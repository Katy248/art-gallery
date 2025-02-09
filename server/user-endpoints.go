package main

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *createUserRequest) validate() error {
	// TODO: implement validation
	return nil
}

func CreateUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r createUserRequest
		log.Info("Here")
		if err := ctx.BindJSON(&r); err != nil {
			log.Errorf("Failed bind JSON data: %s", err)
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if err := r.validate(); err != nil {
			log.Errorf("Create user request data validation fail: %s", err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		user, err := NewUser(r.Name, r.Email, r.Password)
		if err != nil {
			log.Errorf("Error while creating new user: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		db, err := ConnectToDb(ConnectionString)
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
	Id int `uri:"id"`
}
type getUserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func newGetUserResponse(u User) *getUserResponse {
	return &getUserResponse{
		Id:   u.ID,
		Name: u.Name,
	}
}

func GetUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r getUserRequest
		if err := ctx.BindUri(&r); err != nil {
			log.Errorf("Failed bind uri parameters: %s", err)
			ctx.AbortWithError(http.StatusNotFound, err)
			return
		}
		db, _ := ConnectToDb(ConnectionString)
		var user User
		query := fmt.Sprintf("id = %d", r.Id)
		db.Model(&User{}).First(&user, query)
		ctx.JSON(http.StatusOK, newGetUserResponse(user))
	}
}
