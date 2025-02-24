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

	"github.com/gin-gonic/gin"
)

type HiveOsService struct {
	hiveOsRDB    *redis.HiveOsRDB
	farmRDB      *redis.FarmRDB
	minerRDB     *redis.MinerRDB
	taskRDB      *redis.TaskRDB
	fsRDB        *redis.FsRDB
	taskDAO      *mysql.TaskDAO
	minerService *MinerService
}

func NewHiveOsService() *HiveOsService {
	return &HiveOsService{
		hiveOsRDB:    redis.NewHiveOsRDB(),
		farmRDB:      redis.NewFarmRDB(),
		minerRDB:     redis.NewMinerRDB(),
		taskRDB:      redis.NewTaskRDB(),
		fsRDB:        redis.NewFsRDB(),
		taskDAO:      mysql.NewTaskDAO(),
		minerService: NewMinerService(),
	}
}

// 轮询
func (s *HiveOsService) Poll(ctx *gin.Context) {
	rigID := ctx.Query("id_rig")
	method := ctx.Query("method")

	switch method {
	case "hello":
		if rigID != "" {
			s.helloCase(ctx, rigID)
		} else {
			s.helloCaseUseHash(ctx)
		}
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
	s.formatOutput(&req)

	// 从 req 中获取 rigID，根据 rigID 查询 hiveOsRDB farmID:minerID
	userID, farmID, minerID, err := s.hiveOsRDB.GetRigFarmAndMinerID(ctx, rigID)
	if err != nil {
		log.Println(rigID, "hiveOsRDB.GetRigFarmAndMinerID")
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	// 通过 farmID 和 minerID 获取 miner
	miner, err := s.minerRDB.GetByID(ctx, farmID, minerID)
	if err != nil {
		log.Println(farmID, minerID, "minerRDB.GetByID")
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	// 验证密码
	if req.Params.Passwd != miner.Pass {
		log.Println(req.Params.Passwd, miner.Pass, "req.Params.Passwd")
		rsp.Error(ctx, http.StatusInternalServerError, "invalid pass", "")
		return
	}

	// 对 req 的数据进行存储
	s.setMinerInfo(ctx, rigID, &req)

	// 更新 miner GpuNum
	miner.GpuNum = len(req.Params.Gpu)

	// 更新 farm GpuNum
	farm, err := s.farmRDB.GetByID(ctx, userID, farmID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "get farm failed", "")
		return
	}
	farm.GpuNum += miner.GpuNum
	s.farmRDB.Set(ctx, userID, farm, farm.Perm)

	config := utils.GenerateHiveOsConfig(&miner.HiveOsConfig)
	wallet := utils.GenerateHiveOsWallet(&miner.HiveOsWallet)
	autofan := utils.GenerateHiveOsAutofan(&miner.HiveOsAutoFan)

	rsp := &dto.ServerRsp{
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

	log.Println("rsp", rsp)

	ctx.JSON(http.StatusOK, rsp)
}

// poll hello case use hash
func (s *HiveOsService) helloCaseUseHash(ctx *gin.Context) {
	var req dto.HelloReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	s.formatOutput(&req)

	// 如果用户在已连接的情况下，再次使用 hash 连接，此时已经存在连接
	// 第二次请求不会携带上一次使用的 rig_id 及 pass

	userID, farmID, err := s.farmRDB.GetUserAndFarmIDByHash(ctx, req.Params.FarmHash)
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	miner, err := s.minerService.CreateMinerByUserID(ctx, userID, farmID, req.Params.Mb.Bios)
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}

	config := utils.GenerateHiveOsConfig(&miner.HiveOsConfig)
	wallet := utils.GenerateHiveOsWallet(&miner.HiveOsWallet)
	autofan := utils.GenerateHiveOsAutofan(&miner.HiveOsAutoFan)

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

	s.formatOutput(&rsp)

	ctx.JSON(http.StatusOK, rsp)
}

// Poll stats case
func (s *HiveOsService) statsCase(ctx *gin.Context, rigID string) {
	var req dto.HiveOsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	s.formatOutput(&req)

	// 从 req 中获取 rigID，根据 rigID 查询 hiveOsRDB farmID:minerID
	_, farmID, minerID, err := s.hiveOsRDB.GetRigFarmAndMinerID(ctx, rigID)
	if err != nil {
		rsp.Error(ctx, http.StatusNotAcceptable, err.Error(), "")
		return
	}
	// 通过 farmID 和 minerID 获取 miner
	miner, err := s.minerRDB.GetByID(ctx, farmID, minerID)
	if err != nil {
		log.Println(farmID, minerID, "minerRDB.GetByID")
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	// 验证密码
	if req.Params.Passwd != miner.Pass {
		log.Println(req.Params.Passwd, miner.Pass, "req.Params.Passwd")
		rsp.Error(ctx, http.StatusInternalServerError, "invalid pass", "")
		return
	}

	// 对 req 的数据进行存储
	s.setMinerStats(ctx, rigID, &req)

	// 通过 minerID 获取其 fsID

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
			Result: dto.ServerRsp_Result{
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
			Result: dto.ServerRsp_Result{
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
	var req dto.HiveOsResReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	s.formatOutput(&req)

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

	log.Println("=======================")
	log.Println("task:", task)
	log.Println("=======================")

	_, farmID, minerID, err := s.hiveOsRDB.GetRigFarmAndMinerID(ctx, rigID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), "hiveOsRDB.GetRigFarmAndMinerID")
		return
	}
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

func (s *HiveOsService) setMinerStats(ctx context.Context, rigID string, req *dto.HiveOsReq) error {
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
	return s.hiveOsRDB.GetMinerStats(ctx, rigID)
}

func (s *HiveOsService) GetMinerInfo(ctx context.Context, rigID string) (*info.MinerInfo, error) {
	return s.hiveOsRDB.GetMinerInfo(ctx, rigID)
}

// 格式化输出
func (s *HiveOsService) formatOutput(req any) {
	jsonInd, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		fmt.Println("json marshal indent error:", err)
		return
	}
	fmt.Printf("%s\n", jsonInd)
}
