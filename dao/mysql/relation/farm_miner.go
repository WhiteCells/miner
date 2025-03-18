package relation

type FarmMiner struct {
	FarmID  int `gorm:"index"`
	MinerID int `gorm:"index"`
}

func (FarmMiner) TableName() string {
	return "farm_miner"
}
