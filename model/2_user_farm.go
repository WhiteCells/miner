package model

import "miner/common/perm"

type UserFarm struct {
	UserID int           `gorm:"not null;index;comment:用户ID" json:"user_id"`
	FarmID int           `gorm:"not null;index;comment:矿场ID" json:"farm_id"`
	Perm   perm.FarmPerm `gorm:"not null;default:owner;comment:权限" json:"perm"`
}

func (UserFarm) TableName() string {
	return "user_farm"
}
