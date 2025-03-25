package mysql

import (
	"context"
	"miner/model"
	"miner/utils"
)

type CoinDAO struct {
}

func NewCoinDAO() *CoinDAO {
	return &CoinDAO{}
}

func (CoinDAO) CreateCoin(ctx context.Context, coin *model.Coin) error {
	return utils.DB.WithContext(ctx).Create(coin).Error
}

func (CoinDAO) DelCoin(ctx context.Context, coinID int) error {
	return utils.DB.WithContext(ctx).Delete(&model.Coin{}, coinID).Error
}

func (CoinDAO) UpdateCoin(ctx context.Context, coinID int, coin *model.Coin) error {
	return utils.DB.WithContext(ctx).Save(coin).Error
}

func (CoinDAO) GetCoin(ctx context.Context, coinID int) (*model.Coin, error) {
	var coin model.Coin
	err := utils.DB.WithContext(ctx).First(&coin, coinID).Error
	return &coin, err
}

func (CoinDAO) GetCoins(ctx context.Context, query map[string]any) ([]model.Coin, int64, error) {
	var coin []model.Coin
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	if err := utils.DB.WithContext(ctx).
		Model(&model.Coin{}).Count(&total).Error; err != nil {
		return nil, -1, err
	}

	err := utils.DB.WithContext(ctx).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&coin).
		Error
	return coin, total, err
}
