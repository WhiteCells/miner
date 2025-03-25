package dto

type CreateFarmReq struct {
	Name     string `json:"name" binding:"required,min=1,max=40"`
	TimeZone string `json:"time_zone" binding:"required,min=1,max=40"`
}

type DeleteFarmReq struct {
	FarmID int `json:"farm_id" binding:"required"`
}

type UpdateFarmReq struct {
	FarmID     int            `json:"farm_id" binding:"required"`
	UpdateInfo map[string]any `json:"update_info" binding:"required"`
}

type ApplyFarmFsReq struct {
	FarmID int `json:"farm_id" binding:"required"`
	FsID   int `json:"fs_id" binding:"required"`
}

type UnApplyFarmFsReq struct {
	FarmID int `json:"farm_id" binding:"required"`
	FsID   int `json:"fs_id" binding:"required"`
}

type TransferFarmReq struct {
	FarmID   int `json:"farm_id" binding:"required"`
	ToUserID int `json:"to_user_id" binding:"required"`
}
