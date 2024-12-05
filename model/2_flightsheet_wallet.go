package model

type FlightsheetWallet struct {
	FlightsheetID int `gorm:"not null;index" json:"flightsheet_id"`
	WalletID      int `gorm:"not null;index" json:"wallet_id"`
}

func (FlightsheetWallet) TableName() string {
	return "flightsheet_wallet"
}
