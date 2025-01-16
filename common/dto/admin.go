package dto

import (
	"miner/common/status"
	"miner/model/info"
)

type AdminSwitchRegisterReq struct {
	Status status.RegisterStatus `json:"status"`
}

type AdminSetGlobalFsReq struct {
	Name     string `json:"name" binding:"required"`
	Coin     string `json:"coin" binding:"required"`
	WalletID string `json:"wallet_id" binding:"required"`
	Pool     string `json:"pool" binding:"required"`
	Soft     string `json:"soft" binding:"required"`
}

type AdminSetInviteRewardReq struct {
	Reward int `json:"reward" binding:"required"`
}

type AdminSetRechargeRewardReq struct {
	Ratio float64 `json:"ratio" binding:"required"`
}

type AdminSetUserStatusReq struct {
	UserID string            `json:"user_id" binding:"required"`
	Status status.UserStatus `json:"status" binding:"required"`
}

type AdminSetMinePoolCostReq struct {
	MinepoolID string  `json:"minerpool_id" binding:"required"`
	Cost       float64 `json:"cost" binding:"required"`
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
	Coin info.Coin `json:"coin" binding:"required"`
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
	Pool info.Pool `json:"pool" binding:"required"`
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
	Soft info.Soft `json:"soft" binding:"required"`
}

type AdminDelSoftReq struct {
	Name string `json:"name"`
}
