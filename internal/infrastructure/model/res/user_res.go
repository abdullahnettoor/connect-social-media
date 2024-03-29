package res

import "github.com/abdullahnettoor/connect-social-media/internal/domain/entity"

type SignUpRes struct {
	CommonRes
	Token string `json:"token,omitempty"`
}

type LoginRes struct {
	CommonRes
	Token       string `json:"token,omitempty"`
	entity.User `json:"user,omitempty"`
}

type CreatePostRes struct {
	CommonRes
}

type UserProfileRes struct {
	CommonRes
	Followers []*entity.User `json:"followers,omitempty"`
	Following []*entity.User `json:"following,omitempty"`
}

