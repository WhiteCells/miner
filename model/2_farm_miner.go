package model

type FarmMiner struct {
	FarmID  int `gorm:"not null;index" json:"farm_id"`
	MinerID int `gorm:"not null;index" json:"miner_id"`
}

func (FarmMiner) TableName() string {
	return "farm_miner"
}
