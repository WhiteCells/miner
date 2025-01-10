package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"miner/common/dto"
	"miner/common/rsp"
	"miner/dao/mysql"
	"miner/dao/redis"
	"miner/model"
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
	taskDAO   *mysql.TaskDAO
}

func NewHiveOsService() *HiveOsService {
	return &HiveOsService{
		hiveOsRDB: redis.NewHiveOsRDB(),
		farmRDB:   redis.NewFarmRDB(),
		minerRDB:  redis.NewMinerRDB(),
		taskRDB:   redis.NewTaskRDB(),
		fsRDB:     redis.NewFsRDB(),
		taskDAO:   mysql.NewTaskDAO(),
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
	var req dto.HelloReq
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
	// s.setMinerStatus(ctx, rigID, &req)
	s.setMinerInfo(ctx, rigID, &req)
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
		log.Println("=============================================")
		log.Println(err.Error())
		log.Println("=============================================")
		return
	}

	taskIDInt, err := strconv.Atoi(task.ID)
	if err != nil {
		return
	}

	log.Println("===> stats  taskID", taskIDInt)

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
	task, err := s.taskDAO.GetTask(taskID)
	if err != nil {
		log.Println(taskID)
		log.Println("taskRDB.Get")
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}

	task.Result = req.Params.Payload
	task.Status = info.Done

	if err := s.taskDAO.UpdateTask(task); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "update result failed")
		return
	}

	log.Println("=======================", task)

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

func (s *HiveOsService) setMinerInfo(ctx context.Context, rigID string, req *dto.HelloReq) error {
	var info info.MinerInfo
	info.RigID = req.Params.RigID
	info.Passwd = req.Params.Passwd
	info.ServerUrl = req.Params.ServerUrl
	info.UID = req.Params.UID
	info.RefId = req.Params.RigID
	info.BootTime = req.Params.BootTime
	info.BootEvent = req.Params.BootEvent
	info.Ip = req.Params.Ip
	info.NetInterfaces = req.Params.NetInterfaces
	info.Openvpn = req.Params.Openvpn
	info.LanConfig = req.Params.LanConfig
	info.Gpu = req.Params.Gpu
	info.GpuCountAmd = req.Params.GpuCountAmd
	info.GpuCountNvidia = req.Params.GpuCountNvidia
	info.GpuCountIntel = req.Params.GpuCountIntel
	info.Mb = req.Params.Mb
	info.Cpu = req.Params.Cpu
	info.DiskModel = req.Params.DiskModel
	info.ImageVersion = req.Params.ImageVersion
	info.Kernel = req.Params.Kernel
	info.AmdVersion = req.Params.AmdVersion
	info.NvidiaVersion = req.Params.NvidiaVersion
	info.IntelVersion = req.Params.IntelVersion
	info.Version = req.Params.Version
	info.ShellinaboxEnable = req.Params.ShellinaboxEnable
	info.SshEnable = req.Params.SshEnable
	info.SshPasswordEnable = req.Params.SshPasswordEnable
	return s.hiveOsRDB.SetMinerInfo(ctx, rigID, &info)
}

func (s *HiveOsService) setMinerStatus(ctx context.Context, rigID string, req *dto.HiveosReq) error {
	var stats info.MinerStats
	stats.Algo = req.Params.MinerStats.Algo
	stats.BusNumbers = req.Params.MinerStats.BusNumbers
	stats.Coin = req.Params.Meta.Custom.Coin
	stats.Cpuavg = req.Params.Cpuavg
	stats.Cputemp = req.Params.Cputemp
	stats.Df = req.Params.Df
	stats.Fan = req.Params.Fan
	stats.FsID = req.Params.Meta.FsID
	stats.Hs = req.Params.MinerStats.Hs
	stats.HsUnits = req.Params.MinerStats.HsUnits
	stats.Khs = req.Params.MinerStats.Khs
	stats.Mem = req.Params.Mem
	stats.Miner = req.Params.Miner
	stats.Power = req.Params.Power
	stats.Status = req.Params.MinerStats.Status
	stats.Temp = req.Params.Temp
	stats.TotalKhs = req.Params.TotalKhs
	return s.hiveOsRDB.SetMinerStats(ctx, rigID, &stats)
}

func (s *HiveOsService) PostTask(ctx context.Context, req *dto.PostTaskReq) (string, error) {
	// TODO 限制命令长度
	// TODO 限制 QPS
	task := &model.Task{
		Type:    req.Type,
		Status:  info.Pending,
		Content: req.Content,
	}
	if err := s.taskDAO.AddTask(ctx, req.RigID, task); err != nil {
		return "", err
	}

	return task.ID, nil
}

func (s *HiveOsService) GetTaskRes(ctx context.Context, taskID string) (*model.Task, error) {
	return s.taskDAO.GetTask(taskID)
}

func (s *HiveOsService) GetTaskStats(ctx context.Context, taskID string) (info.TaskStatus, error) {
	task, err := s.GetTaskRes(ctx, taskID)
	return task.Status, err
}

func (s *HiveOsService) GetMinerStats(ctx context.Context, rigID string) (*info.MinerStats, error) {
	return s.hiveOsRDB.GetMinerStatus(ctx, rigID)
}

func (s *HiveOsService) GetMinerInfo(ctx context.Context, rigID string) (*info.MinerInfo, error) {
	return s.hiveOsRDB.GetMinerInfo(ctx, rigID)
}

// 生成Config字符串
func (s *HiveOsService) generateConfig(rigID int, passwrod string, farmID int, workerName, hiveHostURL, apiHostURL, sshServer string) string {
	return fmt.Sprintf(`\
HIVE_HOST_URL="%s"
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
	return fmt.Sprintf(`
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
