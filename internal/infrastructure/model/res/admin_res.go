package res

import "github.com/abdullahnettoor/connect-social-media/internal/domain/entity"

type AdminLoginRes struct {
	CommonRes
	Token   string       `json:"token,omitempty"`
	Admin   entity.Admin `json:"admin,omitempty"`
}
