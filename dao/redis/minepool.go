package redis

import (
	"context"
	"encoding/json"
	"miner/model/info"
	"miner/utils"
)

type MinepoolRDB struct {
}

func NewMinpoolRDB() *MinepoolRDB {
	return &MinepoolRDB{}
}

// 添加矿池
// 更新矿池
// +-------+----------+-------+
// | field |  key     |  val  |
// +-------+----------+-------+
// | mp    |  <mp_id> |  info |
// +-------+----------+-------+
func (r *MinepoolRDB) Set(ctx context.Context, mp *info.Minepool) error {
	field := MakeField(MpField)
	mpJSON, err := json.Marshal(mp)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, field, mp.ID, string(mpJSON))
}

func (r *MinepoolRDB) SetCost(ctx context.Context, mpID string, cost float64) error {
	mp, err := r.GetByID(ctx, mpID)
	if err != nil {
		return err
	}
	mp.Cost = cost
	return r.Set(ctx, mp)
}

func (r *MinepoolRDB) GetByID(ctx context.Context, mpID string) (*info.Minepool, error) {
	field := MakeField(MpField)
	minepoolJSON, err := utils.RDB.HGet(ctx, field, mpID)
	if err != nil {
		return nil, err
	}
	var minepool info.Minepool
	if err := json.Unmarshal([]byte(minepoolJSON), &minepool); err != nil {
		return nil, err
	}
	return &minepool, nil
}

// 获取所有矿池
func (r *MinepoolRDB) GetAll(ctx context.Context) (*[]info.Minepool, error) {
	field := MakeField(MpField)
	idInfo, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}
	var minepools []info.Minepool
	for mpID := range idInfo {
		minepool, err := r.GetByID(ctx, mpID)
		if err != nil {
			return nil, err
		}
		minepools = append(minepools, *minepool)
	}
	return &minepools, nil
}
