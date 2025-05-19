package models

type Post struct {
	*BaseModel

	Description    string `json:"description"`
	Publisher      User   `json:"publisher"`
	PublisherID    int    `json:"publisherId"`
	ImageUrl       string `json:"imageUrl"`
	WarningMessage string `json:"warningMessage"`

	PostSaves []PostSave
	Comments  []Comment
}

type PostSave struct {
	*BaseModel
	User   User
	UserID int
	Post   Post
	PostID int
}

type Comment struct {
	*BaseModel
	Post      Post
	PostID    int
	Creator   User
	CreatorID int

	Text string
}
