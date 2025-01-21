package redis

import (
	"context"
	"encoding/json"
	"miner/model/info"
	"miner/utils"
)

type WalletRDB struct{}

func NewWalletRDB() *WalletRDB {
	return &WalletRDB{}
}

// 添加钱包
// 更新钱包
// +--------------------+-------------+-------+
// | field              |  key        | val   |
// ---------------------+-------------+-------+
// | wallet:<user_id>   | <wallet_id> | info  |
// +--------------------+-------------+-------+
func (c *WalletRDB) Set(ctx context.Context, userID string, wallet *info.Wallet) error {
	field := MakeField(WalletField, userID)
	walletJSON, err := json.Marshal(wallet)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, field, wallet.ID, string(walletJSON))
}

// 删除钱包
func (c *WalletRDB) Del(ctx context.Context, userID string, walletID string) error {
	field := MakeField(WalletField, userID)
	return utils.RDB.HDel(ctx, field, walletID)
}

// 通过 ID 获取钱包
func (c *WalletRDB) GetByID(ctx context.Context, userID string, walletID string) (*info.Wallet, error) {
	field := MakeField(WalletField, userID)
	walletJSON, err := utils.RDB.HGet(ctx, field, walletID)
	if err != nil {
		return nil, err
	}
	var wallet info.Wallet
	err = json.Unmarshal([]byte(walletJSON), &wallet)
	return &wallet, err
}

// 获取用户的所有钱包
func (c *WalletRDB) GetAll(ctx context.Context, userID string) (*[]info.Wallet, error) {
	field := MakeField(WalletField, userID)
	idInfo, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}
	var wallets []info.Wallet
	for walletID := range idInfo {
		wallet, err := c.GetByID(ctx, userID, walletID)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, *wallet)
	}
	return &wallets, nil
}

// 获取用户的所有钱包
func (c *WalletRDB) GetAllByCoin(ctx context.Context, userID string, coin string) (*[]info.Wallet, error) {
	field := MakeField(WalletField, userID)
	idInfo, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}
	var wallets []info.Wallet
	for walletID := range idInfo {
		wallet, err := c.GetByID(ctx, userID, walletID)
		if wallet.Coin == coin {
			if err != nil {
				return nil, err
			}
			wallets = append(wallets, *wallet)
		}
	}
	return &wallets, nil
}
