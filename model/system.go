package model

import "miner/common/status"

type System struct {
	InviteReward   int                   `json:"invite_reward" gorm:"column:invite_reward;type:int;comment:邀请积分奖励"`
	RechargeReward int                   `json:"recharge_reward" gorm:"column:recharge_reward;type:int;comment:充值积分奖励"`
	SwitchRegister status.RegisterStatus `json:"switch_register" gorm:"column:switch_register;type:int;comment:注册开关"`
}

func (System) TableName() string {
	return "system"
}
