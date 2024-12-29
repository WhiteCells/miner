package dto

type CreateFsReq struct {
	Name     string `json:"name" binding:"required"`
	CoinID   string `json:"coin" binding:"required"`
	WalletID string `json:"wallet_id" binding:"required"`
	MineID   string `json:"mine_id" binding:"required"`
	SoftID   string `json:"soft_id" binding:"required"`
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
