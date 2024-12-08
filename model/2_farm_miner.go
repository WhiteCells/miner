package model

type FarmMiner struct {
	FarmID  int `gorm:"not null;index;comment:矿场ID" json:"farm_id"`
	MinerID int `gorm:"not null;index;comment:矿机ID" json:"miner_id"`
}

func (FarmMiner) TableName() string {
	return "farm_miner"
}
