package model

type UserFlightsheet struct {
	UserID        int `json:"user_id" gorm:"not null;index;column:user_id;type:int"`
	FlightsheetID int `json:"flightsheet_id" gorm:"not null;index;column:flightsheet_id;type:int"`
}

func (UserFlightsheet) TableName() string {
	return "user_flightsheet"
}
