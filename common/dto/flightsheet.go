package dto

type CreateFsReq struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Coin     string `json:"coin" binding:"required,min=2,max=20"`
	WalletID string `json:"wallet_id" binding:"required"`
	Pool     string `json:"pool" binding:"required,min=2,max=20"`
	Soft     string `json:"soft" binding:"required,min=2,max=20"`
}

type DeleteFsReq struct {
	FsID string `json:"fs_id" binding:"required"`
}

type UpdateFsReq struct {
	FsID       string                 `json:"fs_id" binding:"required"`
	UpdateInfo map[string]interface{} `json:"update_info" binding:"required"`
}

type ApplyWalletReq struct {
	FsID     string `json:"fs_id" binding:"required"`
	WaleltID string `json:"walelt_id" binding:"required"`
}
