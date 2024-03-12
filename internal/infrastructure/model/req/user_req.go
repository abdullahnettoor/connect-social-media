package req

type SignUpReq struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserId struct {
	UserID int64 `json:"userId"`
}
