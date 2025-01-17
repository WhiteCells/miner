package dto

type CreateFarmReq struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	TimeZone string `json:"time_zone" binding:"required,min=3,max=20"`
}

type DeleteFarmReq struct {
	FarmID string `json:"farm_id" binding:"required,min=3,max=20"`
}

type UpdateFarmReq struct {
	FarmID     string                 `json:"farm_id" binding:"required,max=19"`
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
