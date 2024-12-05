package model

import (
	"miner/common/role"
	"miner/common/status"
	"time"
)

type User struct {
	ID          int               `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name        string            `json:"name" gorm:"unique;not null;column:name;type:varchar(255)"`
	Password    string            `json:"password" gorm:"not null;column:password;type:varchar(255)"`
	Email       string            `json:"email" gorm:"not null;column:email;type:varchar(255)"`
	Role        role.RoleType     `json:"role" gorm:"column:role;type:varchar(255)"`
	Points      int               `json:"points" gorm:"column:points;type:int"`
	Status      status.UserStatus `json:"status" gorm:"column:status;type:int"`
	LastLoginAt time.Time         `json:"last_login_at" gorm:"column:last_login_at;type:datetime"`
	LastLoginIP string            `json:"last_login_ip" gorm:"column:last_login_ip;type:varchar(255)"`
	InviteCode  string            `json:"invite_code" gorm:"unique;column:invite_code;type:varchar(255)"`
	InvitedBy   int               `json:"invited_by" gorm:"column:invited_by;type:int"`
}

func (User) TableName() string {
	return "user"
}
