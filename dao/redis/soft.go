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
func (c *SoftRDB) Set(ctx context.Context, coin string, name string, info *info.Soft) error {
	field := MakeField(SoftField, coin)
	key := MakeField(name)
	infoByte, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, field, key, string(infoByte))
}

func (c *SoftRDB) Del(ctx context.Context, coin string, name string) error {
	field := MakeField(SoftField, coin)
	key := MakeField(name)
	return utils.RDB.HDel(ctx, field, key)
}

func (c *SoftRDB) Get(ctx context.Context, coin string, name string) (*info.Soft, error) {
	field := MakeField(SoftField, coin)
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

func (c *SoftRDB) GetAll(ctx context.Context, coin string) ([]info.Soft, error) {
	field := MakeField(SoftField, coin)

	infos, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}

	var softList []info.Soft
	for _, infoStr := range infos {
		var info_ info.Soft
		if err := json.Unmarshal([]byte(infoStr), &info_); err != nil {
			return nil, err
		}
		softList = append(softList, info_)
	}

	return softList, nil
}

func (c *SoftRDB) Exists(ctx context.Context, coin string, name string) bool {
	_, err := c.Get(ctx, coin, name)
	return err == nil
}
