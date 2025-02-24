package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Server struct {
		HiveOsUrl string `mapstructure:"hive_os_url"`
		Host      string `mapstructure:"host"`
		Port      int    `mapstructure:"port"`
		Mode      string `mapstructure:"mode"`
	} `mapstructure:"server"`

	MySQL struct {
		Host         string `mapstructure:"host"`
		Port         int    `mapstructure:"port"`
		User         string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		DBName       string `mapstructure:"dbname"`
		MaxIdleConns int    `mapstructure:"max_idle_conns"`
		MaxOpenConns int    `mapstructure:"max_open_conns"`
	} `mapstructure:"mysql"`

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
	} `mapstructure:"redis"`

	JWT struct {
		Secret string `mapstructure:"secret"`
		Expire int    `mapstructure:"expire"`
	} `mapstructure:"jwt"`

	Log struct {
		Level      string `mapstructure:"level"`
		Filename   string `mapstructure:"filename"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxAge     int    `mapstructure:"max_age"`
		MaxBackups int    `mapstructure:"max_backups"`
	} `mapstructure:"log"`

	Mnemonic struct {
		Key  string `mapstructure:"key"`
		Path string `mapstructure:"path"`
	} `mapstructure:"mnemonic"`

	Bsc struct {
		Api string `mapstructure:"api"`
	} `mapstructure:"bsc"`
}

var Config ServerConfig

func InitConfig(configFile string, configType string) error {
	viper.SetConfigFile(configFile)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		return errors.New("Failed to read config " + err.Error())
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return errors.New("Failed to Unmarshal config " + err.Error())
	}

	return nil
}

type HiveOsConfig struct {
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
	WdEnable        string `json:"wd_enable"`
	WdMiner         string `json:"wd_miner"`
	WdReboot        string `json:"wd_reboot"`
	WdMaxLA         string `json:"wd_max_la"`
	WdASR           string `json:"wd_asr"`
	WdPowerEnabled  string `json:"wd_power_enabled"`
	WdPowerMin      string `json:"wd_power_min"`
	WdPowerMax      string `json:"wd_power_max"`
	WdPowerAction   string `json:"wd_power_action"`
	WdCheckConn     string `json:"wd_check_conn"`
	WdShareTime     string `json:"wd_share_time"`
	WdMinhashes     string `json:"wd_minhashes"`
	WdMinhashesAlgo string `json:"wd_minhashes_algo"`
	WdType          string `json:"wd_type"`
	// Options
	XDisabled         string `json:"x_disabled"`
	PushInterval      string `json:"push_interval"`
	Amintenance       string `json:"amintenance"`
	MinerDelay        string `json:"miner_delay"`
	DohEnable         string `json:"doh_enable"`
	PowerRecycle      string `json:"power_recycle"`
	ShellinaboxEnable string `json:"shellinabox_enable"`
	SshEnable         string `json:"ssh_enable"`
	SshPasswordEnable string `json:"ssh_password_enable"`
}

type HiveOsWallet struct {
	CustomMiner      string `json:"custom_miner"`
	CustomInstallURL string `json:"custom_install_url"`
	CustomAlgo       string `json:"custom_algo"`
	CustomTemplate   string `json:"custom_template"`
	CustomUrl        string `json:"custom_url"`
	CustomPass       string `json:"custom_pass"`
	CustomUserConfig string `json:"custom_user_config"`
	CustomTLS        string `json:"custom_tls"`
	FsID             string `json:"fs_id"`
	Coin             string `json:"coin"`
	// Meta             struct {
	// 	FsID   string `json:"fs_id"`
	// 	Custom struct {
	// 		Coin string `json:"coin"`
	// 	} `json:"custom"`
	// }
}

type HiveOsAutoFan struct {
	CriticalTemp       string `json:"critical_temp"`
	CriticalTempAction string `json:"critical_temp_action"`
	Enable             string `json:"enable"`
	TargetTemp         string `json:"target_temp"`
	MinFan             string `json:"min_fan"`
	MaxFan             string `json:"max_fan"`
	NoAMD              string `json:"no_amd"`
	TargetMemTemp      string `json:"target_mem_temp"`
	RebootOnError      string `json:"reboot_on_error"`
	SmartMode          string `json:"smart_mode"`
}

// 生成字符串会被写入 /hive-config/rig.conf
func GenerateHiveOsConfig(data *HiveOsConfig) string {
	template := `
### MINERS HIVE CONFIGS ###

# URL
HIVE_HOST_URL="%s"
API_HOST_URLs="%s"

# Id of the rig
RIG_ID=%s

# Rig password as in admin panel
RIG_PASSWD="%s"

# Rig hostname
WORKER_NAME="%s"

# Id of the farm
FARM_ID=%s

# Selected miners
MINER=%s
MINER2=%s

# Rig timezone
TIMEZONE="%s"

# Watchdog
WD_ENABLED=%s
WD_MINER=%s
WD_REBOOT=%s
WD_MAX_LA=%s
WD_ASR=%s
WD_POWER_ENABLED=%s
WD_POWER_MIN=%s
WD_POWER_MAX=%s
WD_POWER_ACTION=%s
WD_CHECK_CONN=%s
WD_SHARE_TIME=%s
WD_MINHASHES='%s'
WD_MINHASHES_ALGO='%s'
WD_TYPE='%s'

# Options
X_DISABLED=%s
PUSH_INTERVAL=%s
AMINTENANCE=%s
MINER_DELAY=%s
DOH_ENABLE=%s
POWERCYCLE=%s
SHELLINABOX_ENABLE=%s
SSH_ENABLE=%s
SSH_PASSWORD_ENABLE=%s
`
	return fmt.Sprintf(template,
		// URL
		data.HiveOsUrl,
		data.ApiHiveOsUrls,
		// Id of the rig
		data.RigID,
		// Password of the rig
		data.RigPasswd,
		// Rig hostname
		data.WorkerName,
		// Id of the farm
		data.FarmID,
		// Selected miners
		data.Miner,
		data.Miner2,
		// Rig timezone
		data.TimeZone,
		// watchdog
		data.WdEnable,
		data.WdMiner,
		data.WdReboot,
		data.WdMaxLA,
		data.WdASR,
		data.WdPowerEnabled,
		data.WdPowerMin,
		data.WdPowerMax,
		data.WdPowerAction,
		data.WdCheckConn,
		data.WdShareTime,
		data.WdMinhashes,
		data.WdMinhashesAlgo,
		data.WdType,
		// options
		data.XDisabled,
		data.PushInterval,
		data.Amintenance,
		data.MinerDelay,
		data.DohEnable,
		data.PowerRecycle,
		data.ShellinaboxEnable,
		data.SshEnable,
		data.SshPasswordEnable,
	)
}

// 生成字符串会被写入 /hive-config/wallet.conf
func GenerateHiveOsWallet(data *HiveOsWallet) string {
	template := `
### FLIGHT SHEET

# Miner custom
CUSTOM_MINER="%s"
CUSTOM_INSTALL_URL="%s"
CUSTOM_ALGO="%s"
CUSTOM_TEMPLATE="%s"
CUSTOM_URL="%s"
CUSTOM_PASS="%s"
CUSTOM_USER_CONFIG='%s'
CUSTOM_TLS=""
META='{
	"fs_id":%s,
	"custom": {
		"coin":"%s"
	}
}'
`
	return fmt.Sprintf(template,
		data.CustomMiner, data.CustomInstallURL, data.CustomAlgo,
		data.CustomTemplate, data.CustomUrl, data.CustomPass,
		data.CustomUserConfig, data.FsID, data.Coin,
	)
}

// 生成字符串会被写入 /hive-config/autofan.conf
func GenerateHiveOsAutofan(data *HiveOsAutoFan) string {
	template := `\
CRITICAL_TEMP=%s
CRITICAL_TEMP_ACTION="%s"
ENABLED=%s
TARGET_TEMP=%s
MIN_FAN=%s
MAX_FAN=%s
NO_AMD=%s
TARGET_MEM_TEMP=%s
REBOOT_ON_ERROR=%s
SMART_MODE=%s
`
	return fmt.Sprintf(template,
		data.CriticalTemp, data.CriticalTempAction, data.Enable,
		data.TargetTemp, data.MinFan, data.MaxFan, data.NoAMD,
		data.TargetMemTemp, data.RebootOnError, data.SmartMode,
	)
}

// 生成 hiveosurl
func GenerateHiveOsUrl() string {
	host := Config.Server.Host
	port := Config.Server.Port
	return fmt.Sprintf("http://%s:%d", host, port)
	// hiveOsUrl := Config.Server.HiveOsUrl
	// return fmt.Sprintf(hiveOsUrl)
}

func GeneratePort() string {
	return ":" + strconv.Itoa(Config.Server.Port)
}
