package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

	"github.com/gin-gonic/gin"
)

type HiveosService struct {
	minerDAO     *mysql.MinerDAO
	hiveosRDB    *redis.HiveOsRDB
	farmDAO      *mysql.FarmDAO
	farmRDB      *redis.FarmRDB
	minerRDB     *redis.MinerRDB
	taskRDB      *redis.TaskRDB
	taskDAO      *mysql.TaskDAO
	minerService MinerService
}

func NewHiveosService() *HiveosService {
	return &HiveosService{
		minerDAO:     mysql.NewMinerDAO(),
		hiveosRDB:    redis.NewHiveOsRDB(),
		farmDAO:      mysql.NewFarmDAO(),
		farmRDB:      redis.NewFarmRDB(),
		minerRDB:     redis.NewMinerRDB(),
		taskRDB:      redis.NewTaskRDB(),
		taskDAO:      mysql.NewTaskDAO(),
		minerService: *NewMinerService(),
	}
}

func (m *HiveosService) Poll(ctx *gin.Context) {
	rigID := ctx.Query("id_rig")
	method := ctx.Query("method")

	switch method {
	case "hello":
		if rigID != "" {
			m.helloCase(ctx, rigID)
		} else {
			m.helloCaseUseHash(ctx)
		}
	case "stats":
		m.statsCase(ctx, rigID)
	case "message":
		m.messageCase(ctx, rigID)
	}
}

func (m *HiveosService) helloCase(ctx *gin.Context, rigID string) {
	var req dto.HelloReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Logger.Error("helloCase ShouldBindJSON, error" + err.Error())
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	m.formatOutput(&req)

	miner, err := m.minerRDB.GetMinerByRigID(ctx, rigID)
	if err != nil {
		log.Println(rigID, "minerRDB.GetMinerByRigID")
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// 验证密码
	if req.Params.Passwd != miner.HiveOsConfig.RigPasswd {
		log.Println(req.Params.Passwd, miner.HiveOsConfig.RigPasswd, "req.Params.Passwd")
		rsp.Error(ctx, http.StatusInternalServerError, "invalid pass", "")
		return
	}

	// 对 req 的数据进行存储
	errSetInfo := m.setMinerInfo(ctx, rigID, &req)
	if errSetInfo != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to set hiveos info", "")
		return
	}

	// 更新 miner GpuNum
	miner.GpuNum = len(req.Params.Gpu)
	if err := m.minerRDB.UpdateMinerByRigID(ctx, rigID, miner); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to update miner info", "")
		return
	}

	config := utils.GenerateHiveOsConfig(&miner.HiveOsConfig)
	wallet := utils.GenerateHiveOsWallet(&miner.HiveOsWallet)
	autofan := utils.GenerateHiveOsAutofan(&miner.HiveOsAutoFan)

	rsp_ := &dto.ServerRsp{
		ID:      0,
		Jsonrpc: "2.0",
		Result: dto.ServerRsp_Result{
			ID:        99999,
			Config:    config,
			Wallet:    wallet,
			Autofan:   autofan,
			Justwrite: 1,
			Confseq:   1,
		},
	}

	//log.Println("rsp", rsp_)

	ctx.JSON(http.StatusOK, rsp_)
}

// poll hello case use hash
func (m *HiveosService) helloCaseUseHash(ctx *gin.Context) {
	var req dto.HelloReq
	body, err := io.ReadAll(ctx.Request.Body)
	err = json.Unmarshal(body, &req)
	if err != nil {
		utils.Logger.Error("helloCaseUseHash Unmarshal req error:" + err.Error())
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}

	// 获取 farm
	farm, err := m.farmDAO.GetFarmByHash(ctx, req.Params.FarmHash)
	// 创建 miner
	rigID, err := m.minerService.generateRigID(ctx, 8)
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	pass, err := m.minerService.generateRigPass(ctx, 8)
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	miner := &model.Miner{
		Name:  req.Params.WorkerName,
		RigID: rigID,
		Pass:  pass,
	}
	err = m.minerDAO.CreateMiner(ctx, farm.ID, miner)

	// 创建 miner 缓存
	minerInfo := &info.Miner{}
	if err := m.minerRDB.CreateMinerByRigID(ctx, miner.RigID, minerInfo); err != nil {
		utils.Logger.Error("helloCaseUseHash CreateMinerByRigID error" + err.Error())
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}

	if err := m.setMinerInfo(ctx, rigID, &req); err != nil {
		utils.Logger.Error("helloCaseUseHash setMinerInfo error" + err.Error())
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}

	// farm, err := m.farmDAO.GetFarmByHash(ctx, req.Params.FarmHash)

	// createMinerReq := dto.CreateMinerReq{
	// 	Name:   req.Params.WorkerName,
	// 	FarmID: farm.ID,
	// }
	// // 这里会有一个逻辑问题
	// // userID 应该是 farm 的管理员的 userID 还是使用 hash 的 userID
	// m.minerService.CreateMiner(ctx, userID, farm.ID, &createMinerReq)

	config := utils.GenerateHiveOsConfig(&minerInfo.HiveOsConfig)
	wallet := utils.GenerateHiveOsWallet(&minerInfo.HiveOsWallet)
	autofan := utils.GenerateHiveOsAutofan(&minerInfo.HiveOsAutoFan)

	rsp := &dto.ServerHashRsp{
		Jsonrpc: "2.0",
		ID:      0,
		Result: dto.ServerHashRsp_Result{
			RigName:         miner.Name,
			RespositoryList: "",
			Config:          config,
			Wallet:          wallet,
			NvidiaOc:        "",
			Autofan:         autofan,
			Confseq:         1,
		},
	}

	m.formatOutput(&rsp)

	ctx.JSON(http.StatusOK, rsp)
}

// Poll stats case
func (m *HiveosService) statsCase(ctx *gin.Context, rigID string) {
	var req dto.HiveOsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	m.formatOutput(&req)

	// 通过 rigID 获取 miner
	miner, err := m.minerDAO.GetMinerByRigID(ctx, rigID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// 从 req 中获取 rigID，根据 rigID 查询 hiveOsRDB farmID:minerID
	// _, farmID, minerID, err := m.hiveosDAO.GetRigFarmAndMinerID(ctx, rigID)
	// if err != nil {
	// 	rsp.Error(ctx, http.StatusNotAcceptable, err.Error(), "")
	// 	return
	// }
	// // 通过 farmID 和 minerID 获取 miner
	// miner, err := m.minerRDB.GetByID(ctx, farmID, minerID)
	// if err != nil {
	// 	log.Println(farmID, minerID, "minerRDB.GetByID")
	// 	rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
	// 	return
	// }
	// 验证密码
	if req.Params.Passwd != miner.Pass {
		log.Println("stats case valid password error", req.Params.Passwd, miner.Pass)
		rsp.Error(ctx, http.StatusInternalServerError, "invalid pass", "")
		return
	}

	//
	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, rigID)

	// 对 req 的数据进行存储
	m.setMinerStats(ctx, rigID, &req)

	// 通过 minerID 获取其 fsID

	rigIDInt, err := strconv.Atoi(rigID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "convertion failed", err.Error())
		return
	}

	config := utils.GenerateHiveOsConfig(&minerInfo.HiveOsConfig)
	wallet := utils.GenerateHiveOsWallet(&minerInfo.HiveOsWallet)
	autofan := utils.GenerateHiveOsAutofan(&minerInfo.HiveOsAutoFan)

	// 从 taskRDB 中拿出对应的 task
	task, err := m.taskRDB.GetTask(ctx, rigID)
	if err != nil {
		// 没有任务则结束
		log.Println("=============================================")
		log.Println(rigID, err.Error())
		log.Println("=============================================")
		return
	}

	log.Println("===> stats  taskID", task.ID)

	// 对任务分类讨论
	switch task.Type {
	case info.Cmd:
		ctx.JSON(http.StatusOK, &dto.ServerRsp{
			ID:      rigIDInt,
			Jsonrpc: "2.0",
			Result: dto.ServerRsp_Result{
				ID:        task.ID,
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
			Result: dto.ServerRsp_Result{
				ID:        task.ID,
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
func (m *HiveosService) messageCase(ctx *gin.Context, rigID string) {
	// 这一次 hiveos 的请求为上一次 服务器 回包（命令、配置）的结果
	var req dto.HiveOsResReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	m.formatOutput(&req)

	// 查找命令
	// 根绝请求生成新的任务，更新任务中的 result
	taskID := req.Params.ID
	task, err := m.taskDAO.GetTask(ctx, taskID)
	if err != nil {
		log.Println(taskID)
		log.Println("taskRDB.Get")
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}

	task.Result = req.Params.Payload
	task.Status = info.Done

	if err := m.taskDAO.UpdateTask(ctx, task); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "update result failed")
		return
	}

	log.Println("=======================")
	log.Println("task:", task)
	log.Println("=======================")

	_, farmID, minerID, err := m.hiveosDAO.GetRigFarmAndMinerID(ctx, rigID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "hiveOsRDB.GetRigFarmAndMinerID")
		return
	}
	// 通过 farmID 和 minerID 获取 miner
	miner, err := m.minerRDB.GetByID(ctx, farmID, minerID)
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

	config := utils.GenerateHiveOsConfig(&miner.HiveOsConfig)
	wallet := utils.GenerateHiveOsWallet(&miner.HiveOsWallet)
	autofan := utils.GenerateHiveOsAutofan(&miner.HiveOsAutoFan)

	ctx.JSON(http.StatusOK, &dto.ServerRsp{
		ID:      0,
		Jsonrpc: "2.0",
		Result: dto.ServerRsp_Result{
			ID:        0,
			Config:    config,
			Wallet:    wallet,
			Autofan:   autofan,
			Justwrite: 1,
			Confseq:   1,
		},
	})
}

func (m *HiveosService) setMinerInfo(ctx context.Context, rigID string, req *dto.HelloReq) error {
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
	return m.hiveosRDB.SetMinerInfo(ctx, rigID, &info)
}

func (m *HiveosService) setMinerStats(ctx context.Context, rigID string, req *dto.HiveOsReq) error {
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
	return m.hiveosRDB.SetMinerStats(ctx, rigID, &stats)
}

func (m *HiveosService) PostTask(ctx context.Context, req *dto.PostTaskReq) (int, error) {
	task := &model.Task{
		Type:    req.Type,
		Status:  info.Pending,
		Content: req.Content,
	}
	if err := m.taskDAO.AddTask(ctx, req.RigID, task); err != nil {
		return -1, err
	}

	return task.ID, nil
}

func (m *HiveosService) GetTaskRes(ctx context.Context, taskID string) (*model.Task, error) {
	return m.taskDAO.GetTask(ctx, taskID)
}

func (m *HiveosService) GetTaskStats(ctx context.Context, taskID string) (info.TaskStatus, error) {
	task, err := m.GetTaskRes(ctx, taskID)
	return task.Status, err
}

func (m *HiveosService) GetMinerStats(ctx context.Context, rigID string) (*info.MinerStats, error) {
	return m.hiveosRDB.GetMinerStats(ctx, rigID)
}

func (m *HiveosService) GetMinerInfo(ctx context.Context, rigID string) (*info.MinerInfo, error) {
	return m.hiveosRDB.GetMinerInfo(ctx, rigID)
}

// 格式化输出
func (HiveosService) formatOutput(req any) {
	jsonInd, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		fmt.Println("json marshal indent error:", err)
		return
	}
	fmt.Printf("%s\n", jsonInd)
}
