package services

import (
	"context"
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

func (m *WalletService) CreateWallet(ctx context.Context, userID int, wallet *model.Wallet) error {
	return m.walletDAO.CreateWallet(ctx, userID, wallet)
}

func (m *WalletService) DelWallet(ctx context.Context, userID, walletID int) error {
	return m.walletDAO.DelWallet(ctx, userID, walletID)
}

func (m *WalletService) UpdateWallet(ctx context.Context, userID, walletID int, wallet *model.Wallet) error {
	return m.walletDAO.UpdateWallet(ctx, wallet)
}

func (m *WalletService) GetWalletByID(ctx context.Context, userID, walletID int) (*model.Wallet, error) {
	return m.walletDAO.GetWalletByID(ctx, walletID)
}

func (m *WalletService) GetWallets(ctx context.Context, userID int, query map[string]any) (*[]model.Wallet, int64, error) {
	return m.walletDAO.GetWallets(ctx, userID, query)
}
