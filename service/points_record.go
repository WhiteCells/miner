package service

import (
	"errors"
	"miner/dao/mysql"
	"miner/model"

	"github.com/gin-gonic/gin"
)

type PointsRecordService struct {
	pointsRecordDAO *mysql.PointsRecordDAO
}

func NewPointRecordService() *PointsRecordService {
	return &PointsRecordService{
		pointsRecordDAO: mysql.NewPointRecordDAO(),
	}
}

// GetPointsRecords 获取用户积分记录
func (s *PointsRecordService) GetPointsRecords(ctx *gin.Context, query map[string]interface{}) (*[]model.PointsRecord, int64, error) {
	records, total, err := s.pointsRecordDAO.GetUserPointsRecords(query)
	if err != nil {
		return nil, -1, errors.New("get user points records failed")
	}
	return records, total, err
}
