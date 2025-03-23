package controller

import (
	"miner/common/dto"
	"miner/common/params"
	"miner/common/rsp"
	"miner/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MinerController struct {
	minerService *services.MinerService
}

func NewMinerController() *MinerController {
	return &MinerController{
		minerService: services.NewMinerService(),
	}
}

// 创建矿机
func (m *MinerController) CreateMiner(ctx *gin.Context) {
	var req dto.CreateMinerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := m.minerService.CreateMiner(ctx, userID, req.FarmID, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "create miner success", nil)
}

// 删除矿机
func (m *MinerController) DeleteMiner(ctx *gin.Context) {
	var req dto.DelMinerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := m.minerService.DelMiner(ctx, userID, req.FarmID, req.MinerID); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "delete miner success", nil)
}

// 更新矿机
func (m *MinerController) UpdateMiner(ctx *gin.Context) {
	var req dto.UpdateMinerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := m.minerService.UpdateMiner(ctx, userID, req.FarmID, req.MinerID, req.UpdateInfo); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "update miner success", nil)
}

// 更新矿机 watchdog
func (c *MinerController) UpdateMinerWatchdog(ctx *gin.Context) {
	var req dto.UpdateMinerWatchdogReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := c.minerService.UpdateMinerWatchdog(ctx, userID, req.FarmID, req.MinerID, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "update miner watchdog success", nil)
}

// 更新矿机 options
func (c *MinerController) UpdateMinerOptions(ctx *gin.Context) {
	var req dto.UpdateMinerOptionsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := c.minerService.UpdateMinerOptions(ctx, userID, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "update miner options success", nil)
}

// 更新矿机 autofan
func (c *MinerController) UpdateMinerAutofan(ctx *gin.Context) {
	var req dto.UpdateMinerAutofanReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := c.minerService.UpdateMinerAutofan(ctx, userID, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "update miner autofan success", nil)
}

// 更新矿机 wallet
func (c *MinerController) UpdateMinerWallet(ctx *gin.Context) {
	var req dto.UpdateMinerWalletReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := c.minerService.UpdateMinerWallet(ctx, userID, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "update miner wallet success", nil)
}

// 获取矿场下的矿机
func (c *MinerController) GetFarmAllMiner(ctx *gin.Context) {
	var params params.PageParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	query := map[string]any{
		"page_num":  params.PageNum,
		"page_size": params.PageSize,
	}
	farmID, err := strconv.Atoi(ctx.Query("farm_id"))
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	miners, total, err := c.minerService.GetMinersByFarmID(ctx, farmID, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "get user all miner success", miners, total)
}

// 获取指定矿机
func (c *MinerController) GetMinerByMinerID(ctx *gin.Context) {
	minerID, err := strconv.Atoi(ctx.Query("miner_id"))
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	miner, err := c.minerService.GetMinerByMinerID(ctx, userID, minerID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "get miner by id success", miner)
}

// 矿机应用飞行表
func (c *MinerController) ApplyFs(ctx *gin.Context) {
	var req dto.ApplyMinerFsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := c.minerService.ApplyFs(ctx, userID, req.FarmID, req.MinerID, req.FsID); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "apply miner flightsheet success", nil)
}

// 转移矿机所有权
func (c *MinerController) Transfer(ctx *gin.Context) {
	var req dto.TransferMinerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if err := c.minerService.Transfer(ctx, req.FromFarmID, req.MinerID, req.ToFarmHash); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "transfer miner success", nil)
}

// GetRigConf 获取 rig.conf
// func (c *MinerController) GetRigConf(ctx *gin.Context) {
// 	farmID := ctx.Query("farm_id")
// 	minerID := ctx.Query("miner_id")
// 	conf, err := c.minerService.GetRigConf(ctx, farmID, minerID)
// 	if err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}
// 	rsp.Success(ctx, http.StatusOK, "get rig.conf", conf)
// }

// 设置矿机 watchdog 选项
func (c *MinerController) SetWatchdog(ctx *gin.Context) {
	var req dto.SetWatchdogReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := c.minerService.SetWatchdog(ctx, userID, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "set watchdog success", nil)
}

// 获取矿机 watchdog 选项
func (c *MinerController) GetWatchdog(ctx *gin.Context) {
	farmID, err := strconv.Atoi(ctx.Query("farm_id"))
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	minerID, err := strconv.Atoi(ctx.Query("miner_id"))
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	watchdog, err := c.minerService.GetWatchdog(ctx, userID, farmID, minerID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get watchdog success", watchdog)
}

// 设置矿机 fan 选项
func (c *MinerController) SetAutoFan(ctx *gin.Context) {
	var req dto.SetAutoFanReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if err := c.minerService.SetAutoFan(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get watchdog success", nil)
}

// 获取矿机 fan 选项
func (c *MinerController) GetAutoFan(ctx *gin.Context) {
	farmID, err := strconv.Atoi(ctx.Query("farm_id"))
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	minerID, err := strconv.Atoi(ctx.Query("miner_id"))
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	autofan, err := c.minerService.GetAutoFan(ctx, farmID, minerID)
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get auto fan success", autofan)
}

// 设置矿机 worker 选项
func (c *MinerController) SetOptions(ctx *gin.Context) {
	var req dto.SetOptionsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	userID := ctx.GetInt("user_id")

	if err := c.minerService.SetOptions(ctx, userID, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "set options failed", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "set options success", nil)
}

// 获取矿机 worker 选项
func (c *MinerController) GetOptions(ctx *gin.Context) {
	farmID, err := strconv.Atoi(ctx.Query("farm_id"))
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	minerID, err := strconv.Atoi(ctx.Query("miner_id"))
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	options, err := c.minerService.GetOptions(ctx, userID, farmID, minerID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "get options failed", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get options success", options)
}
