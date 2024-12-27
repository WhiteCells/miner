package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
	"miner/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FarmController struct {
	farmService *service.FarmService
}

func NewFarmController() *FarmController {
	return &FarmController{
		farmService: service.NewFarmService(),
	}
}

// CreateFarm 创建矿场
func (c *FarmController) CreateFarm(ctx *gin.Context) {
	var req dto.CreateFarmReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	farm, err := c.farmService.CreateFarm(ctx, &req)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "create farm success", farm)
}

// DeleteFarm 删除矿场
func (c *FarmController) DeleteFarm(ctx *gin.Context) {
	var req dto.DeleteFarmReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.farmService.DeleteFarm(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "delete farm success", nil)
}

// UpdateFarm 更新矿场
func (c *FarmController) UpdateFarm(ctx *gin.Context) {
	var req dto.UpdateFarmReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.farmService.UpdateFarm(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "update farm success", nil)
}

// GetFarm 获取用户所有的矿场
func (c *FarmController) GetFarm(ctx *gin.Context) {
	farms, err := c.farmService.GetFarm(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "get farm success", farms)
}

// ApplyFs 矿场应用飞行表
func (c *FarmController) ApplyFs(ctx *gin.Context) {
	var req dto.ApplyFarmFlightsheetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.farmService.ApplyFs(ctx, &req); err != nil {
		return
	}

	rsp.Success(ctx, http.StatusOK, "get user all farm", nil)
}

// Transfer 转移矿场
func (c *FarmController) Transfer(ctx *gin.Context) {
	var req dto.TransferFarmReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.farmService.Transfer(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "transfer farm success", nil)
}
