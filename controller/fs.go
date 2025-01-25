package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
	"miner/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FsController struct {
	fsService *service.FsService
}

func NewFsController() *FsController {
	return &FsController{
		fsService: service.NewFsService(),
	}
}

// CreateFs 创建飞行表
func (c *FsController) CreateFs(ctx *gin.Context) {
	var req dto.CreateFsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	flightsheet, err := c.fsService.CreateFs(ctx, &req)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "create flightsheet success", flightsheet)
}

// DeleteFs 删除飞行表
func (c *FsController) DeleteFs(ctx *gin.Context) {
	var req dto.DeleteFsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.fsService.DeleteFs(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "delete flightsheet success", nil)
}

// UpdateFs 更新飞行表
func (c *FsController) UpdateFs(ctx *gin.Context) {
	var req dto.UpdateFsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.fsService.UpdateFs(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "update flightsheet success", nil)
}

// GetAllFs 获取所有飞行表
func (c *FsController) GetAllFs(ctx *gin.Context) {
	fss, err := c.fsService.GetAllFs(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInsufficientStorage, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "get user all flightsheet success", fss)
}

// GetFsByID 获取指定 fs
func (c *FsController) GetFsByID(ctx *gin.Context) {
	fsID := ctx.Param("fs_id")
	fs, err := c.fsService.GetFsByID(ctx, fsID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "get user all flightsheet success", fs)
}

// ApplyWallet 飞行表应用钱包
// func (c *FsController) ApplyWallet(ctx *gin.Context) {
// 	var req dto.ApplyWalletReq
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
// 		return
// 	}

// 	if err := c.fsService.ApplyWallet(ctx, &req); err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}

// 	rsp.Success(ctx, http.StatusOK, "apply wallet success", nil)
// }
