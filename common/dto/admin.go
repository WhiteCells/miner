package dto

import (
	"miner/common/status"
	"miner/model/info"
)

type AdminSwitchRegisterReq struct {
	Status status.RegisterStatus `json:"status" binding:"required,oneof=1 0"`
}

type AdminSetGlobalFsReq struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Coin     string `json:"coin" binding:"required,min=3,max=20"`
	WalletID string `json:"wallet_id" binding:"required"`
	Pool     string `json:"pool" binding:"required,min=2,max=20"`
	Soft     string `json:"soft" binding:"required,min=2,max=20"`
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

// test
type AdminIncrBscApiKeyReq struct {
	Apikey string  `json:"apikey" binding:"required"`
	Score  float64 `json:"score" binding:"required"`
}

// coin
/*
{
	"coin": {
		"name": ""
	}
}
*/
type AdminAddCoinReq struct {
	Coin info.Coin `json:"coin" binding:"required,min=2,max=20"`
}

type AdminDelCoinReq struct {
	Name string `json:"name" binding:"required"`
}

// pool
/*
{
	"pool": {
		"name": "",
		"server": []
	}
}
*/
type AdminAddPoolReq struct {
	Pool info.Pool `json:"pool" binding:"required,min=2,max=20"`
}

type AdminDelPoolReq struct {
	Name string `json:"name" binding:"required"`
}

// soft
/*
{
	"soft": {
		"name": "",
		...
	}
}
*/
type AdminAddSoftReq struct {
	Soft info.Soft `json:"soft" binding:"required,min=2,max=20"`
}

type AdminDelSoftReq struct {
	Name string `json:"name" binding:"required"`
}

type AdminSetFreeGpuNumReq struct {
	GpuNum int `json:"gpu_num" binding:"required"`
}
