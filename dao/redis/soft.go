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
// +----------------+--------+
// | key            | val    |
// +----------------+--------+
// | custom:<fs_id> | <info> |
// +----------------+--------+
// info 存储 Custom 对象json序列化
func (c *SoftRDB) Set(ctx context.Context, fsID string, info *info.Soft) error {
	key := MakeField(CustomField, fsID)
	infoByte, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, string(infoByte))
}

func (c *SoftRDB) Get(ctx context.Context, fsID string) (*info.Soft, error) {
	key := MakeField(CustomField, fsID)
	softStr, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	var info info.Soft
	if err := json.Unmarshal([]byte(softStr), &info); err != nil {
		return nil, err
	}
	return &info, nil
}
