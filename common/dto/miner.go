package dto

import "miner/utils"

type CreateMinerReq struct {
	FarmID string `json:"farm_id" binding:"required,min=1,max=20"`
	Name   string `json:"name" binding:"required,min=1,max=20"`
}

type DeleteMinerReq struct {
	FarmID  string `json:"farm_id" binding:"required"`
	MinerID string `json:"miner_id" binding:"required"`
}

/*
update_info

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
		// URL
		HiveOsUrl     string `json:"hive_os_url"`
		ApiHiveOsUrls string `json:"api_hive_os_urls"`
		// Id of the rig
		RigID string `json:"rig_id"`
		// Password of the rig
		RigPasswd string `json:"rig_passwd"`
		// Rig hostname
		WorkerName string `json:"worker_name"`
		// Id of the farm
		FarmID string `json:"farm_id"`
		// Selected miners
		Miner  string `json:"miner"`
		Miner2 string `json:"miner2"`
		// Rig timezone
		TimeZone string `json:"time_zone"`
		// Watchdog
		Watchdog Watchdog `json:"watchdog"`
		// WdEnable        string `json:"wd_enable"`
		// WdMiner         string `json:"wd_miner"`
		// WdReboot        string `json:"wd_reboot"`
		// WdMaxLA         string `json:"wd_max_la"`
		// WdASR           string `json:"wd_asr"`
		// WdPowerEnabled  string `json:"wd_power_enabled"`
		// WdPowerMin      string `json:"wd_power_min"`
		// WdPowerMax      string `json:"wd_power_max"`
		// WdPowerAction   string `json:"wd_power_action"`
		// WdCheckConn     string `json:"wd_check_conn"`
		// WdShareTime     string `json:"wd_share_time"`
		// WdMinhashes     string `json:"wd_minhashes"`
		// WdMinhashesAlgo string `json:"wd_minhashes_algo"`
		// WdType          string `json:"wd_type"`
		// Options
		Options Options `json:"options"`
		// XDisabled         string `json:"x_disabled"`
		// PushInterval      string `json:"push_interval"`
		// Amintenance       string `json:"amintenance"`
		// MinerDelay        string `json:"miner_delay"`
		// DohEnable         string `json:"doh_enable"`
		// PowerRecycle      string `json:"power_recycle"`
		// ShellinaboxEnable string `json:"shellinabox_enable"`
		// SshEnable         string `json:"ssh_enable"`
		// SshPasswordEnable string `json:"ssh_password_enable"`
	HiveOsWallet  utils.HiveOsWallet  `json:"hive_os_wallet"`
	HiveOsAutoFan utils.HiveOsAutoFan `json:"hive_os_auto_fan"`
*/
type UpdateMinerReq struct {
	FarmID     string         `json:"farm_id" binding:"required"`
	MinerID    string         `json:"miner_id" binding:"required"`
	UpdateInfo map[string]any `json:"update_info" binding:"required"`
}

type UpdateMinerWatchdogReq struct {
	FarmID   string         `json:"farm_id"`
	MinerID  string         `json:"miner_id"`
	Watchdog map[string]any `json:"watchdog"`
}

type UpdateMinerOptionsReq struct {
	FarmID  string         `json:"farm_id"`
	MinerID string         `json:"miner_id"`
	Options map[string]any `json:"options"`
}

type UpdateMinerWalletReq struct {
	FarmID  string         `json:"farm_id"`
	MinerID string         `json:"miner_id"`
	Wallet  map[string]any `json:"wallet"`
}

type UpdateMinerAutofanReq struct {
	FarmID  string         `json:"farm_id"`
	MinerID string         `json:"miner_id"`
	Autofan map[string]any `json:"autofan"`
}

type ApplyMinerFsReq struct {
	FarmID  string `json:"farm_id" binding:"required"`
	MinerID string `json:"miner_id" binding:"required"`
	FsID    string `json:"fs_id" binding:"required"`
}

type TransferMinerReq struct {
	FromFarmID string `json:"from_farm_id" binding:"required"`
	MinerID    string `json:"from_miner_id" binding:"required"`
	ToUserID   string `json:"to_user_id" binding:"required"`
	ToFarmID   string `json:"to_farm_id" binding:"required"`
}

type SetWatchdogReq struct {
	FarmID   string         `json:"farm_id"`
	MinerID  string         `json:"miner_id"`
	Watchdog utils.Watchdog `json:"watchdog"`
}

type SetAutoFanReq struct {
	FarmID  string              `json:"farm_id"`
	MinerID string              `json:"miner_id"`
	AutoFan utils.HiveOsAutoFan `json:"autofan"`
}

type SetOptionsReq struct {
	FarmID  string        `json:"farm_id"`
	MinerID string        `json:"miner_id"`
	Options utils.Options `json:"options"`
}
