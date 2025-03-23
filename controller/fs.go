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

type FsController struct {
	fsService *services.FsService
}

func NewFsController() *FsController {
	return &FsController{
		fsService: services.NewFsService(),
	}
}

// 创建飞行表
func (m *FsController) CreateFs(ctx *gin.Context) {
	var req dto.CreateFsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	err := m.fsService.CreateFs(ctx, userID, &req)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "create flightsheet success", nil)
}

// DeleteFs 删除飞行表
func (c *FsController) DeleteFs(ctx *gin.Context) {
	var req dto.DelFsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := c.fsService.DelFs(ctx, userID, req.FsID); err != nil {
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

	if err := c.fsService.UpdateFs(ctx, req.FsID, req.UpdateInfo); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "update flightsheet success", nil)
}

// 获取所有飞行表
func (c *FsController) GetFss(ctx *gin.Context) {
	var params params.PageParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	query := map[string]any{
		"page_num":  params.PageNum,
		"page_size": params.PageSize,
	}

	fss, total, err := c.fsService.GetFss(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInsufficientStorage, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "get user all flightsheet success", fss, total)
}

// 获取指定 fs
func (c *FsController) GetFsByID(ctx *gin.Context) {
	fsID, err := strconv.Atoi(ctx.Param("fs_id"))
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	fs, err := c.fsService.GetFsByFsID(ctx, fsID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "get user all flightsheet success", fs)
}

// 获取用户飞行表
func (m *FsController) GetFsByUserID(ctx *gin.Context) {

}

// 获取 farm 使用的飞行表
func (m *FsController) GetFsByFarmID(ctx *gin.Context) {

}

// 获取 miner 使用的飞行表
func (m *FsController) GetFsByMinerID(ctx *gin.Context) {

}
