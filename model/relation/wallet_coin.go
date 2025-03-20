package relation

type WalletCoin struct {
	WalletID int `json:"wallet_id" gorm:"column:wallet_id;type:int"`
	CoinID   int `json:"coin_id" gorm:"column:coin_id;type:int"`
}

func (WalletCoin) TableName() string {
	return "wallet_coin"
}
