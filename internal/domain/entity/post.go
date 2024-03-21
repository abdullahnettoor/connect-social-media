package entity

type Post struct {
	ID          string   `json:"postId,omitempty"`
	Description string   `json:"description,omitempty"`
	Location    string   `json:"location,omitempty"`
	MediaUrls   []string `json:"mediaUrls,omitempty"`
	IsBlocked   bool     `json:"isBlocked,omitempty"`
	CreatedAt   string   `json:"createdAt,omitempty"`
	UpdatedAt   string   `json:"updatedAt,omitempty"`
}

// User 		User     `json:"user,omitempty"`
// Likes       []*User    `json:"likes,omitempty"`
// Comments    []*Comment `json:"comments,omitempty"`
