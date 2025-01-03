package redis

import (
	"context"
	"encoding/json"
	"miner/model/info"
	"miner/utils"
)

type HiveOsRDB struct {
}

func NewHiveOsRDB() *HiveOsRDB {
	return &HiveOsRDB{}
}

// 设置 OS 矿机
// 更新 OS 矿机
// +-------------+------------------------+
// | key         | value                  |
// +-------------+------------------------+
// | os:<rig_id> | <farm_id>:<miner_id>   |
// +-------------+------------------------+
func (c *HiveOsRDB) SetRig(ctx context.Context, rigID string, farmID string, minerID string) error {
	key := MakeField(OsField, rigID)
	val := MakeVal(farmID, minerID)
	return utils.RDB.Set(ctx, key, val)
}

// 获取 OS 矿机
func (c *HiveOsRDB) GetRigMinerID(ctx context.Context, rigID string) (string, error) {
	key := MakeField(OsField, rigID)
	return utils.RDB.Get(ctx, key)
}

// rig_id 是否存在
func (c *HiveOsRDB) ExistsRigID(ctx context.Context, rigID string) bool {
	field := MakeField(OsField, rigID)
	return utils.RDB.Exists(ctx, field)
}

// 矿机状态
// +---------------------+-----------+
// | key                 | val       |
// +---------------------+-----------+
// | os_miner:<rig_id>   | <info>    |
// +---------------------+-----------+
func (c *HiveOsRDB) SetMinerStatus(ctx context.Context, rigID string, status *info.MinerStatus) error {
	key := MakeKey(OsMinerField, rigID)
	minerStatusJSON, err := json.Marshal(status)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, string(minerStatusJSON))
}

func (c *HiveOsRDB) GetMinerStatus(ctx context.Context, rigID string) error {

	return nil
}
