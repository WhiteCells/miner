package dto

type CreateMinerReq struct {
	FarmID string `json:"farm_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Model  string `json:"model" binding:"required"`
	IP     string `json:"ip" binding:"required,ip"`
}

type DeleteMinerReq struct {
	FarmID  string `json:"farm_id" binding:"required"`
	MinerID string `json:"miner_id" binding:"required"`
}

type UpdateMinerReq struct {
	FarmID     string                 `json:"farm_id" binding:"required"`
	MinerID    string                 `json:"miner_id" binding:"required"`
	UpdateInfo map[string]interface{} `json:"update_info" binding:"required"`
}

type ApplyMinerFlightsheetReq struct {
	MinerID       string `json:"miner_id" binding:"required"`
	FlightsheetID string `json:"fs_id" binding:"required"`
}

type TransferMinerReq struct {
	FromFarmID string `json:"from_farm_id" binding:"required"`
	MinerID    string `json:"from_miner_id" binding:"required"`
	ToUserID   string `json:"to_user_id" binding:"required"`
	ToFarmID   string `json:"to_farm_id" binding:"required"`
}
