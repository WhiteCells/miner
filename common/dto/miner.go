package dto

type CreateMinerReq struct {
	FarmID string `json:"farm_id" binding:"required,min=3,max=20"`
	Name   string `json:"name" binding:"required,min=3,max=20"`
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
	FarmID        string `json:"farm_id" binding:"required"`
	MinerID       string `json:"miner_id" binding:"required"`
	FlightsheetID string `json:"fs_id" binding:"required"`
}

type TransferMinerReq struct {
	FromFarmID string `json:"from_farm_id" binding:"required"`
	MinerID    string `json:"from_miner_id" binding:"required"`
	ToUserID   string `json:"to_user_id" binding:"required"`
	ToFarmID   string `json:"to_farm_id" binding:"required"`
}

type GetRigConfReq struct {
	FarmID  string `json:"farm_id" binding:"required"`
	MinerID string `json:"miner_id" binding:"required"`
}
