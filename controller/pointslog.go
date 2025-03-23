package controller

import (
	"miner/common/params"
	"miner/common/rsp"
	"miner/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PointslogController struct {
	pointslogService *services.PointslogService
}

func NewPointsRecordController() *PointslogController {
	return &PointslogController{
		pointslogService: services.NewPointslogService(),
	}
}

// 获取积分日志
func (m *PointslogController) GetPointslogs(ctx *gin.Context) {
	var params params.PageParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	query := map[string]any{
		"page_num":  params.PageNum,
		"page_size": params.PageSize,
	}

	records, total, err := m.pointslogService.GetPointslogs(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "get points records success", records, total)
}

// 获取指定用户积分日志
func (m *PointslogController) GetPointslogByUserID(ctx *gin.Context) {
	var params params.PageParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	query := map[string]any{
		"page_num":  params.PageNum,
		"page_size": params.PageSize,
	}
	userID := ctx.GetInt("user_id")

	records, total, err := m.pointslogService.GetPointslogByUserID(ctx, userID, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "get points records success", records, total)

}
