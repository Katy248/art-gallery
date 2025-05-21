package create

import (
	"art-gallery-server/database"
	auth "art-gallery-server/middleware/auth"
	"art-gallery-server/models/posts"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func CreatePostHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		auth.Middleware(),
		createPost(),
	}
}

func createPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Description string `form:"description" json:"description"`
		}
		if err := ctx.ShouldBind(&request); err != nil {
			log.Errorf("Failed bind request data: %s", err)
		}

		file, err := ctx.FormFile("picture")
		if err != nil {
			log.Errorf("Failed get file: %s", err)
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		post := posts.Post{
			PublisherID: auth.UserId(ctx),
			Description: request.Description,
		}

		database.Conn.Model(&post).Create(&post)

		imageName := fmt.Sprintf("picture-%d", post.ID)
		post.ImageUrl = fmt.Sprintf("/api/images/%s", imageName)

		database.Conn.Save(&post)
		err = ctx.SaveUploadedFile(file, fmt.Sprintf("%s/%s", database.ImagesDir(), imageName))
		if err != nil {
			log.Errorf("Failed save file: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"success": true, "id": post.ID})
	}
}
