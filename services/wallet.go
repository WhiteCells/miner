package services

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/dao/mysql"
	"miner/model"
)

type WalletService struct {
	walletDAO *mysql.WalletDAO
}

func NewWalletService() *WalletService {
	return &WalletService{
		walletDAO: mysql.NewWalletDAO(),
	}
}

func (m *WalletService) CreateWallet(ctx context.Context, userID, coinID int, req *dto.CreateWalletReq) error {
	wallet := &model.Wallet{
		Name:    req.Name,
		Address: req.Addr,
	}
	return m.walletDAO.CreateWallet(ctx, userID, coinID, wallet)
}

func (m *WalletService) DelWallet(ctx context.Context, userID, walletID int) error {
	return m.walletDAO.DelWallet(ctx, userID, walletID)
}

func (m *WalletService) UpdateWallet(ctx context.Context, userID, walletID int, updateInfo map[string]any) error {
	allow := model.GetWalletAllowChangeField()
	updates := make(map[string]any)
	for key, val := range updateInfo {
		if allow[key] {
			updates[key] = val
		}
	}
	if len(updates) == 0 {
		return errors.New("no field update")
	}
	return m.walletDAO.UpdateWallet(ctx, userID, walletID, updates)
}

func (m *WalletService) GetWalletByCoinID(ctx context.Context, userID, coinID int, query map[string]any) ([]model.Wallet, int64, error) {
	return m.walletDAO.GetWalletByCoinID(ctx, userID, coinID, query)
}

func (m *WalletService) GetWalletByWalletID(ctx context.Context, userID, walletID int) (*model.Wallet, error) {
	return m.walletDAO.GetWalletByID(ctx, walletID)
}

func (m *WalletService) GetWalletsByUserID(ctx context.Context, userID int, query map[string]any) ([]model.Wallet, int64, error) {
	return m.walletDAO.GetWallets(ctx, userID, query)
}
