package dto

type CreateWalletReq struct {
	Name string `json:"name" binding:"required,min=1,max=20"`
	Addr string `json:"addr" binding:"required,min=1,max=256"`
	Coin string `json:"coin" binding:"required,min=1,max=20"`
}

type DeleteWalletReq struct {
	WalletID string `json:"wallet_id" binding:"required"`
}

type UpdateWalletReq struct {
	WalletID   string         `json:"wallet_id" binding:"required"`
	UpdateInfo map[string]any `json:"update_info" binding:"required"`
}

type GetUserWalletByIDReq struct {
	WalletID string `json:"wallet_id" binding:"required"`
}
