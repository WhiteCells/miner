package redis

import (
	"context"
	"encoding/json"
	"errors"
	"miner/common/perm"
	"miner/model/info"
	"miner/utils"
)

type MinerRDB struct{}

func NewMinerRDB() *MinerRDB {
	return &MinerRDB{}
}

// 添加矿机
// 更新矿机
// +-----------------+------------+-------+
// | field           |    key     |  val  |
// +-----------------+------------+-------+
// | miner:<farm_id> | <miner_id> |  info |
// +-----------------+------------+-------+
func (c *MinerRDB) Set(ctx context.Context, farmID string, miner *info.Miner) error {
	field := MakeField(MinerField, farmID)
	minerJSON, err := json.Marshal(miner)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, field, miner.ID, string(minerJSON))
}

// 删除矿机
func (c *MinerRDB) Del(ctx context.Context, farmID string, minerID string) error {
	field := MakeField(MinerField, farmID)
	return utils.RDB.HDel(ctx, field, minerID)
}

// 通过 ID 获取矿机
func (c *MinerRDB) GetByID(ctx context.Context, farmID string, minerID string) (*info.Miner, error) {
	field := MakeField(MinerField, farmID)
	minerJSON, err := utils.RDB.HGet(ctx, field, minerID)
	if err != nil {
		return nil, err
	}
	var miner info.Miner
	err = json.Unmarshal([]byte(minerJSON), &miner)
	return &miner, err
}

// 获取矿场下的所有矿机
func (c *MinerRDB) GetAll(ctx context.Context, farmID string) (*[]info.Miner, error) {
	field := MakeField(MinerField, farmID)
	farmIDMinerID, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}
	var miners []info.Miner
	for minerID := range farmIDMinerID {
		miner, err := c.GetByID(ctx, farmID, minerID)
		if err != nil {
			return nil, err
		}
		miners = append(miners, *miner)
	}
	return &miners, nil
}

// 转移矿机
// fromUserID
// fromFarmID
// fromMinerID
// toUserID
// toFarmID
func (c *MinerRDB) Transfer(ctx context.Context, fromUserID, fromFarmID, MinerID, toUserID, toFarmID string) error {
	// 检查 fromUserID 对 fromFarmID 的权限
	// 检查 toUserID 对 toFarmID 的权限
	if !c.validPerm(ctx, fromUserID, fromFarmID, perm.FarmOwner) ||
		!c.validPerm(ctx, toUserID, toFarmID, perm.FarmOwner) {
		return errors.New("permission denied")
	}

	pipe := utils.RDB.Client.TxPipeline()

	// 删除原有关联
	field := MakeField(FarmField, fromFarmID)
	pipe.HDel(ctx, field, MinerID)

	// 建立新的关联
	field = MakeField(FarmField, toFarmID)
	miner, err := c.GetByID(ctx, fromFarmID, MinerID)
	if err != nil {
		return err
	}
	pipe.HSet(ctx, field, MinerID, miner)

	_, err = pipe.Exec(ctx)

	return err
}

// 应用飞行表
// +---------------------+------------+
// | key                 |    val     |
// +---------------------+------------+
// | miner_fs:<miner_id> |  <fs_id>   |
// +---------------------+------------+
func (c *MinerRDB) ApplyFs(ctx context.Context, farmID string, minerID string, fsID string) error {
	pipe := utils.RDB.Client.TxPipeline()
	// 更新 miner
	// 获取 miner
	field := MakeField(MinerField, farmID)
	minerJSON, err := pipe.HGet(ctx, field, minerID).Result()
	if err != nil {
		return err
	}
	var miner info.Miner
	if err = json.Unmarshal([]byte(minerJSON), &miner); err != nil {
		return err
	}
	miner.FS = fsID
	miner.HiveOsWallet.FsID = fsID
	minerByte, err := json.Marshal(miner)
	if err != nil {
		return err
	}
	pipe.HSet(ctx, field, miner.ID, string(minerByte))

	key := MakeKey(MinerFsField, minerID)
	pipe.Set(ctx, key, fsID, 0)

	_, err = pipe.Exec(ctx)

	return err
}

// 获取应用的飞行表
func (c *MinerRDB) GetApplyFs(ctx context.Context, minerID string) (string, error) {
	key := MakeKey(MinerFsField, minerID)
	return utils.RDB.Get(ctx, key)
}

// 添加管理员
// fromUserID 操作的用户 ID
// fromFarmID 操作的用户矿场 ID
// minerID    用户矿机 ID
// toUserID   添加的管理员 ID
// toFarmID   添加的管理员矿场 ID
// func (c *MinerRDB) AddManager(ctx context.Context, fromUserID, fromFarmID, minerID, toUserID, toFarmID string) error {
// 	return nil
// }

// 是否有权限
func (c *MinerRDB) validPerm(ctx context.Context, userID string, farmID string, p perm.FarmPerm) bool {
	farmField := MakeField(FarmField, userID)
	farmJSON, err := utils.RDB.HGet(ctx, farmField, farmID)
	if err != nil {
		return false
	}

	var farm info.Farm
	if err := json.Unmarshal([]byte(farmJSON), &farm); err != nil {
		return false
	}

	return farm.Perm == p
}
