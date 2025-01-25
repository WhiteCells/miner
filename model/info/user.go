package info

import (
	"miner/common/role"
	"miner/common/status"
	"time"
)

type User struct {
	ID             string            `json:"id"`
	Name           string            `json:"name"`
	Email          string            `json:"email"`
	Password       string            `json:"password"`
	Secret         string            `json:"secret"`
	MnemonicVer    string            `json:"mnemonic_ver"` // 助记词版本
	Address        string            `json:"address"`
	Role           role.RoleType     `json:"role"`
	InvitePoints   float32           `json:"invite_points"`
	RechargePoints float32           `json:"recharge_points"`
	LastBalance    float32           `json:"last_balance"`
	Status         status.UserStatus `json:"status"`
	LastLoginAt    time.Time         `json:"last_login_at"`
	LastLoginIP    string            `json:"last_login_ip"`
	LastCheckAt    time.Time         `json:"last_check_at"`
	InviteCode     string            `json:"invite_code"`
	InviteBy       string            `json:"invite_by"`

	// test
	Key string `json:"key"`
}
