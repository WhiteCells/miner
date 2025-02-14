package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
	"miner/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MinerController struct {
	minerService *service.MinerService
}

func NewMinerController() *MinerController {
	return &MinerController{
		minerService: service.NewMinerService(),
	}
}

// CreateMiner 创建矿机
func (c *MinerController) CreateMiner(ctx *gin.Context) {
	var req dto.CreateMinerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	miner, err := c.minerService.CreateMiner(ctx, &req)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "create miner success", miner)
}

// DeleteMiner 删除矿机
func (c *MinerController) DeleteMiner(ctx *gin.Context) {
	var req dto.DeleteMinerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if err := c.minerService.DeleteMiner(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "delete miner success", nil)
}

// UpdateMiner 更新矿机
func (c *MinerController) UpdateMiner(ctx *gin.Context) {
	var req dto.UpdateMinerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if err := c.minerService.UpdateMiner(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "update miner success", nil)
}

// GetMiner 获取用户矿机
func (c *MinerController) GetFarmAllMiner(ctx *gin.Context) {
	farmID := ctx.Query("farm_id")
	miners, err := c.minerService.GetFarmAllMiner(ctx, farmID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "get user all miner success", miners)
}

// GetMinerByID
func (c *MinerController) GetMinerByID(ctx *gin.Context) {
	farmID := ctx.Query("farm_id")
	minerID := ctx.Query("miner_id")
	miner, err := c.minerService.GetMinerByID(ctx, farmID, minerID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "get miner by id success", miner)
}

// ApplyFlightsheet 矿机应用飞行表
func (c *MinerController) ApplyFs(ctx *gin.Context) {
	var req dto.ApplyMinerFsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if err := c.minerService.ApplyFs(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "apply miner flightsheet success", nil)
}

// Transfer 转移矿机所有权
func (c *MinerController) Transfer(ctx *gin.Context) {
	var req dto.TransferMinerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if err := c.minerService.Transfer(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "transfer miner success", nil)
}

// 获取 rig.conf
func (c *MinerController) GetRigConf(ctx *gin.Context) {
	farmID := ctx.Query("farm_id")
	minerID := ctx.Query("miner_id")
	conf, err := c.minerService.GetRigConf(ctx, farmID, minerID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get rig.conf", conf)
}
