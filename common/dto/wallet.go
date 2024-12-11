package dto

type CreateWalletReq struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	CoinType string `json:"coin_type"`
}

type DeleteWalletReq struct {
	WalletID int `json:"wallet_id"`
}

type UpdateWalletReq struct {
	WalletID   int                    `json:"wallet_id"`
	UpdateInfo map[string]interface{} `json:"update_info"`
}

// type GetUserAllWalletReq struct {
// }

type GetUserWalletByIDReq struct {
	WalletID int `json:"wallet_id"`
}
