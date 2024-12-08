package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"miner/model"
	"miner/utils"
	"time"
)

type WalletCache struct{}

func NewWalletCache() *WalletCache {
	return &WalletCache{}
}

const (
	walletInfoTimeout = 30 * time.Minute
)

func (c *WalletCache) SetWalletInfoByID(ctx context.Context, wallet model.Wallet) error {
	key := fmt.Sprintf("wallet:%d:info", wallet.ID)
	walletJSON, err := json.Marshal(wallet)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, walletJSON, walletInfoTimeout)
}

func (c *WalletCache) GetWalletInfoByID(ctx context.Context, walletID int) (*model.Wallet, error) {
	key := fmt.Sprintf("wallet:%d:info", walletID)
	walletJSON, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	var wallet model.Wallet
	if err = json.Unmarshal([]byte(walletJSON), &wallet); err != nil {
		return nil, err
	}
	return &wallet, nil
}
