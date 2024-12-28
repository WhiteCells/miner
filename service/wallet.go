package service

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/dao/redis"
	"miner/model"
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
	_, exists := ctx.Value("user_id").(int)
	if !exists {
		return errors.New("invalid user_id in context")
	}

	return nil
}

func (s *WalletService) UpdateWallet(ctx context.Context, req *dto.UpdateWalletReq) error {
	_, exists := ctx.Value("user_id").(int)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	// wallet, err := s.walletDAO.GetWalletByID(req.WalletID)
	// if err != nil {
	// 	return errors.New("wallet not found")
	// }
	// for key, value := range req.UpdateInfo {
	// 	switch key {
	// 	case "name":
	// 		wallet.Name = value.(string)
	// 	case "address":
	// 		wallet.Address = value.(string)
	// 	case "coin_type":
	// 		wallet.CoinType = value.(string)
	// 	}
	// }
	// if err := s.walletDAO.UpdateWallet(wallet); err != nil {
	// 	return errors.New("update wallet failed")
	// }
	return nil
}

func (s *WalletService) GetWallet(ctx context.Context, query map[string]interface{}) (*[]model.Wallet, error) {
	// userID, exists := ctx.Value("user_id").(int)
	// if !exists {
	// 	return nil, -1, errors.New("invalid user_id in context")
	// }
	// var wallets *[]model.Wallet
	// wallets, total, err := s.walletDAO.GetWallet(userID, query)
	// if err != nil {
	// 	return nil, -1, errors.New("get user all wallet failed")
	// }
	// return wallets, total, nil
	return nil, nil
}

func (s *WalletService) GetUserWalletByID(ctx context.Context, req *dto.GetUserWalletByIDReq) (*model.Wallet, error) {
	// _, exists := ctx.Value("user_id").(int)
	// if !exists {
	// 	return nil, errors.New("invalid user_id in context")
	// }
	// var wallet *model.Wallet
	// wallet, err := s.walletDAO.GetWalletByID(req.WalletID)
	// if err != nil {
	// 	return nil, errors.New("get user wallet failed")
	// }
	// return wallet, nil
	return nil, nil
}
