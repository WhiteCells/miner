package service

import (
	"context"
	"errors"
	"fmt"
	"miner/common/dto"
	"miner/common/rsp"
	"miner/dao/redis"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type HiveOsService struct {
	hiveOsRDB *redis.HiveOsRDB
	farmRDB   *redis.FarmRDB
	minerRDB  *redis.MinerRDB
	taskRDB   *redis.TaskRDB
	fsRDB     *redis.FsRDB
}

func NewHiveOsService() *HiveOsService {
	return &HiveOsService{
		hiveOsRDB: redis.NewHiveOsRDB(),
		farmRDB:   redis.NewFarmRDB(),
		minerRDB:  redis.NewMinerRDB(),
		taskRDB:   redis.NewTaskRDB(),
		fsRDB:     redis.NewFsRDB(),
	}
}

func (s *HiveOsService) Poll(ctx *gin.Context, rigID string, method string) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}

	// 通过 rigID 查找 Miner
	farmMinerID, err := s.hiveOsRDB.GetRigMinerID(ctx, rigID)
	if err != nil {
		return err
	}
	parts := strings.Split(farmMinerID, ":")
	farmID := parts[0]
	minerID := parts[1]

	farm, err := s.farmRDB.GetByID(ctx, userID, farmID)
	if err != nil {
		return err
	}

	miner, err := s.minerRDB.GetByID(ctx, farmID, minerID)
	if err != nil {
		return err
	}

	// 通过 miner 找到其飞行表 ID
	fsID, err := s.minerRDB.GetApplyFs(ctx, minerID)
	if err != nil {
		return err
	}

	// 通过飞行表 ID 查找对应飞行表
	_, err = s.fsRDB.GetByID(ctx, userID, fsID)
	if err != nil {
		return err
	}

	// 通过飞行表 ID 查找对应的钱包

	// 构造 config、wallet及autofun
	// RigID
	rigIDInt, err := strconv.Atoi(miner.ID)
	if err != nil {
		return err
	}
	config.RigID = rigIDInt
	// RigPasswd
	config.RigPasswd = miner.Pass
	// WorkerName
	config.WorkerName = rigID
	// FarmID
	farmIDInt, err := strconv.Atoi(farmID)
	if err != nil {
		return err
	}
	config.FarmID = farmIDInt
	// Miner 挖矿软件
	// config.Miner = miner.Miner
	// Timezone
	config.Timezone = farm.TimeZone

	// autofan

	// wallet

	// 从缓存中拿出任务
	task, err := s.taskRDB.LPop(ctx)
	if err != nil {
		return err
	}

	// 任务类型
	switch task.Type {
	case "stats":
		var req dto.HiveosReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
			return err
		}
		// 检查 RigID 及 RigPasswd
		if miner.ID != req.Params.RigID || miner.Pass != req.Params.Passwd {
			rsp.Error(ctx, http.StatusUnauthorized, err.Error(), "")
			return err
		}
		//
	case "message":
		var req dto.HiveosResReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
			return err
		}
		// 检查 RigID 及 RigPasswd
		if miner.ID != req.Params.RigID || miner.Pass != req.Params.Passwd {
			rsp.Error(ctx, http.StatusUnauthorized, err.Error(), "")
			return err
		}
		//
		// 存储任务结果
	}

	return nil
}

func (s *HiveOsService) SendCmd(ctx context.Context) error {

	return nil
}

func (s *HiveOsService) GetCmdRes(ctx context.Context) error {

	return nil
}

func (s *HiveOsService) SetConfig(ctx context.Context) error {

	return nil
}

func (s *HiveOsService) GetConfigRes(ctx context.Context) error {

	return nil
}

func (s *HiveOsService) GetStats(ctx context.Context) error {

	return nil
}

// 生成

// 不知道大小会不会有影响
type ConfigParams struct {
	HiveHostURL string `json:"hive_host_url"`
	APIHostURLs string `json:"api_host_ur_ls"`
	RigID       int    `json:"rig_id"`
	RigPasswd   string `json:"rig_passwd"`
	WorkerName  string `json:"worker_name"`
	FarmID      int    `json:"farm_id"`
	Miner       string `json:"miner"`
	Timezone    string `json:"timezone"`
	WDEnabled   int    `json:"wd_enabled"`
	SSHEnable   int    `json:"ssh_enable"`
	SSHPassword int    `json:"ssh_password"`
	ShellInABox int    `json:"shell_in_a_box"`
}

type WalletParams struct {
	CustomMiner      string `json:"custom_miner"`
	CustomInstallURL string `json:"custom_install_url"`
	CustomAlgo       string `json:"custom_algo"`
	CustomTemplate   string `json:"custom_template"`
	CustomURL        string `json:"custom_url"`
	CustomUserConfig string `json:"custom_user_config"`
	CustomTLS        string `json:"custom_tls"`
	// {\"fs_id\":20216083,\"custom\":{\"coin\":\"smh\"}}
	Meta struct {
		FsID   int `json:"fs_id"`
		Custom struct {
			Coin string `json:"coin"`
		} `json:"custom"`
	} `json:"meta"`
}

type AutoFanParams struct {
	Enabled            string `json:"enabled"`
	TargetTemp         string `json:"target_temp"`
	TargetMemTemp      string `json:"target_mem_temp"`
	MinFan             string `json:"min_fan"`
	MaxFan             string `json:"max_fan"`
	CriticalTemp       string `json:"critical_temp"`
	CriticalTempAction string `json:"critical_temp_action"`
}

func (c *HiveOsService) generateConfig(params ConfigParams) string {
	return fmt.Sprintf(`HIVE_HOST_URL="%s"
API_HOST_URLs="%s"
RIG_ID=%d
RIG_PASSWD="%s"
WORKER_NAME="%s"
FARM_ID=%d
MINER=%s
TIMEZONE="%s"
WD_ENABLED=%d
SHELLINABOX_ENABLE=%d
SSH_ENABLE=%d
SSH_PASSWORD_ENABLE=%d
`, params.HiveHostURL, params.APIHostURLs, params.RigID, params.RigPasswd, params.WorkerName, params.FarmID, params.Miner, params.Timezone, params.WDEnabled, params.ShellInABox, params.SSHEnable, params.SSHPassword)
}

func (c *HiveOsService) generateWallet(params WalletParams) string {
	return fmt.Sprintf(`### Wallet 
CUSTOM_MINER="%s"
CUSTOM_INSTALL_URL="%s"
CUSTOM_ALGO="%s"
CUSTOM_TEMPLATE="%s"
CUSTOM_URL="%s"
CUSTOM_USER_CONFIG='%s'
CUSTOM_TLS="%s"
META='%s'
`, params.CustomMiner, params.CustomInstallURL, params.CustomAlgo, params.CustomTemplate, params.CustomURL, params.CustomUserConfig, params.CustomTLS, params.Meta)
}

func (c *HiveOsService) GenerateAutofan(params AutoFanParams) string {
	return fmt.Sprintf(`ENABLED=%s
TARGET_TEMP=%s
TARGET_MEM_TEMP=%s
MIN_FAN=%s
MAX_FAN=%s
CRITICAL_TEMP=%s
CRITICAL_TEMP_ACTION="%s"
`, params.Enabled, params.TargetTemp, params.TargetMemTemp, params.MinFan, params.MaxFan, params.CriticalTemp, params.CriticalTempAction)
}

var config = ConfigParams{
	HiveHostURL: "http://172.16.0.176:9090/hiveos",
	APIHostURLs: "http://172.16.0.176:9090/hiveos",
	RigID:       10101,
	RigPasswd:   "1q2w3e4r",
	WorkerName:  "15",
	FarmID:      3335302,
	Miner:       "custom",
	Timezone:    "Europe/Kiev",
	WDEnabled:   1,
	ShellInABox: 1,
	SSHEnable:   1,
	SSHPassword: 1,
}

var wallet = WalletParams{
	CustomMiner:      "",
	CustomInstallURL: "",
	CustomAlgo:       "",
	CustomTemplate:   "",
	CustomURL:        "",
	CustomUserConfig: "",
	CustomTLS:        "",
	Meta: struct {
		FsID   int `json:"fs_id"`
		Custom struct {
			Coin string `json:"coin"`
		} `json:"custom"`
	}{
		FsID: 1,
		Custom: struct {
			Coin string `json:"coin"`
		}{
			Coin: "",
		},
	},
}

var autofan = AutoFanParams{
	Enabled:            "",
	TargetTemp:         "",
	TargetMemTemp:      "",
	MinFan:             "",
	MaxFan:             "",
	CriticalTemp:       "",
	CriticalTempAction: "",
}
