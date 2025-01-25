package info

import (
	"miner/common/perm"
	"miner/utils"
	"time"
)

type Miner struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	RigID       string         `json:"rig_id"`
	Pass        string         `json:"pass"`
	FS          string         `json:"fs"`
	Perm        perm.MinerPerm `json:"perm"`
	GpuNum      int            `json:"gpu_num"`
	UUID        string         `json:"uuid"`
	LastFlushAt time.Time      `json:"last_flush_at"`

	// HiveOS 单独客户端配置
	HiveOsConfig  utils.HiveOsConfig  `json:"hive_os_config"`
	HiveOsWallet  utils.HiveOsWallet  `json:"hive_os_wallet"`
	HiveOsAutoFan utils.HiveOsAutoFan `json:"hive_os_auto_fan"`
}
