package controller

import "miner/service"

type WalletController struct {
	walletService *service.WalletService
}

func NewWalletController() *WalletController {
	return &WalletController{
		walletService: service.NewWalletService(),
	}
}
