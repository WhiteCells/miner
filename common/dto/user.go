package dto

type LoginReq struct {
	Username string `json:"username" binding:"required,min=3,max=32"`
	Password string `json:"password" binding:"required,min=6,max=32"`
	// todo 图形验证码
	// google 验证码
	// GoogleCode string `json:"google_code" binding:"required"`
}

type RegisterReq struct {
	Username   string `json:"username" binding:"required,min=3,max=32"`
	Password   string `json:"password" binding:"required,min=6,max=32"`
	Email      string `json:"email" binding:"required,email,max=32"`
	InviteCode string `json:"invite_code"`
}
