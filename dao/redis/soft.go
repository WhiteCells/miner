package redis

import (
	"context"
	"encoding/json"
	"miner/model/info"
	"miner/utils"
)

type SoftRDB struct {
}

func NewSoftRDB() *SoftRDB {
	return &SoftRDB{}
}

// 使用 hash 便于查找所有的 soft
// soft 无需与 coin 进行关联
// soft 目前只做 Custom
// +--------+---------+--------+
// | field  | key     | val    |
// +--------+---------+--------+
// | soft   | <name>  | <info> |
// +--------+---------+--------+

func (c *SoftRDB) Set(ctx context.Context, info *info.Soft) error {
	field := MakeField(SoftField)
	key := MakeKey(info.Name)
	infoByte, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, field, key, string(infoByte))
}

func (c *SoftRDB) Del(ctx context.Context, name string) error {
	field := MakeField(SoftField)
	key := MakeKey(name)
	return utils.RDB.HDel(ctx, field, key)
}

func (c *SoftRDB) GetByName(ctx context.Context, name string) (*info.Soft, error) {
	field := MakeField(SoftField)
	key := MakeKey(name)

	softStr, err := utils.RDB.HGet(ctx, field, key)
	if err != nil {
		return nil, err
	}

	var info info.Soft
	if err := json.Unmarshal([]byte(softStr), &info); err != nil {
		return nil, err
	}

	return &info, nil
}

func (c *SoftRDB) GetAll(ctx context.Context) (*[]info.Soft, error) {
	field := MakeField(SoftField)

	infos, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}

	var softs []info.Soft
	for _, infoStr := range infos {
		var soft info.Soft
		if err := json.Unmarshal([]byte(infoStr), &soft); err != nil {
			return nil, err
		}
		softs = append(softs, soft)
	}

	return &softs, nil
}

func (c *SoftRDB) Exists(ctx context.Context, name string) bool {
	_, err := c.GetByName(ctx, name)
	return err == nil
}
