package redis

import (
	"context"
	"encoding/json"
	"miner/model/info"
	"miner/utils"
)

type PoolsRDB struct {
}

func NewPoolsRDB() *PoolsRDB {
	return &PoolsRDB{}
}

// 使用 hash 便于查找所有的 Pool
// +--------+---------+--------+
// | field  | key     | val    |
// +--------+---------+--------+
// | Pools  | <name>  | <info> |
// +--------+---------+--------+

func (c *PoolsRDB) Set(ctx context.Context, info_ *info.Pool) error {
	field := MakeField(PoolsField)
	key := MakeKey(info_.Name)
	infoByte, err := json.Marshal(info_)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, field, key, string(infoByte))
}

func (c *PoolsRDB) Del(ctx context.Context, name string) error {
	field := MakeField(PoolsField)
	key := MakeKey(name)
	return utils.RDB.HDel(ctx, field, key)
}

func (c *PoolsRDB) Get(ctx context.Context, name string) (*info.Pool, error) {
	field := MakeField(PoolsField)
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

func (c *PoolsRDB) GetAll(ctx context.Context) ([]info.Pool, error) {
	field := MakeField(PoolsField)

	infos, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}

	var Pools []info.Pool
	for _, infoStr := range infos {
		var info_ info.Pool
		if err := json.Unmarshal([]byte(infoStr), &info_); err != nil {
			return nil, err
		}
		Pools = append(Pools, info_)
	}

	return Pools, nil
}

func (c *PoolsRDB) Exists(ctx context.Context, name string) bool {
	_, err := c.Get(ctx, name)
	return err == nil
}
