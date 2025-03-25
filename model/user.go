package model

import (
	"miner/common/role"
	"miner/common/status"
	"time"
)

type User struct {
	ID             int               `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:用户ID"`
	Name           string            `json:"name" gorm:"unique;not null;column:name;type:varchar(255);comment:用户名"`
	Email          string            `json:"email" gorm:"unique;column:email;type:varchar(255);comment:用户邮箱"`
	Password       string            `json:"password" gorm:"not null;column:password;type:varchar(255);comment:用户密码"`
	Secret         string            `json:"secret" gorm:"column:secret;type:varchar(64);comment:用户登陆密钥"`
	Address        string            `json:"address" gorm:"column:address;type:varchar(255);comment:充值地址"`
	Role           role.RoleType     `json:"role" gorm:"column:role;type:varchar(16);comment:用户角色"`
	InvitePoints   float32           `json:"invite_points" gorm:"column:invite_points;type:float;comment:邀请积分"`
	RechargePoints float32           `json:"recharge_points" gorm:"column:recharge_points;type:float;comment:充值积分"`
	LastBalance    float32           `json:"last_balance" gorm:"column:last_balance;type:float;comment:上次余额"`
	Status         status.UserStatus `json:"status" gorm:"column:status;type:varchar(16);comment:用户状态"`
	UID            string            `json:"uid" gorm:"column:uid;type:varchar(255);comment:用户MN路径"`
	LastCheckAt    time.Time         `json:"last_check_at" gorm:"column:last_check_at;type:datetime;comment:上次支付时间"`
	LastLoginAt    time.Time         `json:"last_login_at" gorm:"column:last_login_at;type:datetime;comment:用户登陆时间"`
	LastLoginIP    string            `json:"last_login_ip" gorm:"column:last_login_ip;type:varchar(255);comment:用户登陆IP"`
	InviteCode     string            `json:"invite_code" gorm:"unique;column:invite_code;type:varchar(255);comment:用户邀请码"`
	InvitedBy      string            `json:"invited_by" gorm:"column:invited_by;type:varchar(255);comment:被邀请用户ID"`
	// todo key 不能返回给用户
	Key string `json:"key" gorm:"column:key;type:varchar(255);comment:钱包密钥"`
}

func (User) TableName() string {
	return "user"
}
