package redis

import (
	"context"
	"encoding/json"
	"miner/model/info"
	"miner/utils"
	"strconv"
	"strings"
)

type HiveOsRDB struct {
}

func NewHiveOsRDB() *HiveOsRDB {
	return &HiveOsRDB{}
}

// 设置 OS 矿机
// 更新 OS 矿机
// +-------------+----------------------------------+
// | key         | value                            |
// +-------------+----------------------------------+
// | os:<rig_id> | <user_id>:<farm_id>:<miner_id>   |
// +-------------+----------------------------------+
func (c *HiveOsRDB) SetRigMapping(ctx context.Context, userID string, rigID string, farmID string, minerID string) error {
	key := MakeField(OsField, rigID)
	val := MakeVal(userID, farmID, minerID)
	return utils.RDB.Set(ctx, key, val)
}

// 删除映射关系
func (c *HiveOsRDB) DelRigMapping(ctx context.Context, rigID string) error {
	key := MakeField(OsField, rigID)
	return utils.RDB.Del(ctx, key)
}

// 获取 OS 矿机对应的 farmID 及 minerID
func (c *HiveOsRDB) GetRigFarmAndMinerID(ctx context.Context, rigID string) (userID string, farmID string, minerID string, err error) {
	key := MakeField(OsField, rigID)
	userFarmMinerID, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return "", "", "", err
	}
	parts := strings.Split(userFarmMinerID, ":")
	userID, farmID, minerID = parts[0], parts[1], parts[2]
	return userID, farmID, minerID, err
}

// rig_id 是否存在
func (c *HiveOsRDB) ExistsRigID(ctx context.Context, rigID string) bool {
	field := MakeField(OsField, rigID)
	return utils.RDB.Exists(ctx, field)
}

// 矿机统计信息
// +--------------------+----------+
// | key                | val      |
// +--------------------+----------+
// | os:stats:<rig_id>  | <stats>  |
// +--------------------+----------+
func (c *HiveOsRDB) SetMinerStats(ctx context.Context, rigID int, stats *info.MinerStats) error {
	rigIDStr := strconv.Itoa(rigID)
	key := MakeKey(OsStatsField, rigIDStr)
	minerStatsByte, err := json.Marshal(stats)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, string(minerStatsByte))
}

func (c *HiveOsRDB) GetMinerStats(ctx context.Context, rigID string) (*info.MinerStats, error) {
	key := MakeKey(OsStatsField, rigID)
	minerStatsByte, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	var minerStats info.MinerStats
	if err := json.Unmarshal([]byte(minerStatsByte), &minerStats); err != nil {
		return nil, err
	}
	return &minerStats, nil
}

// 矿机信息
// +-------------------+-----------+
// | key               | val       |
// +-------------------+-----------+
// | os:info:<rig_id>  | <info>    |
// +-------------------+-----------+
func (c *HiveOsRDB) SetMinerInfo(ctx context.Context, rigID int, info *info.MinerInfo) error {
	rigIDStr := strconv.Itoa(rigID)
	key := MakeKey(OsInfoField, rigIDStr)
	minerInfoByte, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, string(minerInfoByte))
}

func (c *HiveOsRDB) GetMinerInfo(ctx context.Context, rigID string) (*info.MinerInfo, error) {
	key := MakeKey(OsInfoField, rigID)
	minerInfoByte, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	var minerInfo info.MinerInfo
	if err := json.Unmarshal([]byte(minerInfoByte), &minerInfo); err != nil {
		return nil, err
	}
	return &minerInfo, nil
}

// 矿场 hash 作为索引
// +-------------+------------------------+
// | key         | value                  |
// +-------------+------------------------+
// | os:<hash>   | <farm_id>:<miner_id>   |
// +-------------+------------------------+
// func (c *HiveOsRDB) SetRigFarmHash(ctx context.Context, farmHash string, farmID string, minerID string) error {
// 	key := MakeKey(OsFarmHashField, farmHash)
// 	val := MakeVal(farmID, minerID)
// 	return utils.RDB.Set(ctx, key, val)
// }

// func (c *HiveOsRDB) GetRigFarmMinerByHash(ctx context.Context, farmHash string) (string, error) {
// 	key := MakeKey(OsFarmHashField, farmHash)
// 	return utils.RDB.Get(ctx, key)
// }
