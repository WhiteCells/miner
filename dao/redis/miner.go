package redis

import (
	"context"
	"encoding/json"
	"errors"
	"miner/model/info"
	"miner/utils"
	"strconv"
)

type MinerRDB struct{}

func NewMinerRDB() *MinerRDB {
	return &MinerRDB{}
}

// key
// rigID:<info>
func (MinerRDB) CreateMinerByRigID(ctx context.Context, rigID int, miner *info.Miner) error {
	rigIDStr := strconv.Itoa(rigID)
	field := MakeField(MinerField, rigIDStr)
	minerJSON, err := json.Marshal(miner)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, field, string(minerJSON))
}

func (MinerRDB) DelMinerByRigID(ctx context.Context, rigID int) error {
	rigIDStr := strconv.Itoa(rigID)
	field := MakeField(MinerField, rigIDStr)
	return utils.RDB.Del(ctx, field)
}

func (MinerRDB) UpdateMinerByRigID(ctx context.Context, rigID int, miner *info.Miner) error {
	rigIDStr := strconv.Itoa(rigID)
	field := MakeField(MinerField, rigIDStr)
	minerJSON, err := json.Marshal(miner)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, field, string(minerJSON))
}

func (MinerRDB) GetMinerByRigID(ctx context.Context, rigID int) (*info.Miner, error) {
	rigIDStr := strconv.Itoa(rigID)
	field := MakeField(MinerField, rigIDStr)
	minerJSON, err := utils.RDB.Get(ctx, field)
	if err != nil {
		return nil, errors.New("redis miner not found")
	}
	var miner info.Miner
	err = json.Unmarshal([]byte(minerJSON), &miner)
	return &miner, err
}

// 添加矿机
// 更新矿机
// +-----------------+------------+-------+
// | field           |    key     |  val  |
// +-----------------+------------+-------+
// | miner:<farm_id> | <miner_id> |  info |
// +-----------------+------------+-------+
// key
// rig_id:<farm_id:miner_id>:<info>
// func (c *MinerRDB) Set(ctx context.Context, farmID int, miner *info.Miner) error {
// 	farmIDStr := strconv.Itoa(farmID)
// 	field := MakeField(MinerField, farmIDStr)
// 	minerJSON, err := json.Marshal(miner)
// 	if err != nil {
// 		return err
// 	}
// 	return utils.RDB.HSet(ctx, field, miner.ID, string(minerJSON))
// }

// 删除矿机
// func (c *MinerRDB) Del(ctx context.Context, farmID string, minerID string) error {
// 	field := MakeField(MinerField, farmID)
// 	return utils.RDB.HDel(ctx, field, minerID)
// }

// // 通过 ID 获取矿机
// func (c *MinerRDB) GetByID(ctx context.Context, farmID, minerID int) (*info.Miner, error) {
// 	farmIDStr := strconv.Itoa(farmID)
// 	minerIDStr := strconv.Itoa(minerID)
// 	field := MakeField(MinerField, farmIDStr)
// 	minerJSON, err := utils.RDB.HGet(ctx, field, minerIDStr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var miner info.Miner
// 	err = json.Unmarshal([]byte(minerJSON), &miner)
// 	return &miner, err
// }

// 获取矿场下的所有矿机
// func (c *MinerRDB) GetAll(ctx context.Context, farmID int) ([]info.Miner, error) {
// 	farmIDStr := strconv.Itoa(farmID)
// 	field := MakeField(MinerField, farmIDStr)
// 	farmIDMinerID, err := utils.RDB.HGetAll(ctx, field)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var miners []info.Miner
// 	for minerID := range farmIDMinerID {
// 		miner, err := c.GetByID(ctx, farmID, minerID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		miners = append(miners, *miner)
// 	}
// 	return &miners, nil
// }

// 转移矿机
// fromUserID
// fromFarmID
// fromMinerID
// toUserID
// toFarmID
// func (c *MinerRDB) Transfer(ctx context.Context, fromUserID, fromFarmID, MinerID, toUserID, toFarmID string) error {
// 	// 检查 fromUserID 对 fromFarmID 的权限
// 	// 检查 toUserID 对 toFarmID 的权限
// 	if !c.validPerm(ctx, fromUserID, fromFarmID, perm.FarmOwner) ||
// 		!c.validPerm(ctx, toUserID, toFarmID, perm.FarmOwner) {
// 		return errors.New("permission denied")
// 	}

// 	pipe := utils.RDB.Client.TxPipeline()

// 	// 删除原有关联
// 	field := MakeField(FarmField, fromFarmID)
// 	pipe.HDel(ctx, field, MinerID)

// 	// 建立新的关联
// 	field = MakeField(FarmField, toFarmID)
// 	miner, err := c.GetByID(ctx, fromFarmID, MinerID)
// 	if err != nil {
// 		return err
// 	}
// 	pipe.HSet(ctx, field, MinerID, miner)

// 	_, err = pipe.Exec(ctx)

// 	return err
// }

// 应用飞行表
// +---------------------+------------+
// | key                 |    val     |
// +---------------------+------------+
// | miner:fs:<miner_id> |  <fs_id>   |
// +---------------------+------------+
// func (c *MinerRDB) ApplyFs(ctx context.Context, farmID string, minerID string, fsID string, softInfo *info.Soft) error {
// 	pipe := utils.RDB.Client.TxPipeline()
// 	// 获取 miner
// 	field := MakeField(MinerField, farmID)
// 	minerJSON, err := utils.RDB.Client.HGet(ctx, field, minerID).Result()
// 	if err != nil {
// 		return err
// 	}
// 	var miner info.Miner
// 	if err = json.Unmarshal([]byte(minerJSON), &miner); err != nil {
// 		return err
// 	}
// 	// 更新 miner
// 	miner.HiveOsWallet.FsID = fsID

// 	miner.HiveOsWallet.CustomMiner = softInfo.MinerName
// 	miner.HiveOsWallet.CustomUserConfig = softInfo.CustomUserConfig
// 	miner.HiveOsWallet.CustomAlgo = softInfo.CustomAlgo
// 	miner.HiveOsWallet.CustomInstallURL = softInfo.CustomInstallUrl
// 	miner.HiveOsWallet.CustomPass = softInfo.CustomPass
// 	miner.HiveOsWallet.CustomTemplate = softInfo.CustomTemplate
// 	miner.HiveOsWallet.CustomUrl = softInfo.CustomUrl
// 	miner.HiveOsWallet.CustomTLS = softInfo.CustomTls

// 	minerByte, err := json.Marshal(miner)
// 	if err != nil {
// 		return err
// 	}
// 	pipe.HSet(ctx, field, miner.ID, string(minerByte))
// 	// 建立关联
// 	key := MakeKey(MinerFsField, minerID)
// 	pipe.Set(ctx, key, fsID, 0)

// 	_, err = pipe.Exec(ctx)

// 	return err
// }

// 获取应用的飞行表
// func (c *MinerRDB) GetApplyFs(ctx context.Context, minerID string) (string, error) {
// 	key := MakeKey(MinerFsField, minerID)
// 	return utils.RDB.Get(ctx, key)
// }

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
// func (c *MinerRDB) validPerm(ctx context.Context, userID string, farmID string, p perm.FarmPerm) bool {
// 	farmField := MakeField(FarmField, userID)
// 	farmJSON, err := utils.RDB.HGet(ctx, farmField, farmID)
// 	if err != nil {
// 		return false
// 	}

// 	var farm info.Farm
// 	if err := json.Unmarshal([]byte(farmJSON), &farm); err != nil {
// 		return false
// 	}

// 	return farm.Perm == p
// }
