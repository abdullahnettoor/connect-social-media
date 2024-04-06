package req

import (
	"mime/multipart"

	"github.com/abdullahnettoor/connect-social-media/internal/domain/constants"
)

type SignUpReq struct {
	FullName  string            `json:"fullName"`
	Email     string            `json:"email"`
	Username  string            `json:"username"`
	Password  string            `json:"password"`
	AccType   constants.AccType `json:"accountType"`
	CreatedAt string            `json:"createdAt"`
	UpdatedAt string            `json:"updatedAt"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type FollowUnfollowUserReq struct {
	UserID     string `json:"userId"`
	FollowedID string `json:"followedId"`
}

type UserId struct {
	UserID string `json:"userId"`
}

type VerifyOtp struct {
	UserID string `json:"userId"`
	Otp    string `json:"Otp"`
}

type CreatePostReq struct {
	Description string `form:"description,omitempty"`
	Location    string `form:"location,omitempty"`
	Files       []*multipart.FileHeader
	UserID      string `json:"userId"`
}

type LikeUnlikePostReq struct {
	PostID string `json:"postId"`
	UserID string `json:"userId"`
}

type CreateCommentReq struct {
	Comment string `json:"comment,omitempty"`
	UserID  string `json:"userId"`
	PostID  string `json:"postId"`
}

type GetCommentsReq struct {
	PostID string `json:"postId"`
}

type DeleteCommentReq struct {
	UserID    string `json:"userId"`
	CommentID string `json:"commentId"`
}

type SendChatReq struct {
	SenderID    string `json:"senderId"`
	RecipientID string `json:"recipientId"`
	Message     string `json:"message"`
	CreatedAt   string `json:"createdAt"`
	ReceivedAt  string `json:"receivedAt"`
}

// type Post struct {
// 	ID          int64    `json:"postId,omitempty"`
// 	Description string   `json:"description,omitempty"`
// 	MediaUrls   []string `json:"mediaUrls,omitempty"`
// 	IsBlocked   bool     `json:"isBlocked,omitempty"`
// 	CreatedAt   string   `json:"createdAt,omitempty"`
// 	UpdatedAt   string   `json:"updatedAt,omitempty"`
// 	User        User     `json:"user,omitempty"`
// }
