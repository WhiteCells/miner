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

// 轮询
func (s *HiveOsService) Poll(ctx *gin.Context) {
	rigID := ctx.Query("id_rig")
	method := ctx.Query("method")
	switch method {
	case "hello":
		s.helloCase(ctx, rigID)
	case "stats":
		s.statsCase(ctx, rigID)
	case "message":
		s.messageCase(ctx, rigID)
	}
}

// Poll hello case
func (s *HiveOsService) helloCase(ctx *gin.Context, rigID string) {
	var req dto.HiveosReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	////////////////////////////////////////////////
	jsonInd, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return
	}
	fmt.Printf("%s\n", jsonInd)
	////////////////////////////////////////////////
	// 对 req 的数据进行存储
	s.setMinerStatus(ctx, rigID, &req)
	// 从 req 中获取 rigID，根据 rigID 查询 hiveOsRDB farmID:minerID
	// rigID := req.Params.RigID
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
	miner, err := s.minerRDB.GetByID(ctx, farmID, minerID)
	if err != nil {
		log.Println(farmID, minerID)
		log.Println("minerRDB.GetByID")
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	// 验证密码
	if req.Params.Passwd != miner.Pass {
		log.Println(req.Params.Passwd, miner.Pass)
		log.Println("req.Params.Passwd")
		rsp.Error(ctx, http.StatusInternalServerError, "invalid pass", "")
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
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	farmIDInt, err := strconv.Atoi(farmID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	//////////////////////////////
	// config
	//////////////////////////////
	host := utils.Config.Server.Host
	port := utils.Config.Server.Port
	hive_host := fmt.Sprintf("http://%s:%d/hiveos", host, port)
	config := s.generateConfig(rigIDInt, req.Params.Passwd, farmIDInt, miner.Name, hive_host, hive_host, "")
	//////////////////////////////
	// wallet
	//////////////////////////////
	wallet := s.generateWallet("custom", "", "template", "custom")
	//////////////////////////////
	// autofan
	//////////////////////////////
	autofan := s.generateAutofan("", "", "")

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
			Config:    config,
			Wallet:    wallet,
			Autofan:   autofan,
			Justwrite: 1,
			Confseq:   1,
		},
	})
}

// Poll stats case
func (s *HiveOsService) statsCase(ctx *gin.Context, rigID string) {
	var req dto.HiveosReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	////////////////////////////////////////////////
	jsonInd, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return
	}
	fmt.Printf("%s\n", jsonInd)
	////////////////////////////////////////////////
	// 对 req 的数据进行存储
	s.setMinerStatus(ctx, rigID, &req)
	// 从 req 中获取 rigID，根据 rigID 查询 hiveOsRDB farmID:minerID
	// rigID := req.Params.RigID
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
	miner, err := s.minerRDB.GetByID(ctx, farmID, minerID)
	if err != nil {
		log.Println(farmID, minerID)
		log.Println("minerRDB.GetByID")
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	// 验证密码
	if req.Params.Passwd != miner.Pass {
		log.Println(req.Params.Passwd, miner.Pass)
		log.Println("req.Params.Passwd")
		rsp.Error(ctx, http.StatusInternalServerError, "invalid pass", "")
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
		rsp.Error(ctx, http.StatusInternalServerError, "convertion failed", err.Error())
		return
	}
	farmIDInt, err := strconv.Atoi(farmID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "convertion failed", err.Error())
		return
	}
	//////////////////////////////
	// config
	//////////////////////////////
	host := utils.Config.Server.Host
	port := utils.Config.Server.Port
	hive_host := fmt.Sprintf("http://%s:%d/hiveos", host, port)
	config := s.generateConfig(rigIDInt, req.Params.Passwd, farmIDInt, miner.Name, hive_host, hive_host, "")
	//////////////////////////////
	// wallet
	//////////////////////////////
	wallet := s.generateWallet("custom", "", "template", "custom")
	//////////////////////////////
	// autofan
	//////////////////////////////
	autofan := s.generateAutofan("", "", "")

	// 从 taskRDB 中拿出对应的 task
	task, err := s.taskRDB.GetTask(ctx, rigID)
	if err != nil {
		// 没有任务则结束
		return
	}

	taskIDInt, err := strconv.Atoi(task.ID)
	if err != nil {
		return
	}

	// 对任务分类讨论
	switch task.Type {
	case info.Cmd:
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
				ID:        taskIDInt,
				Config:    config,
				Wallet:    wallet,
				Autofan:   autofan,
				Justwrite: 1,
				Command:   "exec",
				Exec:      task.Content,
				Confseq:   1,
			},
		})

	case info.Config:
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
				ID:        taskIDInt,
				Config:    config,
				Wallet:    wallet,
				Autofan:   autofan,
				Justwrite: 1,
				Command:   "config",
				Confseq:   1,
			},
		})
	}
}

// Poll message case
func (s *HiveOsService) messageCase(ctx *gin.Context, rigID string) {
	// 这一次 hiveos 的请求为上一次 服务器 回包（命令、配置）的结果
	var req dto.HiveosResReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	////////////////////////////////////////////////
	jsonInd, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return
	}
	fmt.Printf("%s\n", jsonInd)
	////////////////////////////////////////////////

	// 查找命令
	// 根绝请求生成新的任务，更新任务中的 result
	taskID := req.Params.ID
	task, err := s.taskRDB.Get(ctx, taskID)
	if err != nil {
		log.Println(taskID)
		log.Println("taskRDB.Get")
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}

	task.Result = req.Params.Payload
	task.Status = info.Done

	log.Fatalln(task.Content)

	s.taskRDB.Set(ctx, taskID, task)

	// 从 req 中获取 rigID，根据 rigID 查询 hiveOsRDB farmID:minerID
	// rigID := req.Params.RigID
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
	miner, err := s.minerRDB.GetByID(ctx, farmID, minerID)
	if err != nil {
		log.Println(farmID, minerID)
		log.Println("minerRDB.GetByID")
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	// 验证密码
	if req.Params.Passwd != miner.Pass {
		log.Println(req.Params.Passwd, miner.Pass)
		log.Println("req.Params.Passwd")
		rsp.Error(ctx, http.StatusInternalServerError, "invalid pass", "")
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
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	farmIDInt, err := strconv.Atoi(farmID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	//////////////////////////////
	// config
	//////////////////////////////
	host := utils.Config.Server.Host
	port := utils.Config.Server.Port
	hive_host := fmt.Sprintf("http://%s:%d/hiveos", host, port)
	config := s.generateConfig(rigIDInt, req.Params.Passwd, farmIDInt, miner.Name, hive_host, hive_host, "")
	//////////////////////////////
	// wallet
	//////////////////////////////
	wallet := s.generateWallet("custom", "", "template", "custom")
	//////////////////////////////
	// autofan
	//////////////////////////////
	autofan := s.generateAutofan("", "", "")

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
			ID:        000,
			Config:    config,
			Wallet:    wallet,
			Autofan:   autofan,
			Justwrite: 1,
			Confseq:   1,
		},
	})
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

func (s *HiveOsService) PostTask(ctx context.Context, req *dto.PostTaskReq) (string, error) {
	// 生成任务 ID
	taskID, err := utils.GenerateUID()
	if err != nil {
		return "", err
	}
	task := &info.Task{
		ID:      taskID,
		Type:    req.Type,
		Status:  info.Pending,
		Content: req.Content,
	}
	if err := s.taskRDB.AddTask(ctx, req.RigID, taskID, task); err != nil {
		return "", err
	}
	return taskID, nil
}

func (s *HiveOsService) GetTaskRes(ctx context.Context, taskID string) (*info.Task, error) {
	return s.taskRDB.Get(ctx, taskID)
}

func (s *HiveOsService) GetStats(ctx context.Context) error {

	return nil
}

// 生成Config字符串
func (s *HiveOsService) generateConfig(rigID int, passwrod string, farmID int, workerName, hiveHostURL, apiHostURL, sshServer string) string {
	return fmt.Sprintf(`HIVE_HOST_URL="%s"
API_HOST_URLs="%s"
RIG_ID=%d
RIG_PASSWD="%s"
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
`, hiveHostURL, apiHostURL, rigID, passwrod, workerName, farmID, sshServer)
}

// 生成Wallet字符串
func (s *HiveOsService) generateWallet(customMiner, installURL, template, customProxy string) string {
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
func (s *HiveOsService) generateAutofan(enabled, targetTemp, criticalTemp string) string {
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
