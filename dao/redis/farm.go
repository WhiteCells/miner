package redis

import (
	"context"
	"encoding/json"
	"miner/common/perm"
	"miner/model/info"
	"miner/utils"
)

type FarmRDB struct{}

func NewFarmRDB() *FarmRDB {
	return &FarmRDB{}
}

// 添加矿场
// 更新矿场
// +---------------------+-----------+-------+
// | field               |    key    |  val  |
// +---------------------+-----------+-------+
// | farm:<user_id>      | <farm_id> |  info |
// +---------------------+-----------+-------+
func (c *FarmRDB) Set(ctx context.Context, userID string, farm *info.Farm, perm perm.FarmPerm) error {
	field := MakeField(FarmField, userID)
	farmJSON, err := json.Marshal(farm)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, field, farm.ID, string(farmJSON))
}

// 删除矿机
func (c *FarmRDB) Del(ctx context.Context, userID string, farmID string) error {
	field := MakeField(FarmField, userID)
	return utils.RDB.HDel(ctx, field, farmID)
}

// 查询
func (c *FarmRDB) GetAll(ctx context.Context, userID string) (*[]info.Farm, error) {
	field := MakeField(FarmField, userID)
	idInfo, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}
	var farms []info.Farm
	for farmID := range idInfo {
		farm, err := c.GetByID(ctx, userID, farmID)
		if err != nil {
			return nil, err
		}
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
	field := MakeField(FarmField, userID)
	farmJSON, err := utils.RDB.HGet(ctx, field, farmID)
	if err != nil {
		return nil, err
	}
	var farm info.Farm
	err = json.Unmarshal([]byte(farmJSON), &farm)
	return &farm, err
}

// 转移所有权
func (c *FarmRDB) Transfer(ctx context.Context, fromID string, toID string, farmID string) error {
	pipe := utils.RDB.Client.TxPipeline()

	// 删除原有联系
	field := MakeField(FarmField, fromID)
	pipe.HDel(ctx, field, farmID)

	// 建立新的联系
	farm, err := c.GetByID(ctx, fromID, farmID)
	if err != nil {
		return err
	}
	field = MakeField(FarmField, toID)
	pipe.HSet(ctx, field, farmID, farm)

	_, err = pipe.Exec(ctx)

	return err
}

// 添加成员
func (c *FarmRDB) AddMember(ctx context.Context, userID string, farmID string, memID string) error {
	field := MakeField(FarmField, memID)
	return utils.RDB.HSet(ctx, field, farmID, perm.FarmManager)
}

// 删除成员
func (c *FarmRDB) DelMember(ctx context.Context, userID string, farmID string, memID string) error {
	field := MakeField(FarmField, memID)
	return utils.RDB.HDel(ctx, field, farmID)
}
