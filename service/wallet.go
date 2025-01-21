package service

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/dao/redis"
	"miner/model/info"
	"miner/utils"
)

type WalletService struct {
	walletRDB *redis.WalletRDB
}

func NewWalletService() *WalletService {
	return &WalletService{
		walletRDB: redis.NewWalletRDB(),
	}
}

func (s *WalletService) CreateWallet(ctx context.Context, req *dto.CreateWalletReq) (*info.Wallet, error) {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}

	id, err := utils.GenerateUID()
	if err != nil {
		return nil, errors.New("uid create failed")
	}
	wallet := &info.Wallet{
		ID:   id,
		Name: req.Name,
		Addr: req.Addr,
		Coin: req.Coin,
	}

	s.walletRDB.Set(ctx, userID, wallet)

	return wallet, nil
}

func (s *WalletService) DeleteWallet(ctx context.Context, req *dto.DeleteWalletReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	return s.walletRDB.Del(ctx, userID, req.WalletID)
}

func (s *WalletService) UpdateWallet(ctx context.Context, req *dto.UpdateWalletReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	wallet, err := s.walletRDB.GetByID(ctx, userID, req.WalletID)
	if err != nil {
		return errors.New("wallet not found")
	}

	for key, value := range req.UpdateInfo {
		switch key {
		case "name":
			wallet.Name = value.(string)
		case "addr":
			wallet.Addr = value.(string)
		case "coin":
			wallet.Coin = value.(string)
		}
	}

	if err := s.walletRDB.Set(ctx, userID, wallet); err != nil {
		return errors.New("update wallet failed")
	}

	return nil
}

func (s *WalletService) GetAllWallet(ctx context.Context) (*[]info.Wallet, error) {
	userID := ctx.Value("user_id").(string)
	return s.walletRDB.GetAll(ctx, userID)
}

func (s *WalletService) GetUserWalletByID(ctx context.Context, walletID string) (*info.Wallet, error) {
	userID := ctx.Value("user_id").(string)
	return s.walletRDB.GetByID(ctx, userID, walletID)
}

func (s *WalletService) GetAllWalletByCoin(ctx context.Context, coin string) (*[]info.Wallet, error) {
	userID := ctx.Value("user_id").(string)
	if coin == "" {
		return s.walletRDB.GetAll(ctx, userID)
	}
	return s.walletRDB.GetAllByCoin(ctx, userID, coin)
}

// GetAllWalletAllCoin
func (s *WalletService) GetAllWalletAllCoin(ctx context.Context) (*[]string, error) {
	//userID := ctx.Value("user_id").(string)
	// todo 返回[]string coin列表
	return nil, nil
	//if coin == "" {
	//	return s.walletRDB.GetAll(ctx, userID)
	//}
	//return s.walletRDB.GetAllByCoin(ctx, userID, coin)
}
