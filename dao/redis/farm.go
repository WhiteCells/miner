package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"miner/model"
	"miner/utils"
	"time"
)

type FarmCache struct{}

func NewFarmCache() *FarmCache {
	return &FarmCache{}
}

const (
	farmKeyPrefix     = "farm:"
	farmInfoTimeout   = 30 * time.Minute
	farmMinersTimeout = 10 * time.Minute
)

// 缓存矿场信息
func (c *FarmCache) SetFarmInfo(ctx context.Context, farm *model.Farm) error {
	key := fmt.Sprintf("%s%d:info", farmKeyPrefix, farm.ID)
	farmJSON, err := json.Marshal(farm)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, string(farmJSON), farmInfoTimeout)
}

// 获取缓存的矿场信息
func (c *FarmCache) GetFarmInfo(ctx context.Context, farmID int) (*model.Farm, error) {
	key := fmt.Sprintf("%s%d:info", farmKeyPrefix, farmID)
	farmJSON, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var farm model.Farm
	err = json.Unmarshal([]byte(farmJSON), &farm)
	return &farm, err
}

// 缓存矿场下的矿机列表
func (c *FarmCache) SetFarmMiners(ctx context.Context, farmID int, miners []model.Miner) error {
	key := fmt.Sprintf("%s%d:miners", farmKeyPrefix, farmID)
	minersJSON, err := json.Marshal(miners)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, string(minersJSON), farmMinersTimeout)
}

// 获取缓存的矿场矿机列表
func (c *FarmCache) GetFarmMiners(ctx context.Context, farmID int) ([]model.Miner, error) {
	key := fmt.Sprintf("%s%d:miners", farmKeyPrefix, farmID)
	minersJSON, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var miners []model.Miner
	err = json.Unmarshal([]byte(minersJSON), &miners)
	return miners, err
}

func (c *FarmCache) DeleteFarmCache(ctx context.Context, farmID int) error {
	keys := []string{
		fmt.Sprintf("%s%d:info", farmKeyPrefix, farmID),
		fmt.Sprintf("%s%d:miners", farmKeyPrefix, farmID),
	}

	for _, key := range keys {
		if err := utils.RDB.Del(ctx, key); err != nil {
			return err
		}
	}
	return nil
}
