package shared

type ResponsePost struct {
	ID             int    `json:"id"`
	CreatedAt      string `json:"createdAt"`
	Description    string `json:"description"`
	ImageUrl       string `json:"imageUrl"`
	PublisherID    int    `json:"publisherId"`
	Saved          bool   `json:"saved"`
	PublisherName  string `json:"publisherName"`
	WarningMessage string `json:"warningMessage"`
}

const PageSize = 20
