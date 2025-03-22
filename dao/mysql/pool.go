package mysql

import (
	"context"
	"miner/model"
)

type PoolDAO struct {
}

func NewPoolDAO() *PoolDAO {
	return &PoolDAO{}
}

func (PoolDAO) CreatePool(ctx context.Context, userID int, pool *model.Pool) error {
	return nil
}
