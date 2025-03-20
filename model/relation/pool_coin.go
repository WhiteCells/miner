package relation

type PoolCoin struct {
	PoolID int `gorm:"index"`
	CoinID int `gorm:"index"`
}

func (PoolCoin) TableName() string {
	return "pool_coin"
}
