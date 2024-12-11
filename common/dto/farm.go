package dto

type CreateFarmReq struct {
	Name     string `json:"name" binding:"required"`
	TimeZone string `json:"time_zone" binding:"required"`
}

type DeleteFarmReq struct {
	FarmID int `json:"farm_id" binding:"required"`
}

type UpdateFarmReq struct {
	FarmID     int                    `json:"farm_id" binding:"required"`
	UpdateInfo map[string]interface{} `json:"update_info" binding:"required"`
}

type GetUserAllMinerInFarmReq struct {
	FarmID int `json:"farm_id" binding:"required"`
}

type ApplyFarmFlightsheetReq struct {
	FarmID        int `json:"farm_id" binding:"required"`
	FlightsheetID int `json:"flightsheet_id" binding:"required"`
}
