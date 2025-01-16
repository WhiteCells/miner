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

// 使用 hash 便于查找所有的 pool
// +--------+---------+--------+
// | field  | key     | val    |
// +--------+---------+--------+
// | pool   | <name>  | <info> |
// +--------+---------+--------+

func (c *PoolRDB) Set(ctx context.Context, info *info.Pool) error {
	field := MakeField(PoolField)
	key := MakeKey(info.Name)
	infoByte, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, field, key, string(infoByte))
}

func (c *PoolRDB) Del(ctx context.Context, name string) error {
	field := MakeField(PoolField)
	key := MakeKey(name)
	return utils.RDB.HDel(ctx, field, key)
}

func (c *PoolRDB) GetByName(ctx context.Context, name string) (*info.Pool, error) {
	field := MakeField(PoolField)
	key := MakeKey(name)
	infoStr, err := utils.RDB.HGet(ctx, field, key)
	if err != nil {
		return nil, err
	}
	var info info.Pool
	if err := json.Unmarshal([]byte(infoStr), &info); err != nil {
		return nil, err
	}
	return &info, nil
}

func (c *PoolRDB) GetAll(ctx context.Context) (*[]info.Pool, error) {
	field := MakeField(PoolField)

	infos, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}

	var pools []info.Pool
	for _, infoStr := range infos {
		var info info.Pool
		if err := json.Unmarshal([]byte(infoStr), &info); err != nil {
			return nil, err
		}
		pools = append(pools, info)
	}

	return &pools, nil
}

func (c *PoolRDB) Exists(ctx context.Context, name string) bool {
	_, err := c.GetByName(ctx, name)
	return err == nil
}
