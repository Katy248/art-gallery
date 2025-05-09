package models

type Post struct {
	*BaseModel

	Description string `json:"description"` // Comment
	Publisher   User   `json:"publisher"`
	PublisherID int    `json:"publisherId"`
	ImageUrl    string `json:"imageUrl"`

	PostSaves []PostSave
}

type PostSave struct {
	*BaseModel
	User   User
	UserID int
	Post   Post
	PostID int
}
