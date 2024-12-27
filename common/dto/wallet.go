package dto

type CreateWalletReq struct {
	Name string `json:"name" binding:"required"`
	Addr string `json:"address" binding:"required"`
	Coin string `json:"coin_type" binding:"required"`
}

type DeleteWalletReq struct {
	WalletID string `json:"wallet_id" binding:"required"`
}

type UpdateWalletReq struct {
	WalletID   string                 `json:"wallet_id" binding:"required"`
	UpdateInfo map[string]interface{} `json:"update_info" binding:"required"`
}

type GetUserWalletByIDReq struct {
	WalletID string `json:"wallet_id" binding:"required"`
}
