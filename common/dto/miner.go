package dto

type CreateMinerReq struct {
	UserID      int    `json:"user_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Model       string `json:"model" binding:"required"`
	IPAddress   string `json:"ip_address" binding:"required,ip"`
	SSHPort     int    `json:"ssh_port" binding:"required"`
	SSHUser     string `json:"ssh_user" binding:"required"`
	SSHPassword string `json:"ssh_password" binding:"required"`
	FarmID      int    `json:"farm_id" binding:"required"`
}

type TransferMinerReq struct {
	FromUserID int `json:"from_user_id" binding:"required"`
	ToUserID   int `json:"to_user_id" binding:"required"`
	FarmID     int `json:"farm_id" binding:"required"`
}
