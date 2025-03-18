package relation

type FarmFs struct {
	FarmID int `gorm:"index"`
	FsID   int `gorm:"index"`
}

func (FarmFs) TableName() string {
	return "farm_fs"
}
