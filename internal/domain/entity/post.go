package entity

type Post struct {
	ID           string   `json:"postId,omitempty"`
	Description  string   `json:"description,omitempty"`
	Location     string   `json:"location,omitempty"`
	MediaUrls    []string `json:"mediaUrls,omitempty"`
	IsBlocked    bool     `json:"isBlocked,omitempty"`
	CreatedAt    string   `json:"createdAt,omitempty"`
	UpdatedAt    string   `json:"updatedAt,omitempty"`
	LikeCount    int64    `json:"likeCount"`
	CommentCount int64    `json:"commentCount"`
	Username     string   `json:"username,omitempty"`
	Avatar       string   `json:"avatar,omitempty"`
}

