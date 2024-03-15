package entity

type Comment struct {
	ID        int64     `json:"commentId,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	User      User      `json:"user,omitempty"`
}
