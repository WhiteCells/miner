package dto

import "miner/common/status"

type AdminSwitchRegisterReq struct {
	Status status.RegisterStatus `json:"status"`
}

type AdminSetGlobalFsReq struct {
	Name     string `json:"name" binding:"required"`
	Coin     string `json:"coin" binding:"required"`
	WalletID string `json:"wallet_id" binding:"required"`
	Mine     string `json:"miner" binding:"required"`
	Soft     string `json:"soft" binding:"required"`
}

type AdminSetInviteRewardReq struct {
	Reward int `json:"reward" binding:"required"`
}

type AdminSetRechargeRewardReq struct {
	Reward int `json:"reward" binding:"required"`
}

type AdminSetUserStatusReq struct {
	UserID string            `json:"user_id" binding:"required"`
	Status status.UserStatus `json:"status" binding:"required"`
}

type AdminSetMinePoolCostReq struct {
	MinepoolID string  `json:"minerpool_id" binding:"required"`
	Cost       float64 `json:"cost" binding:"required"`
}
