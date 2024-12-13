package controller

import (
	"miner/common/rsp"
	"miner/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PointsRecordController struct {
	pointsRecordService *service.PointsRecordService
}

func NewPointRecordController() *PointsRecordController {
	return &PointsRecordController{
		pointsRecordService: service.NewPointRecordService(),
	}
}

// GetPointsRecords 获取用户积分记录
func (c *PointsRecordController) GetPointsRecords(ctx *gin.Context) {
	records, total, err := c.pointsRecordService.GetPointsRecords(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.GetPointsRecordsSuccess(ctx, http.StatusOK, "get points records success", records, total)
}
