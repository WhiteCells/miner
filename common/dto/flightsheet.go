package dto

type CreateFlightsheetReq struct {
	Name     string `json:"name" binding:"required"`
	CoinType string `json:"coin_type" binding:"required"`
	WalletID int    `json:"wallet_id" binding:"required"`
	MinePool string `json:"mine_pool" binding:"required"`
	MineSoft string `json:"mine_soft" binding:"required"`
}

type DeleteFlightsheetReq struct {
	FlightsheetID int `json:"flightsheet_id" binding:"required"`
}

type UpdateFlightsheetReq struct {
	FlightsheetID int                    `json:"flightsheet_id" binding:"required"`
	UpdateInfo    map[string]interface{} `json:"update_info" binding:"required"`
}

type GetUserAllFlightsheetReq struct {
}
