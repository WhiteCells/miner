package services

import (
	"context"
	"miner/dao/mysql"
	"miner/model"
)

type OperlogService struct {
	operlogDAO *mysql.OperLogDAO
}

func NewOperlogService() *OperlogService {
	return &OperlogService{
		operlogDAO: mysql.NewOperLogDAO(),
	}
}

func (m *OperlogService) GetOperlogByID(ctx context.Context, userID int, query map[string]any) (*[]model.Operlog, int64, error) {
	return m.operlogDAO.GetOperlogByID(ctx, userID, query)
}

func (m *OperlogService) GetOperlogs(ctx context.Context, query map[string]any) (*[]model.Operlog, int64, error) {
	return m.operlogDAO.GetOperlogs(ctx, query)
}
