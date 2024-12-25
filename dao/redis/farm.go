package redis

import (
	"context"
	"encoding/json"
	"errors"
	"miner/common/perm"
	"miner/model/info"
	"miner/utils"
)

type FarmRDB struct{}

func NewFarmCache() *FarmRDB {
	return &FarmRDB{}
}

// 添加矿场
// 更新矿场
// +---------------------+-----------+-------+
// | field               |    key    |  val  |
// ----------------------+-----------+-------+
// | farm                | <user_id> |  info |
// +---------------------+-----------+-------+
// | user_farm_<user_id> | <farm_id> |  perm |
// +---------------------+-----------+-------+
//
// +---------------------+-----------+-------+
// | field               |    key    |  val  |
// +---------------------+-----------+-------+
// | farm:<user_id>      | <farm_id> |  info |
// +---------------------+-----------+-------+
// | user_farm_<user_id> | <farm_id> |  perm |
// +---------------------+-----------+-------+
func (c *FarmRDB) Set(ctx context.Context, userID string, farm *info.Farm, perm perm.Perm) error {
	pipe := utils.RDB.Client.TxPipeline()

	// farm
	key := GenHField(FarmField, farm.ID)
	farmJSON, err := json.Marshal(farm)
	if err != nil {
		return err
	}
	pipe.HSet(ctx, FarmField, key, string(farmJSON))

	// user_farm
	assField := GenHField(UserFarmField, userID)
	pipe.HSet(ctx, assField, farm.ID, perm)

	_, err = pipe.Exec(ctx)

	return err
}

// 删除矿机
func (c *FarmRDB) Del(ctx context.Context, userID string, farmID string) error {
	pipe := utils.RDB.Client.TxPipeline()

	// farm
	key := GenHField(FarmField, farmID)
	pipe.Del(ctx, key)

	// user_farm
	assField := GenHField(UserFarmField, userID)
	pipe.HDel(ctx, assField, farmID)

	_, err := pipe.Exec(ctx)

	return err
}

// 查询
func (c *FarmRDB) GetAll(ctx context.Context, userID string) (*[]info.Farm, error) {
	field := GenHField(UserFarmField, userID)
	farmIDPerm, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}
	var farms []info.Farm
	for farmID, p := range farmIDPerm {
		// 通过 farmID 查找 farm
		farm, err := c.GetByID(ctx, userID, farmID)
		if err != nil {
			return nil, err
		}
		farm.Perm = perm.Perm(p)
		farms = append(farms, *farm)
	}
	return &farms, nil
}

// filter
// 查询指定权限
// func (c *FarmRDB) GetFliter(ctx context.Context, userID string, perm perm.Perm) (*[]info.Farm, error) {

// }

// 通过 ID 查询
func (c *FarmRDB) GetByID(ctx context.Context, userID string, farmID string) (*info.Farm, error) {
	farmJSON, err := utils.RDB.HGet(ctx, FarmField, userID)
	if err != nil {
		return nil, err
	}
	var farm info.Farm
	err = json.Unmarshal([]byte(farmJSON), &farm)
	return &farm, err
}

// 转移所有权
// 删除原有关联
// 更新新关联
func (c *FarmRDB) Transfer(ctx context.Context, fromID string, toID string, farmID string) error {
	// 检查 fromID 对 farmID 的权限
	if !c.validPerm(ctx, fromID, farmID, perm.FarmOwner) {
		return errors.New("permission denied")
	}

	pipe := utils.RDB.Client.TxPipeline()

	assField := GenHField(UserFarmField, fromID)
	pipe.HDel(ctx, assField, farmID)

	assField = GenHField(userField, fromID)
	pipe.HSet(ctx, assField, toID, perm.FarmOwner)

	_, err := pipe.Exec(ctx)

	return err
}

// 添加管理者
func (c *FarmRDB) AddManager(ctx context.Context, userID string, farmID string, mgrID string) error {
	// 检查 userID 对 farmID 的权限
	if !c.validPerm(ctx, userID, farmID, perm.FarmOwner) {
		return errors.New("permission denied")
	}

	assField := GenHField(UserFarmField, mgrID)
	err := utils.RDB.HSet(ctx, assField, farmID, perm.FarmManager)
	return err
}

// 是否有权限
func (c *FarmRDB) validPerm(ctx context.Context, userID string, farmID string, p perm.FarmPerm) bool {
	assKey := GenHField(UserFarmField, userID)
	permStr, err := utils.RDB.HGet(ctx, assKey, farmID)
	if err != nil {
		return false
	}
	return perm.FarmPerm(permStr) == p
}
