package model

type UserWallet struct {
	UserID   int `gorm:"not null;index"`
	WalletID int `gorm:"not null;index"`
}

func (UserWallet) TableName() string {
	return "user_wallet"
}
