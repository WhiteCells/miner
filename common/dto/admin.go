package dto

import (
	"miner/common/status"
	"miner/model/info"
)

type AdminSwitchRegisterReq struct {
	Status status.RegisterStatus `json:"status" binding:"required,oneof=1 0"`
}

type AdminSetGlobalFsReq struct {
	Name     string    `json:"name" binding:"required,min=1,max=20"`
	Coin     info.Coin `json:"coin" binding:"required"`
	WalletID string    `json:"wallet_id" binding:"required"`
	Pool     info.Pool `json:"pool" binding:"required"`
	Soft     info.Soft `json:"soft" binding:"required"`
}

type AdminSetInviteRewardReq struct {
	Reward int `json:"reward" binding:"required,gt=0"`
}

type AdminSetRechargeRewardReq struct {
	Ratio float64 `json:"ratio" binding:"required"`
}

type AdminSetUserStatusReq struct {
	UserID string            `json:"user_id" binding:"required"`
	Status status.UserStatus `json:"status" binding:"required,oneof=1 0"`
}

type AdminSetMinePoolCostReq struct {
	MinepoolID string  `json:"minerpool_id" binding:"required"`
	Cost       float64 `json:"cost" binding:"required,gt=0"`
}

type AdminSetMnemonicReq struct {
	Mnemonic string `json:"mnemonic" binding:"required"`
}

type AdminAddBscApiKeyReq struct {
	Apikey string `json:"apikey" binding:"required"`
}

type AdminDelBscApiKeyReq struct {
	Apikey string `json:"apikey" binding:"required"`
}

type AdminAddCoinReq struct {
	Coin info.Coin `json:"coin" binding:"required"`
}

type AdminDelCoinReq struct {
	CoinName string `json:"coin_name" binding:"required"`
}

type AdminAddPoolReq struct {
	CoinName string    `json:"coin_name" binding:"required"`
	Pool     info.Pool `json:"pool" binding:"required"`
}

type AdminDelPoolReq struct {
	CoinName string `json:"coin_name" binding:"required"`
	PoolName string `json:"pool_name" binding:"required"`
}

type AdminAddSoftReq struct {
	CoinName string    `json:"coinName" binding:"required"`
	Soft     info.Soft `json:"soft" binding:"required"`
}

type AdminDelSoftReq struct {
	CoinName string `json:"coinName" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type AdminSetFreeGpuNumReq struct {
	GpuNum int `json:"gpu_num" binding:"required"`
}
