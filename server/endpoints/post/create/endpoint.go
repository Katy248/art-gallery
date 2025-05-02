package create

import (
	"art-gallery-server/database"
	auth "art-gallery-server/middleware/auth"
	"art-gallery-server/models"
	"art-gallery-server/utils"
	"errors"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type createPostRequest struct {
	Description string `form:"description"`
}

func (r *createPostRequest) Validate() error {
	return errors.Join()
}

func CreatePostHandlers() []gin.HandlerFunc {
	var user auth.AuthUser
	// var request createPostRequest
	return []gin.HandlerFunc{
		auth.Authorization(&user),
		// ValidateRequest(&request),
		createPost(&user),
	}
}

func createPost(user *auth.AuthUser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request createPostRequest
		err := ctx.Bind(&request)
		if err != nil {
			log.Errorf("Failed bind form data: %s", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		file, err := ctx.FormFile("picture")
		if err != nil {
			log.Errorf("Failed get file: %s", err)
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		db := utils.ConnectToDbOrAbort(ctx)
		post := models.Post{
			PublisherID: user.ID,
			Description: request.Description,
		}

		db.Model(&post).Create(&post)

		imageName := fmt.Sprintf("picture-%d", post.ID)
		post.ImageUrl = fmt.Sprintf("/api/images/%s", imageName)

		db.Save(&post)
		err = ctx.SaveUploadedFile(file, fmt.Sprintf("%s/%s", database.ImagesDir(), imageName))
		if err != nil {
			log.Errorf("Failed save file: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"Success": true})
	}
}
