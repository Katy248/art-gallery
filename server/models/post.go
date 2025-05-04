package models

import "gorm.io/gorm"

type Post struct {
	*gorm.Model

	Description string `json:"description"` // Comment
	Publisher   User   `json:"publisher"`
	PublisherID int    `json:"publisherId"`
	ImageUrl    string `json:"imageUrl"`

	PostSaves []PostSave
}

type PostSave struct {
	*gorm.Model
	User   User
	UserID int
	Post   Post
	PostID int
}
