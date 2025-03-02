package models

import "gorm.io/gorm"

type Post struct {
	*gorm.Model

	Description string // Comment
	Publisher   User
	PublisherID int
	ImageUrl    string

	PostSaves []PostSave
}

type PostSave struct {
	*gorm.Model
	User   User
	UserID int
	Post   Post
	PostID int
}
