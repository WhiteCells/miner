package model

import "miner/common/perm"

type UserFarm struct {
	UserID int           `gorm:"not null;index" json:"user_id"`
	FarmID int           `gorm:"not null;index" json:"farm_id"`
	Role   perm.FarmPerm `gorm:"not null;default:owner" json:"role"`
}

func (UserFarm) TableName() string {
	return "user_farm"
}
