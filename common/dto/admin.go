package dto

import (
	"miner/common/role"
	"time"
)

type AdmineGetUserReq struct {
	Role      role.RoleType `json:"role"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
	PageNum   int           `json:"page_num"`
	PageSize  int           `json:"page_size"`
}
