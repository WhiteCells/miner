package info

import (
	"miner/common/role"
	"miner/common/status"
	"time"
)

type User struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Email       string            `json:"email"`
	Password    string            `json:"password"`
	Secret      string            `json:"secret"`
	Role        role.RoleType     `json:"role"`
	Points      int               `json:"points"`
	Status      status.UserStatus `json:"status"`
	LastLoginAt time.Time         `json:"last_login_at"`
	LastLoginIP string            `json:"last_login_ip"`
	InviteCode  string            `json:"invite_code"`
	InviteBy    int               `json:"invite_by"`
}
