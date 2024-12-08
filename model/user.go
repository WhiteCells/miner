package model

import (
	"miner/common/role"
	"miner/common/status"
	"time"
)

type User struct {
	ID          int               `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:用户ID"`
	Name        string            `json:"name" gorm:"unique;not null;column:name;type:varchar(255);comment:用户名"`
	Password    string            `json:"password" gorm:"not null;column:password;type:varchar(255);comment:用户密码"`
	Secret      string            `json:"secret" gorm:"column:secret;type:varchar(64);comment:用户登陆密钥"`
	Email       string            `json:"email" gorm:"unique;not null;column:email;type:varchar(255);comment:用户邮箱"`
	Role        role.RoleType     `json:"role" gorm:"column:role;type:varchar(255);comment:用户角色"`
	Points      int               `json:"points" gorm:"column:points;type:int;comment:用户积分"`
	Status      status.UserStatus `json:"status" gorm:"column:status;type:int;comment:用户状态"`
	LastLoginAt time.Time         `json:"last_login_at" gorm:"column:last_login_at;type:datetime;comment:用户登陆时间"`
	LastLoginIP string            `json:"last_login_ip" gorm:"column:last_login_ip;type:varchar(255);comment:用户登陆IP"`
	InviteCode  string            `json:"invite_code" gorm:"unique;column:invite_code;type:varchar(255);comment:用户邀请码"`
	InvitedBy   int               `json:"invited_by" gorm:"column:invited_by;type:int;comment:被邀请用户ID"`
}

func (User) TableName() string {
	return "user"
}
