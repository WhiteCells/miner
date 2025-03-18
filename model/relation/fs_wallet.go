package relation

type FsWallet struct {
	FsID     int `gorm:"index"`
	WalletID int `gorm:"index"`
}

func (FsWallet) TableName() string {
	return "fs_wallet"
}
