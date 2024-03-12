package entity

import "time"

type AccType string

const (
	Public  AccType = "PUBLIC"
	Private AccType = "PRIVATE"
)

type User struct {
	ID            int64     `json:"id,omitempty"`
	Username      string    `json:"username,omitempty"`
	Email         string    `json:"email,omitempty"`
	Password      string    `json:"password,omitempty"`
	FullName      string    `json:"fullName,omitempty"`
	Status        string    `json:"status,omitempty"`
	Bio           string    `json:"bio,omitempty"`
	ProfilePicUrl string    `json:"profilePicUrl,omitempty"`
	AccountType   AccType   `json:"accountType,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
}

type Profile struct {
	Followers []*User `json:"followers,omitempty"`
	Following []*User `json:"following,omitempty"`
}
