package relation

import "miner/common/perm"

type UserFarm struct {
	UserID int           `gorm:"index"`
	FarmID int           `gorm:"index"`
	Perm   perm.FarmPerm `gorm:"column:perm;type:varchar(255)"`
}

func (UserFarm) TableName() string {
	return "user_farm"
}
