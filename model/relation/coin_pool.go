package relation

type CoinPool struct {
	CoinID int `gorm:"index"`
	PoolID int `gorm:"index"`
}

func (CoinPool) TableName() string {
	return "coin_pool"
}
