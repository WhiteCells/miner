package services

import (
	"context"
	"miner/dao/mysql"
	"miner/model"
)

type PointslogService struct {
	pointslogDAO *mysql.PointslogDAO
}

func NewPointslogService() *PointslogService {
	return &PointslogService{
		pointslogDAO: mysql.NewPointRecordDAO(),
	}
}

func (m *PointslogService) GetPointslogByUserID(ctx context.Context, userID int, query map[string]any) (*[]model.Pointslog, int64, error) {
	return m.pointslogDAO.GetPointslogByID(ctx, userID, query)
}

func (m *PointslogService) GetPointslogs(ctx context.Context, query map[string]any) (*[]model.Pointslog, int64, error) {
	return m.pointslogDAO.GetPointslogs(ctx, query)
}
