package endpoints

import "github.com/gin-gonic/gin"

func CreatePostHandlers() []gin.HandlerFunc {
	handlers := []gin.HandlerFunc{createPost()}
	return handlers
}

func createPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
