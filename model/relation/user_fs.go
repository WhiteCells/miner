package relation

import "miner/common/perm"

type UserFs struct {
	UserID int         `gorm:"index"`
	FsID   int         `gorm:"index"`
	Perm   perm.FsPerm `gorm:"column:perm;type:varchar(16)"`
}

func (UserFs) TableName() string {
	return "user_fs"
}
