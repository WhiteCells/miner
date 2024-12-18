package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
	"miner/service"
	"net/http"
	"strconv"

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

// GetFarm 获取用户矿场
func (c *FarmController) GetFarm(ctx *gin.Context) {
	pageNum, err := strconv.Atoi(ctx.Query("page_num"))
	if err != nil || pageNum <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid page_numt", nil)
		return
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil || pageSize <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid page_size", nil)
		return
	}
	query := map[string]interface{}{
		"page_num":  pageNum,
		"page_size": pageSize,
	}
	farms, total, err := c.farmService.GetFarm(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "get farm success", farms, total)
}

// ApplyFlightsheet 矿场应用飞行表
func (c *FarmController) ApplyFlightsheet(ctx *gin.Context) {
	var req dto.ApplyFarmFlightsheetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.farmService.ApplyFlightsheet(ctx, &req); err != nil {
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
