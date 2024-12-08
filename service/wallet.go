package service

import "miner/dao/mysql"

type WalletService struct {
	walletDAO *mysql.WalletDAO
}

func NewWalletService() *WalletService {
	return &WalletService{
		walletDAO: mysql.NewWalletDAO(),
	}
}
