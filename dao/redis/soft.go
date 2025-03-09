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

// soft 无需与 coin 进行关联
// soft 目前只做 Custom
// Custom 不做 curd
// 每个 fs 都有自己的 Custom
//
// +---------+---------+--------+
// | fields  | key     | val    |
// +---------+---------+--------+
// | soft    | <name>  | <info> |
// +---------+---------+--------+
// info 存储 Custom 对象json序列化
func (c *SoftRDB) Set(ctx context.Context, name string, info *info.Soft) error {
	field := MakeField(SoftField)
	key := MakeField(name)
	infoByte, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, field, key, string(infoByte))
}

func (c *SoftRDB) Del(ctx context.Context, name string) error {
	field := MakeField(SoftField)
	key := MakeField(name)
	return utils.RDB.HDel(ctx, field, key)
}

func (c *SoftRDB) Get(ctx context.Context, name string) (*info.Soft, error) {
	field := MakeField(SoftField)
	key := MakeField(name)
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
