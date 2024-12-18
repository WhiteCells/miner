package dto

import "miner/common/status"

type AdminSwitchRegisterReq struct {
	Status status.RegisterStatus `json:"status"`
}

type AdminSetGlobalFlightsheetReq struct {
	Name      string `json:"name" binding:"required"`
	CoinType  string `json:"coin_type" binding:"required"`
	WalletID  string `json:"wallet_id" binding:"required"`
	MinerPool string `json:"miner_pool" binding:"required"`
	MineSoft  string `json:"mine_soft" binding:"required"`
}

type AdminSetInviteRewardReq struct {
	Reward int `json:"reward" binding:"required"`
}

type AdminSetRechargeRewardReq struct {
	Reward int `json:"reward" binding:"required"`
}

type AdminSetUserStatusReq struct {
	UserID int               `json:"user_id" binding:"required"`
	Status status.UserStatus `json:"status" binding:"required"`
}

type AdminSetMinerPoolCostReq struct {
	MinePoolID int     `json:"miner_pool_id" binding:"required"`
	Cost       float64 `json:"cost" binding:"required"`
}
