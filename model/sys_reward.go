package model

type SysReward struct {
	InviteReward   int `json:"invite_reward" gorm:"column:invite_reward;type:int"`
	RechargeReward int `json:"recharge_reward" gorm:"column:recharge_reward;type:int"`
}

func (SysReward) TableName() string {
	return "sys_reward"
}
