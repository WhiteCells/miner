package relation

type UserMiner struct {
	UserID  int `gorm:"index"`
	MinerID int `gorm:"index"`
}

func (UserMiner) TableName() string {
	return "user_miner"
}
