package req

type AdminLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
