package entity

type Post struct {
	ID          int64  `json:"postId,omitempty"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
	IsBlocked   bool   `json:"isBlocked,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
	User        User   `json:"user,omitempty"`
}

// MediaUrls   []string   `json:"mediaUrls,omitempty"`
// Likes       []*User    `json:"likes,omitempty"`
// Comments    []*Comment `json:"comments,omitempty"`
