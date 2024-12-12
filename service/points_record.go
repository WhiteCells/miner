package service

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/dao/mysql"
	"miner/model"
)

type PointsRecordService struct {
	pointsRecordDAO *mysql.PointsRecordDAO
}

func NewPointRecordService() *PointsRecordService {
	return &PointsRecordService{
		pointsRecordDAO: mysql.NewPointRecordDAO(),
	}
}

func (s *PointsRecordService) GetUserPointsRecords(ctx context.Context, req *dto.GetUserPointsRecordsReq) (*[]model.PointsRecord, int, error) {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return nil, -1, errors.New("invalid user_id in context")
	}
	records, num, err := s.pointsRecordDAO.GetUserPointsRecords(userID, req.PageNum, req.PageSize)
	if err != nil {
		return nil, -1, errors.New("get user points records failed")
	}
	return records, int(num), err
}
