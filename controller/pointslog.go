package controller

import (
	"miner/common/rsp"
	"miner/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PointslogController struct {
	pointsRecordService *services.PointslogService
}

func NewPointsRecordController() *PointslogController {
	return &PointslogController{
		pointsRecordService: services.NewPointlogService(),
	}
}

// GetPointsRecords 获取用户积分记录
func (c *PointslogController) GetPointsRecords(ctx *gin.Context) {
	userIDStr, exists := ctx.Value("user_id").(string)
	if !exists {
		rsp.Error(ctx, http.StatusInternalServerError, "invalid user_id in context", nil)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "invalid user_id in context", nil)
		return
	}
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

	query := map[string]any{
		"user_id":   userID,
		"page_num":  pageNum,
		"page_size": pageSize,
	}

	records, total, err := c.pointsRecordService.GetPointslogs(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "get points records success", records, total)
}
