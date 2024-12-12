package model

type FarmFlightsheet struct {
	FarmID        int `gorm:"not null;index;column:farm_id;comment:矿场ID" json:"farm_id"`
	FlightsheetID int `gorm:"not null;index;comment:flightsheet_id;comment:飞行表ID" json:"flightsheet_id"`
}
