package model

import "miner/common/perm"

type UserMiner struct {
	UserID  int            `gorm:"not null;index"`
	MinerID int            `gorm:"not null;index"`
	Perm    perm.MinerPerm `gorm:"not null;index"`
}

func (UserMiner) TableName() string {
	return "user_miner"
}
