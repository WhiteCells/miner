package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"miner/common/dto"
	"miner/common/perm"
	"miner/common/rsp"
	"miner/dao/mysql"
	"miner/dao/mysql/relationdao"
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
	userFarmDAO  *relationdao.UserFarmDAO
	farmMinerDAO *relationdao.FarmMinerDAO
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
		userFarmDAO:  relationdao.NewUserFarmDAO(),
		farmMinerDAO: relationdao.NewFarmMinerDAO(),
		minerService: *NewMinerService(),
	}
}

func (m *HiveosService) Poll(ctx *gin.Context) {
	rigIDStr := ctx.Query("id_rig")
	method := ctx.Query("method")

	switch method {
	case "hello":
		if rigIDStr != "" {
			m.helloCase(ctx, rigIDStr)
		} else {
			m.helloCaseUseHash(ctx)
		}
	case "stats":
		m.statsCase(ctx, rigIDStr)
	case "message":
		m.messageCase(ctx, rigIDStr)
	}
}

func (m *HiveosService) helloCase(ctx *gin.Context, rigIDStr string) {
	var req dto.HelloReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Logger.Error("helloCase ShouldBindJSON, error" + err.Error())
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	m.formatOutput(&req)

	rigID, err := strconv.Atoi(rigIDStr)
	if err != nil {
		utils.Logger.Error("helloCase rigIDStr conv error" + err.Error())
		log.Println("helloCase rigIDStr conv error", rigID, err.Error())
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// 获取 miner 缓存
	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, rigID)
	if err != nil {
		utils.Logger.Error("helloCase GetMinerByRigID" + err.Error())
		log.Println("helloCase minerRDB.GetMinerByRigID", rigID, err.Error())
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// 验证密码
	if req.Params.Passwd != minerInfo.HiveOsConfig.RigPasswd {
		log.Println(req.Params.Passwd, minerInfo.HiveOsConfig.RigPasswd, "req.Params.Passwd")
		rsp.Error(ctx, http.StatusInternalServerError, "invalid pass", nil)
		return
	}

	// 对 req 的数据进行存储
	if err := m.setMinerInfo(ctx, rigID, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to set hiveos info", nil)
		return
	}

	// 更新 miner GpuNum
	minerInfo.GpuNum = len(req.Params.Gpu)
	if err := m.minerRDB.UpdateMinerByRigID(ctx, rigID, minerInfo); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to update miner info", nil)
		return
	}

	rsp := &dto.ServerRsp{
		ID:      0,
		Jsonrpc: "2.0",
		Result: dto.ServerRsp_Result{
			ID:        99999,
			Config:    utils.GenerateHiveOsConfig(&minerInfo.HiveOsConfig),
			Wallet:    utils.GenerateHiveOsWallet(&minerInfo.HiveOsWallet),
			Autofan:   utils.GenerateHiveOsAutofan(&minerInfo.HiveOsAutoFan),
			Justwrite: 1,
			Confseq:   1,
		},
	}

	// m.formatOutput(&rsp)

	ctx.JSON(http.StatusOK, rsp)
}

// poll hello case use hash
func (m *HiveosService) helloCaseUseHash(ctx *gin.Context) {
	var req dto.HelloReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Logger.Error("helloCase ShouldBindJSON, error" + err.Error())
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	m.formatOutput(&req)

	// 获取 farm
	farm, err := m.farmDAO.GetFarmByHash(ctx, req.Params.FarmHash)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	// 创建 miner
	pass, err := utils.GenerateRigPass(8)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	miner := &model.Miner{
		Name: req.Params.WorkerName,
		Pass: pass,
	}
	if err = m.minerDAO.CreateMiner(ctx, farm.ID, miner); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	// 创建 miner 缓存
	minerInfo := &info.Miner{
		HiveOsConfig: utils.HiveOsConfig{
			HiveOsUrl:     utils.GenerateHiveOsUrl(),
			ApiHiveOsUrls: utils.GenerateHiveOsUrl(),
			WorkerName:    req.Params.Mb.Bios,
			FarmID:        farm.ID,
			RigID:         miner.ID,
			RigPasswd:     pass,
		},
	}
	if err := m.minerRDB.CreateMinerByRigID(ctx, miner.ID, minerInfo); err != nil {
		utils.Logger.Error("helloCaseUseHash CreateMinerByRigID error" + err.Error())
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}

	// 对 req 的数据进行存储
	if err := m.setMinerInfo(ctx, miner.ID, &req); err != nil {
		utils.Logger.Error("helloCaseUseHash setMinerInfo error" + err.Error())
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}

	rsp := &dto.ServerHashRsp{
		Jsonrpc: "2.0",
		ID:      0,
		Result: dto.ServerHashRsp_Result{
			RigName:         miner.Name,
			RespositoryList: "",
			Config:          utils.GenerateHiveOsConfig(&minerInfo.HiveOsConfig),
			Wallet:          utils.GenerateHiveOsWallet(&minerInfo.HiveOsWallet),
			NvidiaOc:        "",
			Autofan:         utils.GenerateHiveOsAutofan(&minerInfo.HiveOsAutoFan),
			Confseq:         1,
		},
	}

	// m.formatOutput(&rsp)

	ctx.JSON(http.StatusOK, rsp)
}

// Poll stats case
func (m *HiveosService) statsCase(ctx *gin.Context, rigIDStr string) {
	var req dto.HiveOsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	m.formatOutput(&req)

	rigID, err := strconv.Atoi(rigIDStr)
	if err != nil {
		utils.Logger.Error("statsCase rigIDStr conv error" + err.Error())
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// minerInfo
	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, rigID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// 验证密码
	if req.Params.Passwd != minerInfo.HiveOsConfig.RigPasswd {
		log.Println("stats case valid password error", req.Params.Passwd, minerInfo.HiveOsConfig.RigPasswd)
		rsp.Error(ctx, http.StatusInternalServerError, "invalid pass", "")
		return
	}

	// 对 req 的数据进行存储
	if err := m.setMinerStats(ctx, rigID, &req); err != nil {
		utils.Logger.Error("statsCase setMinerStats error" + err.Error())
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}

	// 从 taskRDB 中拿出对应的 task
	task, err := m.taskRDB.GetTask(ctx, rigID)
	if err != nil {
		log.Println("=======================")
		log.Println("rigID:", rigID, err.Error())
		log.Println("=======================")
		return
	}

	log.Println("===> stats  taskID", task.ID)

	// 对任务分类讨论
	switch task.Type {
	case info.Cmd:
		ctx.JSON(http.StatusOK, &dto.ServerRsp{
			ID:      rigID,
			Jsonrpc: "2.0",
			Result: dto.ServerRsp_Result{
				ID:        task.ID,
				Config:    utils.GenerateHiveOsConfig(&minerInfo.HiveOsConfig),
				Wallet:    utils.GenerateHiveOsWallet(&minerInfo.HiveOsWallet),
				Autofan:   utils.GenerateHiveOsAutofan(&minerInfo.HiveOsAutoFan),
				Justwrite: 1,
				Command:   "exec",
				Exec:      task.Content,
				Confseq:   1,
			},
		})

	case info.Config:
		ctx.JSON(http.StatusOK, &dto.ServerRsp{
			ID:      rigID,
			Jsonrpc: "2.0",
			Result: dto.ServerRsp_Result{
				ID:        task.ID,
				Config:    utils.GenerateHiveOsConfig(&minerInfo.HiveOsConfig),
				Wallet:    utils.GenerateHiveOsWallet(&minerInfo.HiveOsWallet),
				Autofan:   utils.GenerateHiveOsAutofan(&minerInfo.HiveOsAutoFan),
				Justwrite: 1,
				Command:   "config",
				Confseq:   1,
			},
		})
	}
}

// Poll message case
func (m *HiveosService) messageCase(ctx *gin.Context, rigIDStr string) {
	// 这一次 hiveos 的请求为上一次 服务器 回包（命令、配置）的结果
	var req dto.HiveOsResReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	m.formatOutput(&req)

	rigID, err := strconv.Atoi(rigIDStr)
	if err != nil {
		utils.Logger.Error("messageCase rigIDStr conv error" + err.Error())
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// 通过 rigID 获取 miner
	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, rigID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	// 验证密码
	if req.Params.Passwd != minerInfo.HiveOsConfig.RigPasswd {
		log.Println(req.Params.Passwd, minerInfo.HiveOsConfig.RigPasswd)
		log.Println("req.Params.Passwd")
		rsp.Error(ctx, http.StatusInternalServerError, "invalid pass", "")
		return
	}

	// 更新任务
	taskID := req.Params.ID
	updateInfo := map[string]any{
		"result": req.Params.Payload,
		"status": info.Done,
	}
	taskIDInt, err := strconv.Atoi(taskID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed strconv", nil)
		return
	}
	if err := m.taskDAO.UpdateTask(ctx, taskIDInt, updateInfo); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "update result failed")
		return
	}

	log.Println("=======================")
	log.Println("task:", req.Params.Payload)
	log.Println("=======================")

	ctx.JSON(http.StatusOK, &dto.ServerRsp{
		ID:      0,
		Jsonrpc: "2.0",
		Result: dto.ServerRsp_Result{
			ID:        0,
			Config:    utils.GenerateHiveOsConfig(&minerInfo.HiveOsConfig),
			Wallet:    utils.GenerateHiveOsWallet(&minerInfo.HiveOsWallet),
			Autofan:   utils.GenerateHiveOsAutofan(&minerInfo.HiveOsAutoFan),
			Justwrite: 1,
			Confseq:   1,
		},
	})
}

func (m *HiveosService) setMinerInfo(ctx context.Context, rigID int, req *dto.HelloReq) error {
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

func (m *HiveosService) setMinerStats(ctx context.Context, rigID int, req *dto.HiveOsReq) error {
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

func (m *HiveosService) PostTask(ctx context.Context, userID int, req *dto.PostTaskReq) (int, error) {
	task := &model.Task{
		Type:    req.Type,
		Status:  info.Pending,
		Content: req.Content,
	}

	// 检查 user 对 farm 权限
	p, err := m.userFarmDAO.GetPerm(ctx, userID, req.FarmID)
	if err != nil || (p != perm.FarmManager && p != perm.FarmOwner) {
		return -1, err
	}
	// farm 与 miner 关联
	if err := m.farmMinerDAO.ExistMiner(ctx, req.FarmID, req.MinerID); err != nil {
		return -1, err
	}

	if err := m.taskDAO.AddTask(ctx, userID, req.FarmID, req.MinerID, task); err != nil {
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
