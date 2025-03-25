package redis

import (
	"context"
	"encoding/json"
	"miner/model/info"
	"miner/utils"
)

type CoinRDB struct {
}

func NewCoinRDB() *CoinRDB {
	return &CoinRDB{}
}

// 使用 hash 便于查找所有的 coin
// +--------+---------+--------+
// | field  | key     | val    |
// +--------+---------+--------+
// | coin   | <name>  | <info> |
// +--------+---------+--------+

func (c *CoinRDB) Set(ctx context.Context, info *info.Coin) error {
	field := MakeField(CoinField)
	key := MakeKey(info.Name)
	infoByte, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, field, key, string(infoByte))
}

func (c *CoinRDB) Del(ctx context.Context, name string) error {
	field := MakeField(CoinField)
	key := MakeKey(name)
	return utils.RDB.HDel(ctx, field, key)
}

func (c *CoinRDB) Get(ctx context.Context, name string) (*info.Coin, error) {
	field := MakeField(CoinField)
	key := MakeKey(name)
	infoStr, err := utils.RDB.HGet(ctx, field, key)
	if err != nil {
		return nil, err
	}
	var info_ info.Coin
	if err := json.Unmarshal([]byte(infoStr), &info_); err != nil {
		return nil, err
	}
	return &info_, nil
}

func (c *CoinRDB) GetAll(ctx context.Context) ([]info.Coin, error) {
	field := MakeField(CoinField)

	infos, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}

	var coins []info.Coin
	for _, infoStr := range infos {
		var info_ info.Coin
		if err := json.Unmarshal([]byte(infoStr), &info_); err != nil {
			return nil, err
		}
		coins = append(coins, info_)
	}

	return coins, nil
}

func (c *CoinRDB) Exists(ctx context.Context, name string) bool {
	_, err := c.Get(ctx, name)
	return err == nil
}
