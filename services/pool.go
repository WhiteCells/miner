package services

import (
	"context"
	"miner/dao/mysql"
	"miner/model"
)

type PoolService struct {
	poolDAO *mysql.PoolDAO
}

func NewPoolService() *PoolService {
	return &PoolService{
		poolDAO: mysql.NewPoolDAO(),
	}
}

func (m *PoolService) CreatePool(ctx context.Context, coinID int, pool *model.Pool) error {
	return m.poolDAO.CreatePool(ctx, coinID, pool)
}
