package dto

type CreateFarmReq struct {
	Name     string `json:"name" binding:"required"`
	TimeZone string `json:"time_zone" binding:"required"`
}

type DeleteFarmReq struct {
	FarmID string `json:"farm_id" binding:"required"`
}

type UpdateFarmReq struct {
	FarmID     string                 `json:"farm_id" binding:"required"`
	UpdateInfo map[string]interface{} `json:"update_info" binding:"required"`
}

type ApplyFarmFlightsheetReq struct {
	FarmID        string `json:"farm_id" binding:"required"`
	FlightsheetID string `json:"fs_id" binding:"required"`
}

type TransferFarmReq struct {
	FarmID   string `json:"farm_id" binding:"required"`
	ToUserID string `json:"to_user_id" binding:"required"`
}
