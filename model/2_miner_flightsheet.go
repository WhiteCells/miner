package model

type MinerFlightsheet struct {
	MinerID       int `gorm:"not null;index;comment:矿机ID" json:"miner_id"`
	FlightsheetID int `gorm:"not null;index;comment:飞行表ID" json:"flightsheet_id"`
}

func (MinerFlightsheet) TableName() string {
	return "miner_flightsheet"
}
