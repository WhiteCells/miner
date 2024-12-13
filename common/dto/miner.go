package dto

type CreateMinerReq struct {
	FarmID      int    `json:"farm_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Model       string `json:"model" binding:"required"`
	IP          string `json:"ip" binding:"required,ip"`
	SSHPort     int    `json:"ssh_port" binding:"required"`
	SSHUser     string `json:"ssh_user" binding:"required"`
	SSHPassword string `json:"ssh_password" binding:"required"`
}

type DeleteMinerReq struct {
	FarmID  int `json:"farm_id" binding:"required"`
	MinerID int `json:"miner_id" binding:"required"`
}

type UpdateMinerReq struct {
	FarmID     int                    `json:"farm_id" binding:"required"`
	MinerID    int                    `json:"miner_id" binding:"required"`
	UpdateInfo map[string]interface{} `json:"update_info" binding:"required"`
}

type ApplyMinerFlightsheetReq struct {
	MinerID       int `json:"miner_id" binding:"required"`
	FlightsheetID int `json:"fs_id" binding:"required"`
}

type TransferMinerReq struct {
	FarmID   int    `json:"farm_id" binding:"required"`
	MinerID  int    `json:"miner_id" binding:"required"`
	FarmHash string `json:"farm_hash" binding:"required"`
}
