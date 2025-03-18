package redis

import (
	"context"
	"encoding/json"
	"miner/model/info"
	"miner/utils"
)

type SoftAllRDB struct {
}

func NewSoftAllRDB() *SoftAllRDB {
	return &SoftAllRDB{}
}

// 使用 hash 便于查找所有的 Soft
// +-----------+---------+--------+
// | field     | key     | val    |
// +-----------+---------+--------+
// | soft_all  | <name>  | <info> |
// +-----------+---------+--------+

func (c *SoftAllRDB) Set(ctx context.Context, info_ *info.Soft) error {
	field := MakeField(SoftAllField)
	key := MakeKey(info_.MinerName)
	infoByte, err := json.Marshal(info_)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, field, key, string(infoByte))
}

func (c *SoftAllRDB) Del(ctx context.Context, name string) error {
	field := MakeField(SoftAllField)
	key := MakeKey(name)
	return utils.RDB.HDel(ctx, field, key)
}

func (c *SoftAllRDB) Get(ctx context.Context, name string) (*info.Soft, error) {
	field := MakeField(SoftAllField)
	key := MakeKey(name)
	infoStr, err := utils.RDB.HGet(ctx, field, key)
	if err != nil {
		return nil, err
	}
	var info_ info.Soft
	if err := json.Unmarshal([]byte(infoStr), &info_); err != nil {
		return nil, err
	}
	return &info_, nil
}

func (c *SoftAllRDB) GetAll(ctx context.Context) (*[]info.Soft, error) {
	field := MakeField(SoftAllField)

	infos, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}

	var Softs []info.Soft
	for _, infoStr := range infos {
		var info_ info.Soft
		if err := json.Unmarshal([]byte(infoStr), &info_); err != nil {
			return nil, err
		}
		Softs = append(Softs, info_)
	}

	return &Softs, nil
}

func (c *SoftAllRDB) Exists(ctx context.Context, name string) bool {
	_, err := c.Get(ctx, name)
	return err == nil
}
