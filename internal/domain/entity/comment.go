package entity

type Comment struct {
	ID        string `json:"commentId,omitempty"`
	Comment   string `json:"comment,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	PostID    string `json:"postId,omitempty"`
	UserID    string `json:"userId,omitempty"`
	Username  string `json:"username,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}
