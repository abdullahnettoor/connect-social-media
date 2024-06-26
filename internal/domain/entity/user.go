package entity

import (
	"github.com/abdullahnettoor/connect-social-media/internal/domain/constants"
)

type User struct {
	ID          string               `json:"userId,omitempty"`
	Username    string               `json:"username,omitempty"`
	Email       string               `json:"email,omitempty"`
	Password    string               `json:"-,omitempty"`
	FullName    string               `json:"fullName,omitempty"`
	Status      constants.UserStatus `json:"status,omitempty"`
	Bio         string               `json:"bio,omitempty"`
	Avatar      string               `json:"avatar,omitempty"`
	AccountType constants.AccType    `json:"accountType,omitempty"`
	CreatedAt   string               `json:"createdAt,omitempty"`
	UpdatedAt   string               `json:"updatedAt,omitempty"`
}
