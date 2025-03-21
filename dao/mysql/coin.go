package mysql

import (
	"miner/model"
	"miner/utils"
)

type CoinDAO struct {
}

func NewCoinDAO() *CoinDAO {
	return &CoinDAO{}
}

func (CoinDAO) CreateCoin(coin *model.Coin) error {
	return utils.DB.Create(coin).Error
}

func (CoinDAO) DelCoin(coinID int) error {
	return utils.DB.Delete(&model.Coin{}, coinID).Error
}

func (CoinDAO) UpdateCoin(coinID int, coin *model.Coin) error {
	return utils.DB.Save(coin).Error
}

func (CoinDAO) GetCoin(coinID int) (*model.Coin, error) {
	var coin model.Coin
	err := utils.DB.First(&coin, coinID).Error
	return &coin, err
}

func (CoinDAO) GetAllCoin() (*[]model.Coin, error) {
	var coin []model.Coin
	err := utils.DB.Find(&coin).Error
	return &coin, err
}
