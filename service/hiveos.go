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

	config := utils.GenerateHiveOsConfig(&miner.HiveOsConfig)
	wallet := utils.GenerateHiveOsWallet(&miner.HiveOsWallet)
	autofan := utils.GenerateHiveOsAutofan(&miner.HiveOsAutoFan)

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
	s.setMinerStats(ctx, rigID, &req)
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

	config := utils.GenerateHiveOsConfig(&miner.HiveOsConfig)
	wallet := utils.GenerateHiveOsWallet(&miner.HiveOsWallet)
	autofan := utils.GenerateHiveOsAutofan(&miner.HiveOsAutoFan)

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

	config := utils.GenerateHiveOsConfig(&miner.HiveOsConfig)
	wallet := utils.GenerateHiveOsWallet(&miner.HiveOsWallet)
	autofan := utils.GenerateHiveOsAutofan(&miner.HiveOsAutoFan)

	rigIDInt, err := strconv.Atoi(rigID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "string conversion", err.Error())
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

func (s *HiveOsService) setMinerStats(ctx context.Context, rigID string, req *dto.HiveosReq) error {
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

// // 设置 SetHiveOsConfig
// func (s *HiveOsService) SetHiveOsConfig(conf *utils.HiveOsConfig, from *utils.HiveOsConfig) {
// 	conf.HiveOsUrl = from.HiveOsUrl
// 	conf.ApiHiveOsUrls = from.ApiHiveOsUrls
// 	conf.RigID = from.RigID
// 	conf.RigPasswd = from.RigPasswd
// 	conf.WorkerName = from.WorkerName
// 	conf.FarmID = from.FarmID
// 	conf.Miner = from.Miner
// 	conf.Miner2 = from.Miner2
// 	conf.TimeZone = from.TimeZone
// 	conf.WdEnable = from.WdEnable
// 	conf.WdMiner = from.WdMiner
// 	conf.WdReboot = from.WdReboot
// 	conf.WdMaxLA = from.WdMaxLA
// 	conf.WdASR = from.WdASR
// 	conf.WdShareTime = from.WdShareTime
// 	conf.WdPowerEnabled = from.WdPowerEnabled
// 	conf.WdPowerMin = from.WdPowerMin
// 	conf.WdPowerMax = from.WdPowerMax
// 	conf.WdPowerAction = from.WdPowerAction
// 	conf.WdType = from.WdType
// }

// // 设置 SetHiveOsWallet
// func (s *HiveOsService) SetHiveOsWallet(wallet *utils.HiveOsWallet, from *utils.HiveOsWallet) {
// 	wallet.CustomMiner = from.CustomMiner
// 	wallet.CustomInstallURL = from.CustomInstallURL
// 	wallet.CustomAlgo = from.CustomAlgo
// 	wallet.CustomTemplate = from.CustomTemplate
// 	wallet.CustomUrl = from.CustomUrl
// 	wallet.CustomPass = from.CustomPass
// 	wallet.CustomUserConfig = from.CustomUserConfig
// 	wallet.CustomTLS = from.CustomTLS
// 	wallet.FsID = from.FsID
// 	wallet.Coin = from.Coin
// }

// // 设置 SetHiveOsAutoFan
// func (s *HiveOsService) SetHiveOsAutoFan(autofan *utils.HiveOsAutoFan, from *utils.HiveOsAutoFan) {
// 	autofan.CriticalTemp = from.CriticalTemp
// 	autofan.CriticalTempAction = from.CriticalTempAction
// 	autofan.Enable = from.Enable
// 	autofan.TargetTemp = from.TargetTemp
// 	autofan.MinFan = from.MinFan
// 	autofan.MaxFan = from.MaxFan
// 	autofan.NoAMD = from.NoAMD
// 	autofan.TargetMemTemp = from.TargetMemTemp
// 	autofan.RebootOnError = from.RebootOnError
// 	autofan.SmartMode = from.SmartMode
// }

// // 生成Config字符串
// // 主要是飞行表
// // 暂时去掉可选
// // DOH_ENABLED
// // SHELLINABOX_ENABLE
// // SSH_ENABLE
// // SSH_PASSWORD_ENABLE
// func (s *HiveOsService) generateConfig(
// 	hiveOsUrl string,
// 	apiHiveOsUrls string,
// 	rigID string,
// 	rigPasswd string,
// 	workerName string,
// 	farmID string,
// 	miner string,
// 	miner2 string,
// 	timeZone string,
// 	wdEnable string,
// 	wdMiner string,
// 	wdReboot string,
// 	wdMaxLA string,
// 	wdASR string,
// 	wdShareTime string,
// 	wdPowerEnabled string,
// 	wdPowerMin string,
// 	wdPowerMax string,
// 	wdPowerAction string,
// 	wdType string,
// ) string {
// 	return fmt.Sprintf(`\
// HIVE_HOST_URL="%s"
// API_HOST_URLs="%s"
// RIG_ID=%s
// RIG_PASSWD="%s"
// WORKER_NAME="%s"
// FARM_ID=%s
// MINER=%s
// MINER2=%s
// TIMEZONE="%s"
// # 算力监视器开关(开关)
// WD_ENABLED=%s
// # 软件重启于(分钟)
// WD_MINER=%s
// # 重启于(分钟)
// WD_REBOOT=%s
// # 重启当 LA >=
// WD_MAX_LA=%s
// # Min ASR
// WD_ASR=%s
// # WD Share Time(分钟)
// WD_SHARE_TIME=%s
// # WD Power
// WD_POWER_ENABLED=%s
// # WD Min Power
// WD_POWER_MIN=%s
// # WD Max Power
// WD_POWER_MAX=%s
// # Power Action (Reboot)
// WD_POWER_ACTION=%s
// # wd mode
// WD_TYPE='%s'
// # 未知参数
// WD_CHECK_GPU=%s
// WD_CHECK_CONN=%s
// WD_MINHASHES='%s'
// WD_MINHASHES_ALGO='%s'
// `,
// 		hiveOsUrl,
// 		apiHiveOsUrls,
// 		rigID,
// 		rigPasswd,
// 		workerName,
// 		farmID,
// 		miner,
// 		miner2,
// 		timeZone,
// 		wdEnable,
// 		wdMiner,
// 		wdReboot,
// 		wdMaxLA,
// 		wdASR,
// 		wdShareTime,
// 		wdPowerEnabled,
// 		wdPowerMin,
// 		wdPowerMax,
// 		wdPowerAction,
// 		wdType,
// 	)
// }

// // 生成Wallet字符串
// func (s *HiveOsService) generateWallet(
// 	customMiner string,
// 	customInstallURL string,
// 	customAlgo string,
// 	customTemplate string,
// 	customUrl string,
// 	customPass string,
// 	customUserConfig string,
// 	customTLS string,
// 	fsID string,
// 	coin string,
// ) string {
// 	return fmt.Sprintf(`
// # 软件名称
// CUSTOM_MINER="%s"
// # 下载地址
// CUSTOM_INSTALL_URL="%s"
// # 软件算法
// CUSTOM_ALGO="%s"
// # 模板
// CUSTOM_TEMPLATE="%s"
// # 池地址
// CUSTOM_URL="%s"
// # 密码
// CUSTOM_PASS="%s"
// # 其他模板参数
// CUSTOM_USER_CONFIG='%s'
// CUSTOM_TLS=""
// META='{
// 	"fs_id":%s,
// 	"custom": {
// 		"coin":"%s"
// 	}
// }'
// `,
// 		customMiner,
// 		customInstallURL,
// 		customAlgo,
// 		customTemplate,
// 		customUrl,
// 		customPass,
// 		customUserConfig,
// 	)
// }

// // 生成Autofan字符串
// // false 选项时，默认不填写
// // 使用默认参数时，不填写
// // 静态参数找不到
// func (s *HiveOsService) generateAutofan(
// 	criticalTemp string,
// 	criticalTempAction string,
// 	enable string,
// 	targetTemp string,
// 	minFan string,
// 	maxFan string,
// 	noAMD string,
// 	targetMemTemp string,
// 	rebootOnError string,
// 	smartMode string,
// ) string {
// 	return fmt.Sprintf(`
// # Critical temp
// CRITICAL_TEMP=%s
// # Critical action
// # 默认停止 重启 reboot 关闭 shutdown
// CRITICAL_TEMP_ACTION="%s"
// # 自动风扇
// ENABLED=%s
// # 缺少 Fan mode
// # 缺少 Static speed
// TARGET_TEMP=%s
// # Min fan speed
// MIN_FAN=%s
// # Max fan speed
// MAX_FAN=%s
// # 未知参数
// NO_AMD=%s
// # 缺少 Target core temp
// # Target memory temp
// TARGET_MEM_TEMP=%s
// # Reboot on errors 1 或 0
// REBOOT_ON_ERROR=%s
// # Smart mode 1 或 0
// SMART_MODE=%s
// `,
// 		criticalTemp,
// 		criticalTempAction,
// 		enable,
// 		targetTemp,
// 		minFan,
// 		maxFan,
// 		noAMD,
// 		targetMemTemp,
// 		rebootOnError,
// 		smartMode,
// 	)
// }
