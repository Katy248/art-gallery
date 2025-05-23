package shared

import (
	"art-gallery-server/models/posts"
)

type ResponsePost struct {
	ID             int    `json:"id"`
	CreatedAt      string `json:"createdAt"`
	Description    string `json:"description"`
	ImageUrl       string `json:"imageUrl"`
	PublisherID    int    `json:"publisherId"`
	Saved          bool   `json:"saved"`
	PublisherName  string `json:"publisherName"`
	WarningMessage string `json:"warningMessage"`
	SavesCount     int    `json:"savesCount"`
}

func (p *ResponsePost) UpdateSavesCount() {
	p.SavesCount = posts.GetPostSavesCount(p.ID)
}

const PageSize = 20
