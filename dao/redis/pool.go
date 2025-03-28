package redis

import (
	"context"
	"encoding/json"
	"miner/model/info"
	"miner/utils"
)

type PoolRDB struct {
}

func NewPoolRDB() *PoolRDB {
	return &PoolRDB{}
}

// coin:pool
// 每一种 coin 都对应多个 pool
// +-------------+---------+--------+
// | field       | key     | val    |
// +-------------+---------+--------+
// | <coin>:pool | <name>  | <info> |
// +-------------+---------+--------+

func (c *PoolRDB) Set(ctx context.Context, coin string, info *info.Pool) error {
	field := MakeField(PoolField, coin)
	key := MakeKey(info.Name)
	infoByte, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, field, key, string(infoByte))
}

func (c *PoolRDB) Del(ctx context.Context, coin string, name string) error {
	field := MakeField(PoolField, coin)
	key := MakeKey(name)
	return utils.RDB.HDel(ctx, field, key)
}

func (c *PoolRDB) Get(ctx context.Context, coin string, name string) (*info.Pool, error) {
	field := MakeField(PoolField, coin)
	key := MakeKey(name)
	infoStr, err := utils.RDB.HGet(ctx, field, key)
	if err != nil {
		return nil, err
	}
	var info_ info.Pool
	if err := json.Unmarshal([]byte(infoStr), &info_); err != nil {
		return nil, err
	}
	return &info_, nil
}

func (c *PoolRDB) GetAll(ctx context.Context, coin string) ([]info.Pool, error) {
	field := MakeField(PoolField, coin)

	infos, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}

	var pools []info.Pool
	for _, infoStr := range infos {
		var info_ info.Pool
		if err := json.Unmarshal([]byte(infoStr), &info_); err != nil {
			return nil, err
		}
		pools = append(pools, info_)
	}

	return pools, nil
}

func (c *PoolRDB) Exists(ctx context.Context, coin string, name string) bool {
	_, err := c.Get(ctx, coin, name)
	return err == nil
}
