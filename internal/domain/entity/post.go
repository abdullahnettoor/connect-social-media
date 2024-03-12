package entity

import "time"

// final String? id;
// final String description;
// final String location;
// final List? mediaURL;
// final List? likes;
// final List<CommentModel>? comments;
// final String? createdDate;
// final String? updatedDate;
// final Map<String, dynamic>? user;
// final String? userId;
// bool isBlocked;

// final String id;
// final UserModel user;
// final String comment;
// final String createdDate;

type Post struct {
	ID          int64      `json:"postId,omitempty"`
	Description string     `json:"description,omitempty"`
	Location    string     `json:"location,omitempty"`
	MediaUrls    []string   `json:"mediaUrls,omitempty"`
	Likes       []*User    `json:"likes,omitempty"`
	Comments    []*Comment `json:"comments,omitempty"`
	IsBlocked   bool      `json:"isBlocked,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	User        `json:"user,omitempty"`
}

type Comment struct {
	ID        int64     `json:",omitempty"`
	Username  string    `json:",omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
