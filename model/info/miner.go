package info

import (
	"miner/utils"
)

type Miner struct {
	GpuNum int    `json:"gpu_num"`
	UUID   string `json:"uuid"`

	// HiveOS 单独客户端配置
	HiveOsConfig  utils.HiveOsConfig  `json:"hive_os_config"`
	HiveOsWallet  utils.HiveOsWallet  `json:"hive_os_wallet"`
	HiveOsAutoFan utils.HiveOsAutoFan `json:"hive_os_auto_fan"`
}
