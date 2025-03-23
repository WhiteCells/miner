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
	var params params.PageParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	query := map[string]any{
		"page_num":  params.PageNum,
		"page_size": params.PageSize,
	}

	farms, total, err := c.farmService.GetFarms(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "get farm success", farms, total)
}

// 获取指定矿场
func (c *FarmController) GetFarmByFarmID(ctx *gin.Context) {
	farmID, err := strconv.Atoi(ctx.Param("farm_id"))
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid params", nil)
		return
	}

	farm, err := c.farmService.GetFarmByFarmID(ctx, farmID)
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
	userID := ctx.GetInt("user_id")

	if err := c.farmService.ApplyFs(ctx, userID, req.FarmID, req.FsID); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
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
