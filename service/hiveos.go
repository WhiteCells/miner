package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"miner/common/dto"
	"miner/common/rsp"
	"miner/dao/redis"
	"miner/model/info"
	"miner/utils"
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

func (s *HiveOsService) Poll(ctx *gin.Context, rigID string, method string) {
	switch method {
	case "hello":
		var req dto.HiveosReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
			return
		}
		////////////////////////////////////////////////
		jsonInd, err := json.MarshalIndent(req, "", "  ")
		if err != nil {
			fmt.Println("")
			return
		}
		fmt.Printf("%s\n", jsonInd)
		////////////////////////////////////////////////
		// 对 req 的数据进行存储
		s.setMinerStatus(ctx, rigID, &req)
		// 从 req 中获取 rigID，根据 rigID 查询 hiveOsRDB farmID:minerID
		rigID := req.Params.RigID
		farmMiner, err := s.hiveOsRDB.GetRigMinerID(ctx, rigID)
		if err != nil {
			log.Println(rigID)
			log.Println("hiveOsRDB.GetRigMinerID")
			rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
			return
		}
		parts := strings.Split(farmMiner, ":")
		farmID := parts[0]
		minerID := parts[1]
		// 通过 farmID 和 minerID 获取 miner
		_, err = s.minerRDB.GetByID(ctx, farmID, minerID)
		if err != nil {
			log.Println(farmID, minerID)
			log.Fatalln("minerRDB.GetByID")
			rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
			return
		}
		// 通过 minerID 获取其 fsID
		// _, err = s.minerRDB.GetApplyFs(ctx, minerID)
		// if err != nil {
		// 	log.Println(minerID)
		// 	log.Fatalln("minerRDB.GetApplyFs")
		// 	rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		//  没找到不退出
		// 	return
		// }
		rigIDInt, err := strconv.Atoi(rigID)
		if err != nil {
			log.Fatalln("strconv.Atoi")
			return
		}
		ctx.JSON(http.StatusOK, &dto.ServerRsp{
			ID:      rigIDInt,
			Jsonrpc: "2.0",
			Result: struct {
				ID        int    `json:"id"`
				Config    string `json:"config"`
				Wallet    string `json:"wallet"`
				Autofan   string `json:"autofan"`
				Justwrite int    `json:"justwrite"`
				Command   string `json:"command"`
				Exec      string `json:"exec"`
				Confseq   int    `json:"confseq"`
			}{
				ID:        99999,
				Config:    "HIVE_HOST_URL=\"http://172.16.0.176:9090/hiveos\"\nAPI_HOST_URLs=\"http://172.16.0.176.4:9090/hiveos\"\nRIG_ID=00666304\nRIG_PASSWD=\"ScPzqqnZ\"\nWORKER_NAME=\"15\"\nFARM_ID=3335302\nMINER=custom\nMINER2=\nTIMEZONE=\"Europe/Kiev\"\nWD_ENABLED=1\nWD_MINER=3\nWD_REBOOT=5\nWD_CHECK_GPU=0\nWD_MAX_LA=900\nWD_ASR=\nWD_POWER_ENABLED=0\nWD_POWER_MIN=\nWD_POWER_MAX=\nWD_POWER_ACTION=\nWD_CHECK_CONN=0\nWD_SHARE_TIME=\nWD_MINHASHES='{}'\nWD_MINHASHES_ALGO='{}'\nWD_TYPE='miner'\nHSSH_SRV=\"http://192.168.35.15:9090/hiveos\"\nX_DISABLED=1\nMINER_DELAY=1\nDOH_ENABLED=0\nSHELLINABOX_ENABLE=1\nSSH_ENABLE=1\nSSH_PASSWORD_ENABLE=1\n",
				Wallet:    "# Miner custom\nCUSTOM_MINER=\"k3ok.com-spacemesh-s\"\nCUSTOM_INSTALL_URL=\"https://gitee.com/k3os/spacemesh/releases/download/v4.0.5/k3ok.com-spacemesh-s-v4.0.5.tar.gz\"\nCUSTOM_ALGO=\"\"\nCUSTOM_TEMPLATE=\"15\"\nCUSTOM_URL=\"http://hiveos.vip/\"\nCUSTOM_PASS=\"\"\nCUSTOM_USER_CONFIG='path:\n- /mnt/\nminerName: 15\napiKey: smh00000-0c79-5659-7b8f-565a95961ecf\nextraParams:\n  deleteLoadFail: false\n  device: \"\"\n  disableInitPost: false\n  disablePlot: true\n  disablePoST: false\n  flags: fullmem\n  maxFileSize: 32\n  nonces: 128\n  numUnits: 15\n  plotInstance: 1\n  postAffinity: 0\n  postAffinityStep: 1\n  postCpuIds: \"\"\n  postInstance: 0\n  postThread: 0\n  randomxAffinity: -1\n  randomxAffinityStep: 1\n  randomxThread: 0\n  removeInitFailed: false\n  reservedSize: 1\n  skipUninitialized: false\n  remoteK2Pow: true\nlog:\n  lv: info\n  path: ./log/\n  name: miner.log\nurl:\n  info: \"\"\n  submit: \"\"\n  line: \"\"\n  ws: \"\"\n  proxy: \"http://172.16.10.77:9090\"\nproxy:\n  url: \"\"\n  username: \"\"\n  password: \"\"\nhttp:\n  enable: false\n  host: \"\"\n  port: 0\nscanPath: false\nscanMinute: 60\ndebug: \"\"'\nCUSTOM_TLS=\"\"\n\nMETA='{\"fs_id\":20216083,\"custom\":{\"coin\":\"smh\"}}'\n",
				Autofan:   "ENABLED=\nTARGET_TEMP=\nTARGET_MEM_TEMP=\nMIN_FAN=\nMAX_FAN=\nCRITICAL_TEMP=\nCRITICAL_TEMP_ACTION=\"\"\nNO_AMD=\nREBOOT_ON_ERROR=\nSMART_MODE=\nCUSTOM_MODE=\"\"\nCUSTOM_TARGET_TEMP=\"\"\nCUSTOM_TARGET_MEM_TEMP=\"\"\nCUSTOM_MIN_FAN=\"\"\nCUSTOM_MAX_FAN=\"\"\nCUSTOM_CRITICAL_TEMP=\"\"\n",
				Justwrite: 1,
				Confseq:   1,
			},
		})
	case "stats":
		var req dto.HiveosReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
			return
		}
		////////////////////////////////////////////////
		jsonInd, err := json.MarshalIndent(req, "", "  ")
		if err != nil {
			fmt.Println("")
			return
		}
		fmt.Printf("%s\n", jsonInd)
		////////////////////////////////////////////////
		// 对 req 的数据进行存储
		s.setMinerStatus(ctx, rigID, &req)
		//
		rigIDInt, err := strconv.Atoi(rigID)
		if err != nil {
			log.Fatalln("strconv.Atoi")
			return
		}
		// 从任务缓存中取出 rigID 对应的任务
		// s.taskRDB.Get()
		ctx.JSON(http.StatusOK, &dto.ServerRsp{
			ID:      rigIDInt,
			Jsonrpc: "2.0",
			Result: struct {
				ID        int    `json:"id"`
				Config    string `json:"config"`
				Wallet    string `json:"wallet"`
				Autofan   string `json:"autofan"`
				Justwrite int    `json:"justwrite"`
				Command   string `json:"command"`
				Exec      string `json:"exec"`
				Confseq   int    `json:"confseq"`
			}{
				ID:        99999,
				Config:    "HIVE_HOST_URL=\"http://172.16.0.176:9090/hiveos\"\nAPI_HOST_URLs=\"http://172.16.0.176.4:9090/hiveos\"\nRIG_ID=00666304\nRIG_PASSWD=\"ScPzqqnZ\"\nWORKER_NAME=\"15\"\nFARM_ID=3335302\nMINER=custom\nMINER2=\nTIMEZONE=\"Europe/Kiev\"\nWD_ENABLED=1\nWD_MINER=3\nWD_REBOOT=5\nWD_CHECK_GPU=0\nWD_MAX_LA=900\nWD_ASR=\nWD_POWER_ENABLED=0\nWD_POWER_MIN=\nWD_POWER_MAX=\nWD_POWER_ACTION=\nWD_CHECK_CONN=0\nWD_SHARE_TIME=\nWD_MINHASHES='{}'\nWD_MINHASHES_ALGO='{}'\nWD_TYPE='miner'\nHSSH_SRV=\"http://192.168.35.15:9090/hiveos\"\nX_DISABLED=1\nMINER_DELAY=1\nDOH_ENABLED=0\nSHELLINABOX_ENABLE=1\nSSH_ENABLE=1\nSSH_PASSWORD_ENABLE=1\n",
				Wallet:    "# Miner custom\nCUSTOM_MINER=\"k3ok.com-spacemesh-s\"\nCUSTOM_INSTALL_URL=\"https://gitee.com/k3os/spacemesh/releases/download/v4.0.5/k3ok.com-spacemesh-s-v4.0.5.tar.gz\"\nCUSTOM_ALGO=\"\"\nCUSTOM_TEMPLATE=\"15\"\nCUSTOM_URL=\"http://hiveos.vip/\"\nCUSTOM_PASS=\"\"\nCUSTOM_USER_CONFIG='path:\n- /mnt/\nminerName: 15\napiKey: smh00000-0c79-5659-7b8f-565a95961ecf\nextraParams:\n  deleteLoadFail: false\n  device: \"\"\n  disableInitPost: false\n  disablePlot: true\n  disablePoST: false\n  flags: fullmem\n  maxFileSize: 32\n  nonces: 128\n  numUnits: 15\n  plotInstance: 1\n  postAffinity: 0\n  postAffinityStep: 1\n  postCpuIds: \"\"\n  postInstance: 0\n  postThread: 0\n  randomxAffinity: -1\n  randomxAffinityStep: 1\n  randomxThread: 0\n  removeInitFailed: false\n  reservedSize: 1\n  skipUninitialized: false\n  remoteK2Pow: true\nlog:\n  lv: info\n  path: ./log/\n  name: miner.log\nurl:\n  info: \"\"\n  submit: \"\"\n  line: \"\"\n  ws: \"\"\n  proxy: \"http://172.16.10.77:9090\"\nproxy:\n  url: \"\"\n  username: \"\"\n  password: \"\"\nhttp:\n  enable: false\n  host: \"\"\n  port: 0\nscanPath: false\nscanMinute: 60\ndebug: \"\"'\nCUSTOM_TLS=\"\"\n\nMETA='{\"fs_id\":20216083,\"custom\":{\"coin\":\"smh\"}}'\n",
				Autofan:   "ENABLED=\nTARGET_TEMP=\nTARGET_MEM_TEMP=\nMIN_FAN=\nMAX_FAN=\nCRITICAL_TEMP=\nCRITICAL_TEMP_ACTION=\"\"\nNO_AMD=\nREBOOT_ON_ERROR=\nSMART_MODE=\nCUSTOM_MODE=\"\"\nCUSTOM_TARGET_TEMP=\"\"\nCUSTOM_TARGET_MEM_TEMP=\"\"\nCUSTOM_MIN_FAN=\"\"\nCUSTOM_MAX_FAN=\"\"\nCUSTOM_CRITICAL_TEMP=\"\"\n",
				Justwrite: 1,
				Confseq:   1,
			},
		})
	case "message":

	}
}

func (s *HiveOsService) setMinerStatus(ctx context.Context, rigID string, req *dto.HiveosReq) error {
	var status info.MinerStatus
	status.Algo = req.Params.MinerStats.Algo
	status.BusNumbers = req.Params.MinerStats.BusNumbers
	status.Coin = req.Params.Meta.Custom.Coin
	status.Cpuavg = req.Params.Cpuavg
	status.Cputemp = req.Params.Cputemp
	status.Df = req.Params.Df
	status.Fan = req.Params.Fan
	status.FsID = req.Params.Meta.FsID
	status.Hs = req.Params.MinerStats.Hs
	status.HsUnits = req.Params.MinerStats.HsUnits
	status.Khs = req.Params.MinerStats.Khs
	status.Mem = req.Params.Mem
	status.Miner = req.Params.Miner
	status.Power = req.Params.Power
	status.Status = req.Params.MinerStats.Status
	status.Temp = req.Params.Temp
	status.TotalKhs = req.Params.TotalKhs
	return s.hiveOsRDB.SetMinerStatus(ctx, rigID, &status)
}

func (s *HiveOsService) SendCmd(ctx context.Context, req *dto.SendCmdReq) (string, error) {
	// 生成任务 ID
	taskID, err := utils.GenerateUID()
	if err != nil {
		return "", err
	}
	task := &info.Task{
		ID:      taskID,
		Type:    "cmd",
		Content: req.Cmd,
	}
	// 向 taskRDB 中添加任务
	err = s.taskRDB.RPush(ctx, task)
	if err != nil {
		return "", err
	}
	return taskID, nil
}

func (s *HiveOsService) GetCmdRes(ctx context.Context, taskID string) (string, error) {
	return s.taskRDB.Get(ctx, taskID)
}

func (s *HiveOsService) SetConfig(ctx context.Context) error {
	// return s.taskRDB.Set(ctx)
	return nil
}

func (s *HiveOsService) GetConfigRes(ctx context.Context) error {

	return nil
}

func (s *HiveOsService) GetStats(ctx context.Context) error {

	return nil
}

// 生成

// 不知道大小写会不会有影响
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
MINER="%s"
TIMEZONE="%s"
WD_ENABLED=%d
SHELLINABOX_ENABLE=%d
SSH_ENABLE=%d
SSH_PASSWORD_ENABLE=%d`,
		params.HiveHostURL, params.APIHostURLs, params.RigID, params.RigPasswd, params.WorkerName, params.FarmID, params.Miner, params.Timezone, params.WDEnabled, params.ShellInABox, params.SSHEnable, params.SSHPassword)
}

func (c *HiveOsService) generateWallet(params WalletParams) string {
	return fmt.Sprintf(`CUSTOM_MINER="%s"
CUSTOM_INSTALL_URL="%s"
CUSTOM_ALGO="%s"
CUSTOM_TEMPLATE="%s"
CUSTOM_URL="%s"
CUSTOM_USER_CONFIG="%s"
CUSTOM_TLS="%s"
META="%s"`,
		params.CustomMiner, params.CustomInstallURL, params.CustomAlgo, params.CustomTemplate, params.CustomURL, params.CustomUserConfig, params.CustomTLS, params.Meta)
}

func (c *HiveOsService) generateAutofan(params AutoFanParams) string {
	return fmt.Sprintf(`ENABLED="%s"
TARGET_TEMP="%s"
TARGET_MEM_TEMP="%s"
MIN_FAN="%s"
MAX_FAN="%s"
CRITICAL_TEMP="%s"
CRITICAL_TEMP_ACTION="%s"`,
		params.Enabled, params.TargetTemp, params.TargetMemTemp, params.MinFan, params.MaxFan, params.CriticalTemp, params.CriticalTempAction)
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

// 生成Config字符串
func GenerateConfig(rigID, farmID int, workerName, hiveHostURL, apiHostURL, sshServer string) string {
	return fmt.Sprintf(`HIVE_HOST_URL="%s"
API_HOST_URLs="%s"
RIG_ID=%d
RIG_PASSWD="1q2w3e4r"
WORKER_NAME="%s"
FARM_ID=%d
MINER=custom
MINER2=
TIMEZONE="Europe/Kiev"
WD_ENABLED=1
WD_MINER=3
WD_REBOOT=5
WD_CHECK_GPU=0
WD_MAX_LA=900
WD_ASR=
WD_POWER_ENABLED=0
WD_POWER_MIN=
WD_POWER_MAX=
WD_POWER_ACTION=
WD_CHECK_CONN=0
WD_SHARE_TIME=
WD_MINHASHES='{}'
WD_MINHASHES_ALGO='{}'
WD_TYPE='miner'
HSSH_SRV="%s"
X_DISABLED=1
MINER_DELAY=1
DOH_ENABLED=0
SHELLINABOX_ENABLE=1
SSH_ENABLE=1
SSH_PASSWORD_ENABLE=1
`, hiveHostURL, apiHostURL, rigID, workerName, farmID, sshServer)
}

// 生成Wallet字符串
func GenerateWallet(customMiner, installURL, template, customProxy string) string {
	return fmt.Sprintf(`# Miner custom
CUSTOM_MINER="%s"
CUSTOM_INSTALL_URL="%s"
CUSTOM_ALGO=""
CUSTOM_TEMPLATE="%s"
CUSTOM_URL="http://hiveos.vip/"
CUSTOM_PASS=""
CUSTOM_USER_CONFIG='path:
- /mnt/
minerName: %s
apiKey: smh00000-0c79-5659-7b8f-565a95961ecf
extraParams:
  deleteLoadFail: false
  device: ""
  disableInitPost: false
  disablePlot: true
  disablePoST: false
  flags: fullmem
  maxFileSize: 32
  nonces: 128
  numUnits: %s
  plotInstance: 1
  postAffinity: 0
  postAffinityStep: 1
  postCpuIds: ""
  postInstance: 0
  postThread: 0
  randomxAffinity: -1
  randomxAffinityStep: 1
  randomxThread: 0
  removeInitFailed: false
  reservedSize: 1
  skipUninitialized: false
  remoteK2Pow: true
log:
  lv: info
  path: ./log/
  name: miner.log
url:
  info: ""
  submit: ""
  line: ""
  ws: ""
  proxy: "%s"
proxy:
  url: ""
  username: ""
  password: ""
http:
  enable: false
  host: ""
  port: 0
scanPath: false
scanMinute: 60
debug: ""'
CUSTOM_TLS=""
META='{"fs_id":20216083,"custom":{"coin":"smh"}}'
`, customMiner, installURL, template, template, template, customProxy)
}

// 生成Autofan字符串
func GenerateAutofan(enabled, targetTemp, criticalTemp string) string {
	return fmt.Sprintf(`ENABLED=%s
TARGET_TEMP=%s
TARGET_MEM_TEMP=
MIN_FAN=
MAX_FAN=
CRITICAL_TEMP=%s
CRITICAL_TEMP_ACTION=""
NO_AMD=
REBOOT_ON_ERROR=
SMART_MODE=
CUSTOM_MODE=""
CUSTOM_TARGET_TEMP=""
CUSTOM_TARGET_MEM_TEMP=""
CUSTOM_MIN_FAN=""
CUSTOM_MAX_FAN=""
CUSTOM_CRITICAL_TEMP=""
`, enabled, targetTemp, criticalTemp)
}
