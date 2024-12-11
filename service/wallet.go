package service

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/dao/mysql"
	"miner/dao/redis"
	"miner/model"
)

type WalletService struct {
	walletDAO   *mysql.WalletDAO
	walletCache *redis.WalletCache
}

func NewWalletService() *WalletService {
	return &WalletService{
		walletDAO:   mysql.NewWalletDAO(),
		walletCache: redis.NewWalletCache(),
	}
}

func (s *WalletService) CreateWallet(ctx context.Context, req *dto.CreateWalletReq) (*model.Wallet, error) {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}

	// TODO cache 检查钱包是否存在

	wallet := &model.Wallet{
		Name:     req.Name,
		Address:  req.Address,
		CoinType: req.CoinType,
	}
	// 创建钱包
	if err := s.walletDAO.CreateWallet(wallet, userID); err != nil {
		return nil, errors.New("create wallet failed")
	}

	// TODO cache 钱包信息

	return wallet, nil
}

func (s *WalletService) DeleteWallet(ctx context.Context, req *dto.DeleteWalletReq) error {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	if err := s.walletDAO.DeleteWallet(req.WalletID, userID); err != nil {
		return errors.New("delete wallet failed")
	}
	return nil
}

func (s *WalletService) UpdateWallet(ctx context.Context, req *dto.UpdateWalletReq) error {
	_, exists := ctx.Value("user_id").(int)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	wallet, err := s.walletDAO.GetWalletByID(req.WalletID)
	if err != nil {
		return errors.New("wallet not found")
	}
	for key, value := range req.UpdateInfo {
		switch key {
		case "name":
			wallet.Name = value.(string)
		case "address":
			wallet.Address = value.(string)
		case "coin_type":
			wallet.CoinType = value.(string)
		}
	}
	if err := s.walletDAO.UpdateWallet(wallet); err != nil {
		return errors.New("update wallet failed")
	}
	return nil
}

func (s *WalletService) GetUserAllWallet(ctx context.Context) (*[]model.Wallet, error) {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}
	var wallets *[]model.Wallet
	wallets, err := s.walletDAO.GetUserAllWallet(userID)
	if err != nil {
		return nil, errors.New("get user all wallet failed")
	}
	return wallets, nil
}

func (s *WalletService) GetUserWalletByID(ctx context.Context, req *dto.GetUserWalletByIDReq) (*model.Wallet, error) {
	_, exists := ctx.Value("user_id").(int)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}
	var wallet *model.Wallet
	wallet, err := s.walletDAO.GetWalletByID(req.WalletID)
	if err != nil {
		return nil, errors.New("get user wallet failed")
	}
	return wallet, nil
}
