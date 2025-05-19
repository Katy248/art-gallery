package posts

import (
	"art-gallery-server/models"
	"art-gallery-server/models/users"
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
