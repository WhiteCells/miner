package model

type Fssub struct {
	ID       int `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	CoinID   int `gorm:"column:coin_id;type:int"`
	WalletID int `gorm:"column:wallet_id;type:int"`
	PoolID   int `gorm:"column:pool_id;type:int"`
	SoftID   int `gorm:"column:soft_id;type:int"`
}

func (Fssub) TableName() string {
	return "fssub"
}

func GetFssubAllowChangeField() map[string]bool {
	return map[string]bool{
		"coin_id":   true,
		"wallet_id": true,
		"pool_id":   true,
		"soft_id":   true,
	}
}
