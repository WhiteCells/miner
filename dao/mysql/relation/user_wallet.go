package relation

type UserWallet struct {
	UserID   int `gorm:"index"`
	WalletID int `gorm:"index"`
}

func (UserWallet) TableName() string {
	return "user_wallet"
}
