package relation

type FarmMiner struct {
	FarmID  int    `gorm:"index"`
	MinerID int    `gorm:"index"`
	Perm    string `gorm:"column:perm;type:varchar(16)"`
}

func (FarmMiner) TableName() string {
	return "farm_miner"
}
