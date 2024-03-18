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

type UserId struct {
	UserID int64 `json:"userId"`
}

type VerifyOtp struct {
	UserID int64  `json:"userId"`
	Otp    string `json:"Otp"`
}

type CreatePostReq struct {
	Description string `form:"description,omitempty"`
	Location    string `form:"location,omitempty"`
	Files       []*multipart.FileHeader
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
