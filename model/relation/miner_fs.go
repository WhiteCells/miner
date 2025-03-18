package relation

type MinerFs struct {
	MinerID int `gorm:"index"`
	FsID    int `gorm:"index"`
}

func (MinerFs) TableName() string {
	return "miner_fs"
}
