package model

import "miner/common/perm"

type UserMiner struct {
	UserID  int            `gorm:"not null;index;comment:用户ID"`
	MinerID int            `gorm:"not null;index;comment:矿机ID"`
	Perm    perm.MinerPerm `gorm:"not null;index;comment:权限"`
}

func (UserMiner) TableName() string {
	return "user_miner"
}
