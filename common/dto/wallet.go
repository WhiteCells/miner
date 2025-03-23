package dto

type CreateWalletReq struct {
	Name   string `json:"name" binding:"required,min=1,max=20"`
	Addr   string `json:"addr" binding:"required,min=1,max=256"`
	CoinID int    `json:"coin_id" binding:"required"`
}

type DeleteWalletReq struct {
	WalletID int `json:"wallet_id" binding:"required"`
}

type UpdateWalletReq struct {
	WalletID   int            `json:"wallet_id" binding:"required"`
	UpdateInfo map[string]any `json:"update_info" binding:"required"`
}

type GetUserWalletByIDReq struct {
	WalletID int `json:"wallet_id" binding:"required"`
}
