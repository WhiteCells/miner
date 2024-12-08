package model

type UserWallet struct {
	UserID   int `gorm:"not null;index;comment:用户ID"`
	WalletID int `gorm:"not null;index;comment:钱包ID"`
}

func (UserWallet) TableName() string {
	return "user_wallet"
}
