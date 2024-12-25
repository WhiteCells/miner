package redis

import (
	"context"
	"miner/model/info"
)

type WalletRDB struct{}

func NewWalletRDB() *WalletRDB {
	return &WalletRDB{}
}

// 添加钱包
// 更新钱包
// +---------+-----------+-------+
// | field   |    key    |  val  |
// ----------+-----------+-------+
// | wallet  | <user_id> |  info |
// +---------+-----------+-------+
func (c *WalletRDB) Set(ctx context.Context, userID string, wallet info.Wallet) error {
	return nil
}

// 删除钱包
func (c *WalletRDB) Del(ctx context.Context, userID string, walletID string) error {
	return nil
}

// 通过 ID 获取钱包
func (c *WalletRDB) GetByID(ctx context.Context, userID string, walletID string) (*info.Wallet, error) {
	return nil, nil
}

// 获取用户的所有钱包
func (c *WalletRDB) GetAll(ctx context.Context, userID string) (*[]info.Wallet, error) {
	return nil, nil
}
