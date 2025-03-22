package dto

type CreateFarmReq struct {
	Name     string `json:"name" binding:"required,min=1,max=20"`
	TimeZone string `json:"time_zone" binding:"required,min=1,max=20"`
}

type DeleteFarmReq struct {
	FarmID int `json:"farm_id" binding:"required,min=1,max=20"`
}

type UpdateFarmReq struct {
	FarmID     int            `json:"farm_id" binding:"required"`
	UpdateInfo map[string]any `json:"update_info" binding:"required"`
}

type UpdateFarmHashReq struct {
	FarmID string `json:"farm_id" binding:"required"`
	Hash   string `json:"hash" binding:"required,len=40"`
}

type ApplyFarmFsReq struct {
	FarmID string `json:"farm_id" binding:"required"`
	FsID   string `json:"fs_id" binding:"required"`
}

type TransferFarmReq struct {
	FarmID   int `json:"farm_id" binding:"required"`
	ToUserID int `json:"to_user_id" binding:"required"`
}
