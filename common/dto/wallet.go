package dto

type CreateWalletReq struct {
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	CoinType string `json:"coin_type" binding:"required"`
}

type DeleteWalletReq struct {
	WalletID int `json:"wallet_id" binding:"required"`
}

type UpdateWalletReq struct {
	WalletID   int                    `json:"wallet_id" binding:"required"`
	UpdateInfo map[string]interface{} `json:"update_info" binding:"required"`
}

type GetUserWalletByIDReq struct {
	WalletID int `json:"wallet_id" binding:"required"`
}
