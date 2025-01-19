package dto

type CreateWalletReq struct {
	Name string `json:"name" binding:"required,min=3,max=20"`
	Addr string `json:"addr" binding:"required,len=42"`
	Coin string `json:"coin" binding:"required,min=2,max=20"`
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
