package req

import (
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
