package controller

import (
	"miner/service"

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

func (c *PointsRecordController) GetUserPointsRecords(ctx *gin.Context) {

}
