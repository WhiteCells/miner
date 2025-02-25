package dto

import "miner/model/info"

type LoginReq struct {
	Email    string `json:"email" binding:"required,min=3,max=32"`
	Password string `json:"password" binding:"required,min=6,max=32"`
	// CaptchaID    string `json:"captcha_id" binding:"required"`
	// CaptchaValue string `json:"captcha_value" binding:"required"`
	// google 验证码
	// GoogleCode string `json:"google_code" binding:"required"`
}

type RegisterReq struct {
	Username   string `json:"username" binding:"required,min=3,max=32"`
	Password   string `json:"password" binding:"required,min=6,max=32"`
	Email      string `json:"email" binding:"required,email,max=32"`
	InviteCode string `json:"invite_code"`
}

type GenerateCaptchaRsp struct {
	CaptchaID string `json:"captcha_id"`
	Base64    string `json:"base64"`
}

type VerifyCaptchaReq struct {
	CaptchaID string `json:"captcha_id"`
	Value     string `json:"value"`
}

type UpdatePasswdReq struct {
	OldPasswd string `json:"old_passwd" binding:"required,min=6,max=32"`
	NewPasswd string `json:"new_passwd" binding:"required,min=6,max=32"`
}

type ApplySoftReq struct {
	FsID string    `json:"fs_id" binding:"required"`
	Soft info.Soft `json:"soft" binding:"required"`
}
