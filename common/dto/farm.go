package dto

type CreateFarmReq struct {
	// UserID   int    `json:"user_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	TimeZone string `json:"time_zone" binding:"required"`
}

type GetFarmInfoReq struct {
	FarmID int `json:"farm_id" binding:"required"`
}
