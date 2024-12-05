package dto

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	// todo 图形验证码
	// google 验证码
	GoogleCode string `json:"google_code" binding:"required"`
}

type RegisterReq struct {
	Username   string `json:"username" binding:"required,min=3,max=32"`
	Password   string `json:"password" binding:"required,min=6,max=32"`
	Email      string `json:"email" binding:"required,email"`
	InviteCode string `json:"invite_code"`
}

type UpdateInfoReq struct {
	UserID     int `json:"user_id" binding:"required"`
	UpdateInfo map[string]interface{}
}

type AddPointsReq struct {
	UserID int    `json:"user_id" binding:"required"`
	Type   string `json:"type" binding:"required"`
	Point  int    `json:"point" binding:"required"`
}

type GetUserInfoReq struct {
	UserID int `json:"user_id" binding:"required"`
}
