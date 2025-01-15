package info

import (
	"miner/common/perm"
	"miner/utils"
)

type Miner struct {
	ID    string         `json:"id"`
	Name  string         `json:"name"`
	RigID string         `json:"rig_id"`
	Pass  string         `json:"pass"`
	FS    string         `json:"fs"`
	Perm  perm.MinerPerm `json:"perm"`

	// HiveOS 单独客户端配置
	HiveOsConfig  utils.HiveOsConfig  `json:"hive_os_config"`
	HiveOsWallet  utils.HiveOsWallet  `json:"hive_os_wallet"`
	HiveOsAutoFan utils.HiveOsAutoFan `json:"hive_os_auto_fan"`
}
