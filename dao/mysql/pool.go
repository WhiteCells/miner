package mysql

import (
	"context"
	"miner/model"
	"miner/model/relation"
	"miner/utils"

	"gorm.io/gorm"
)

type PoolDAO struct {
}

func NewPoolDAO() *PoolDAO {
	return &PoolDAO{}
}

func (PoolDAO) CreatePool(ctx context.Context, coinID int, pool *model.Pool) error {
	return utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// pool
		if err := tx.Create(pool).Error; err != nil {
			return err
		}
		// coin-pool
		coinPool := &relation.CoinPool{
			CoinID: coinID,
			PoolID: pool.ID,
		}
		if err := tx.Create(coinPool).Error; err != nil {
			return err
		}
		return nil
	})
}

func (PoolDAO) DelPool(ctx context.Context, poolID int) error {
	return utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// coin-pool
		if err := utils.DB.WithContext(ctx).
			Delete(&relation.CoinPool{}, "pool_id=?", poolID).Error; err != nil {
			return err
		}
		// pool
		if err := utils.DB.WithContext(ctx).
			Delete(&model.Pool{}, "id=?", poolID).Error; err != nil {
			return err
		}
		return nil
	})
}

func (PoolDAO) GetPoolByID(ctx context.Context, poolID int) (*model.Pool, error) {
	var pool model.Pool
	err := utils.DB.WithContext(ctx).Find(&pool, poolID).Error
	return &pool, err
}

func (PoolDAO) GetPools(ctx context.Context, query map[string]any) ([]model.Pool, int64, error) {
	var pools []model.Pool
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	db := utils.DB.WithContext(ctx).
		Model(&model.Pool{})

	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	if err := db.
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&pools).
		Error; err != nil {
		return nil, -1, err
	}

	return pools, total, nil
}
