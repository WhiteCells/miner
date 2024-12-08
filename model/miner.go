package model

import "miner/common/status"

type Miner struct {
	ID          int                `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:矿机唯一标识"`
	Name        string             `json:"name" gorm:"column:name;type:varchar(255);comment:矿机名"`
	GPUInfo     string             `json:"gpu_info" gorm:"column:gpu_info;type:varchar(255);comment:矿机显卡信息"`
	Status      status.MinerStatus `json:"status" gorm:"column:status;type:int;comment:矿机状态"`
	IP          string             `json:"ip" gorm:"column:ip;type:varchar(255);comment:矿机IP"`
	SSHPort     int                `json:"ssh_port" gorm:"column:ssh_port;type:int;comment:矿机SSH端口"`
	SSHUser     string             `json:"ssh_user" gorm:"column:ssh_user;type:varchar(255);comment:矿机SSH用户名"`
	SSHPassword string             `json:"ssh_password" gorm:"column:ssh_password;type:varchar(255);comment:矿机SSH密码"`
}

func (Miner) TableName() string {
	return "miner"
}
