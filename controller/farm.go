package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
	"miner/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FarmController struct {
	farmService *services.FarmService
}

func NewFarmController() *FarmController {
	return &FarmController{
		farmService: services.NewFarmService(),
	}
}

// 创建矿场
func (m *FarmController) CreateFarm(ctx *gin.Context) {
	var req dto.CreateFarmReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := m.farmService.CreateFarm(ctx, userID, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "create farm success", nil)
}

// 删除矿场
func (c *FarmController) DelFarm(ctx *gin.Context) {
	var req dto.DeleteFarmReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := c.farmService.DelFarm(ctx, userID, req.FarmID); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "delete farm success", nil)
}

// 更新矿场
func (c *FarmController) UpdateFarm(ctx *gin.Context) {
	var req dto.UpdateFarmReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := c.farmService.UpdateFarm(ctx, userID, req.FarmID, req.UpdateInfo); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "update farm success", nil)
}

// UpdateFarmHash 更新矿场hash
// func (c *FarmController) UpdateFarmHash(ctx *gin.Context) {
// 	var req dto.UpdateFarmHashReq
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
// 		return
// 	}

// 	if err := c.farmService.UpdateFarmHash(ctx, &req); err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}

// 	rsp.Success(ctx, http.StatusOK, "update farm hash success", nil)
// }

// 获取用户所有的矿场
func (c *FarmController) GetFarms(ctx *gin.Context) {
	pageNum, err := strconv.Atoi(ctx.Query("page_num"))
	if err != nil || pageNum <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil || pageSize <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid params", nil)
		return
	}
	query := map[string]any{
		"page_num":  pageNum,
		"page_size": pageSize,
	}
	userID := ctx.GetInt("user_id")

	farms, total, err := c.farmService.GetFarms(ctx, userID, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "get farm success", farms, total)
}

// 获取指定矿场
func (c *FarmController) GetFarmByID(ctx *gin.Context) {
	farmID := ctx.Param("farm_id")
	farm, err := c.farmService.GetFarmByID(ctx, farmID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "get farm success", farm)
}

// 应用飞行表
func (c *FarmController) ApplyFs(ctx *gin.Context) {
	var req dto.ApplyFarmFsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.farmService.ApplyFs(ctx, &req); err != nil {
		return
	}

	rsp.Success(ctx, http.StatusOK, "get user all farm", nil)
}

// 转移矿场
func (c *FarmController) Transfer(ctx *gin.Context) {
	var req dto.TransferFarmReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := c.farmService.Transfer(ctx, userID, req.ToUserID, req.FarmID); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "transfer farm success", nil)
}
