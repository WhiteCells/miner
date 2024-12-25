package redis

import (
	"context"
	"miner/model/info"
)

type MinerRDB struct{}

func NewMinerCache() *MinerRDB {
	return &MinerRDB{}
}

// 添加矿机
// 更新矿机
// +----------------------+------------+-------+
// | field                |    key     |  val  |
// -----------------------+------------+-------+
// | miner                | <user_id>  |  info |
// +----------------------+------------+-------+
// | farm_miner_<farm_id> | <miner_id> |  perm |
// +----------------------+------------+-------+
func (c *MinerRDB) Set(ctx context.Context, userID string, farmID string, miner *info.Miner) error {
	return nil
}

// 删除矿机
func (c *MinerRDB) Del(ctx context.Context, userID string, farmID string, minerID string) error {
	return nil
}

// 通过 ID 获取矿机
func (c *MinerRDB) GetByID(ctx context.Context, userID string, minerID string) (*info.Miner, error) {
	return nil, nil
}

// 获取矿场下的所有矿机
func (c *MinerRDB) GetAll(ctx context.Context, userID string, farmID string) (*[]info.Miner, error) {
	return nil, nil
}

// 转移矿机
func (c *MinerRDB) Transfer(ctx context.Context, fromUserID, fromFarmID, toUserID, toFarmID string) error {
	return nil
}

// 添加管理员
// fromUserID 操作的用户 ID
// fromFarmID 操作的用户矿场 ID
// minerID    用户矿机 ID
// toUserID   添加的管理员 ID
// toFarmID   添加的管理员矿场 ID
func (c *MinerRDB) AddManager(ctx context.Context, fromUserID, fromFarmID, minerID, toUserID, toFarmID string) error {
	return nil
}
