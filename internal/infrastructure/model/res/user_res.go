package res

import "github.com/abdullahnettoor/connect-social-media/internal/domain/entity"

type SignUpRes struct {
	Code    int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Token   string      `json:"token,omitempty"`
	Error   error       `json:"error,omitempty"`
}

type LoginRes struct {
	Code        int    `json:"status,omitempty"`
	Message     string `json:"message,omitempty"`
	Token       string `json:"token,omitempty"`
	Error       error  `json:"error,omitempty"`
	entity.User `json:"user,omitempty"`
}
