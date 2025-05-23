package posts

import (
	"art-gallery-server/database"
	"art-gallery-server/models"
	"art-gallery-server/models/users"

	"github.com/charmbracelet/log"
)

type Post struct {
	*models.BaseModel

	Description    string     `json:"description"`
	Publisher      users.User `json:"publisher"`
	PublisherID    int        `json:"publisherId"`
	ImageUrl       string     `json:"imageUrl"`
	WarningMessage string     `json:"warningMessage"`

	PostSaves []PostSave
	Comments  []Comment
}

type PostSave struct {
	*models.BaseModel
	User   users.User
	UserID int
	Post   Post
	PostID int
}

type Comment struct {
	*models.BaseModel
	Post      Post
	PostID    int
	Creator   users.User
	CreatorID int

	Text string
}

func GetPostSavesCount(postId int) int {
	var count int64
	result := database.Conn.Raw("SELECT COUNT(*) FROM post_saves WHERE post_id = ?", postId).Count(&count)
	if result.Error != nil {
		log.Errorf("Error getting post saves count: %v", result.Error)
	}
	log.Infof("Post saves count: %d", count)
	return int(count)
}
