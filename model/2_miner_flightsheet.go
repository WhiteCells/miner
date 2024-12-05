package model

type MinerFlightsheet struct {
	MinerID       int `gorm:"not null;index" json:"miner_id"`
	FlightsheetID int `gorm:"not null;index" json:"flightsheet_id"`
}

func (MinerFlightsheet) TableName() string {
	return "miner_flightsheet"
}
