package services

import (
	"context"
	"miner/dao/mysql"
	"miner/model"
)

type CoinService struct {
	coinDAO *mysql.CoinDAO
}

func NewCoinService() *CoinService {
	return &CoinService{
		coinDAO: mysql.NewCoinDAO(),
	}
}

func (m *CoinService) CreateCoin(ctx context.Context, userID int, coin *model.Coin) error {
	return m.coinDAO.CreateCoin(ctx, coin)
}
