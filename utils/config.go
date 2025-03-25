package utils

import (
	"fmt"
	"log"
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

func InitConfig(configFile string, configType string) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read config file: %s, %s", configFile, err.Error())
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("failed to unmarshal config file: %s, %s", configFile, err.Error())
	}
}

type HiveOsConfig struct {
	// URL
	HiveOsUrl     string `json:"hive_os_url"`
	ApiHiveOsUrls string `json:"api_hive_os_urls"`
	// Id of the rig
	RigID int `json:"rig_id"`
	// Password of the rig
	RigPasswd string `json:"rig_passwd"`
	// Rig hostname
	WorkerName string `json:"worker_name"`
	// Id of the farm
	FarmID int `json:"farm_id"`
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
}

type Watchdog struct {
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
	WdCheckGpu      string `json:"wd_check_gpu"`
}

type Options struct {
	XDisabled         string `json:"x_disabled"`
	PushInterval      string `json:"push_interval"`
	Maintenance       string `json:"maintenance"`
	MinerDelay        string `json:"miner_delay"`
	DohEnable         string `json:"doh_enable"`
	PowerRecycle      string `json:"power_recycle"`
	ShellinaboxEnable string `json:"shellinabox_enable"`
	SshEnable         string `json:"ssh_enable"`
	SshPasswordEnable string `json:"ssh_password_enable"`
}

type HiveOsWallet struct {
	// Custom miner soft
	CustomMiner      string `json:"custom_miner"`
	CustomInstallURL string `json:"custom_install_url"`
	CustomAlgo       string `json:"custom_algo"`
	CustomTemplate   string `json:"custom_template"`
	CustomUrl        string `json:"custom_url"`
	CustomPass       string `json:"custom_pass"`
	CustomUserConfig string `json:"custom_user_config"`
	CustomTLS        string `json:"custom_tls"`
	// fs coin
	FsID string `json:"fs_id"`
	Coin string `json:"coin"`
}

type HiveOsAutoFan struct {
	Enable              string `json:"enable"`
	TargetTemp          string `json:"target_temp"`
	TargetMemTemp       string `json:"target_mem_temp"`
	MinFan              string `json:"min_fan"`
	MaxFan              string `json:"max_fan"`
	CriticalTemp        string `json:"critical_temp"`
	CriticalTempAction  string `json:"critical_temp_action"`
	NoAMD               string `json:"no_amd"`
	RebootOnError       string `json:"reboot_on_error"`
	SmartMode           string `json:"smart_mode"`
	CustomMode          string `json:"custom_mode"`
	CustomTargetTemp    string `json:"custom_target_temp"`
	CustomTargetMemTemp string `json:"custom_target_mem_temp"`
	CustomMinFan        string `json:"custom_min_fan"`
	CustomMaxFan        string `json:"custom_max_fan"`
	CustomCriticalTemp  string `json:"custom_critical_temp"`
	CustomStaticFan     string `json:"custom_static_fan"`
}

// 生成字符串会被写入 /hive-config/rig.conf
func GenerateHiveOsConfig(data *HiveOsConfig) string {
	template := `### MINERS HIVE CONFIGS ###
# URL
HIVE_HOST_URL="%s"
API_HOST_URLs="%s"

# Id of the rig
RIG_ID=%d

# Rig password as in admin panel
RIG_PASSWD="%s"

# Rig hostname
WORKER_NAME="%s"

# Id of the farm
FARM_ID=%d

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
		data.Watchdog.WdEnable,
		data.Watchdog.WdMiner,
		data.Watchdog.WdReboot,
		data.Watchdog.WdMaxLA,
		data.Watchdog.WdASR,
		data.Watchdog.WdPowerEnabled,
		data.Watchdog.WdPowerMin,
		data.Watchdog.WdPowerMax,
		data.Watchdog.WdPowerAction,
		data.Watchdog.WdCheckConn,
		data.Watchdog.WdShareTime,
		data.Watchdog.WdMinhashes,
		data.Watchdog.WdMinhashesAlgo,
		data.Watchdog.WdType,
		// options
		data.Options.XDisabled,
		data.Options.PushInterval,
		data.Options.Maintenance,
		data.Options.MinerDelay,
		data.Options.DohEnable,
		data.Options.PowerRecycle,
		data.Options.ShellinaboxEnable,
		data.Options.SshEnable,
		data.Options.SshPasswordEnable,
	)
}

// 生成字符串会被写入 /hive-config/wallet.conf
func GenerateHiveOsWallet(data *HiveOsWallet) string {
	template := `### FLIGHT SHEET
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
		data.CustomMiner,
		data.CustomInstallURL,
		data.CustomAlgo,
		data.CustomTemplate,
		data.CustomUrl,
		data.CustomPass,
		data.CustomUserConfig,
		data.FsID,
		data.Coin,
	)
}

// 生成字符串会被写入 /hive-config/autofan.conf
func GenerateHiveOsAutofan(data *HiveOsAutoFan) string {
	template := `### autofan
ENABLED=%s
TARGET_TEMP=%s
TARGET_MEM_TEMP=%s
MIN_FAN=%s
MAX_FAN=%s
CRITICAL_TEMP=%s
CRITICAL_TEMP_ACTION="%s"
NO_AMD=%s
REBOOT_ON_ERROR=%s
SMART_MODE=%s
CUSTOM_MODE="%s"
CUSTOM_TARGET_TEMP="%s"
CUSTOM_TARGET_MEM_TEMP="%s"
CUSTOM_MIN_FAN="%s"
CUSTOM_MAX_FAN="%s"
CUSTOM_CRITICAL_TEMP="%s"
`
	return fmt.Sprintf(template,
		data.Enable,
		data.TargetTemp,
		data.TargetMemTemp,
		data.MinFan,
		data.MaxFan,
		data.CriticalTemp,
		data.CriticalTempAction,
		data.NoAMD,
		data.RebootOnError,
		data.SmartMode,
		data.CustomMode,
		data.CustomTargetTemp,
		data.CustomTargetMemTemp,
		data.CustomMinFan,
		data.CustomMaxFan,
		data.CustomCriticalTemp,
	)
}

// 生成 hiveosurl
func GenerateHiveOsUrl() string {
	return Config.Server.HiveOsUrl
}

func GeneratePort() string {
	return ":" + strconv.Itoa(Config.Server.Port)
}
