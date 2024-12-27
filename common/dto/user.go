package dto

import (
	"miner/common/points"
	"time"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	// todo 图形验证码
	// google 验证码
	// GoogleCode string `json:"google_code" binding:"required"`
}

type RegisterReq struct {
	Username   string `json:"username" binding:"required,min=3,max=32"`
	Password   string `json:"password" binding:"required,min=6,max=32"`
	Email      string `json:"email" binding:"required,email"`
	InviteCode string `json:"invite_code"`
}

type UpdateInfoReq struct {
	UpdateInfo map[string]interface{}
}

type AddPointsReq struct {
	UserID string            `json:"user_id" binding:"required"`
	Type   points.PointsType `json:"type" binding:"required"`
	Point  int               `json:"point" binding:"required"`
	Time   time.Time         `json:"time" binding:"required"`
}
