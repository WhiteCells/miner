package model

import "miner/common/status"

type System struct {
	InviteReward   float32               `json:"invite_reward" gorm:"column:invite_reward;type:float;comment:邀请积分奖励"`
	RechargeRatio  float32               `json:"recharge_ratio" gorm:"column:recharge_ratio;type:float;comment:充值比率"`
	RechargeReward float32               `json:"recharge_reward" gorm:"column:recharge_reward;type:float;comment:充值积分奖励"`
	SwitchRegister status.RegisterStatus `json:"switch_register" gorm:"column:switch_register;type:varchar(16);comment:注册开关"`
	FreeGpuNum     int                   `json:"free_gpu_num" gorm:"column:free_gpu_num;type:int;comment:免费GPU数量"`
}

func (System) TableName() string {
	return "system"
}
