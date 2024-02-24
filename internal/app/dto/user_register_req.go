package dto

type UserRegisterReq struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
