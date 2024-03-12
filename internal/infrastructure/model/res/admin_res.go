package res

import "github.com/abdullahnettoor/connect-social-media/internal/domain/entity"

type AdminLoginRes struct {
	Code    int          `json:"status,omitempty"`
	Message string       `json:"message,omitempty"`
	Token   string       `json:"token,omitempty"`
	Error   error        `json:"error,omitempty"`
	Admin   entity.Admin `json:"admin,omitempty"`
}
