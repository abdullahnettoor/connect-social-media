package entity

type Comment struct {
	ID        string     `json:"commentId,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	User      User      `json:"user,omitempty"`
}
