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

func (m *CoinService) DelCoin(ctx context.Context, userID, coinID int) error {
	return m.coinDAO.DelCoin(ctx, coinID)
}

func (m *CoinService) UpdateCoin(ctx context.Context, userID, coinID int, updateInfo map[string]any) error {
	//
	return m.coinDAO.UpdateCoin(ctx, userID, coinID, updateInfo)
}

func (m *CoinService) GetCoinByID(ctx context.Context, userID, coinID int) (*model.Coin, error) {
	//
	return m.coinDAO.GetCoin(ctx, coinID)
}
