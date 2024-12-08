package model

type FlightsheetWallet struct {
	FlightsheetID int `gorm:"not null;index;comment:飞行表ID" json:"flightsheet_id"`
	WalletID      int `gorm:"not null;index;comment:钱包ID" json:"wallet_id"`
}

func (FlightsheetWallet) TableName() string {
	return "flightsheet_wallet"
}
