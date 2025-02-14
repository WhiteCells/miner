package controller

import (
	"miner/common/rsp"
	"miner/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PointsRecordController struct {
	pointsRecordService *service.PointsRecordService
}

func NewPointsRecordController() *PointsRecordController {
	return &PointsRecordController{
		pointsRecordService: service.NewPointRecordService(),
	}
}

// GetPointsRecords 获取用户积分记录
func (c *PointsRecordController) GetPointsRecords(ctx *gin.Context) {
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

	query := map[string]interface{}{
		"user_id":   userID,
		"page_num":  pageNum,
		"page_size": pageSize,
	}

	records, total, err := c.pointsRecordService.GetPointsRecords(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "get points records success", records, total)
}
