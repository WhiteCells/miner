package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"miner/model"
	"miner/utils"
	"time"
)

const (
	minerKeyPrefix    = "miner:"
	minerInfoTimeout  = 30 * time.Minute
	minerStatsTimeout = 5 * time.Minute
)

type MinerCache struct{}

func NewMinerCache() *MinerCache {
	return &MinerCache{}
}

// SetMinerInfo 缓存矿机信息
func (c *MinerCache) SetMinerInfo(ctx context.Context, miner *model.Miner) error {
	key := fmt.Sprintf("%s%d:info", minerKeyPrefix, miner.ID)
	minerJSON, err := json.Marshal(miner)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, string(minerJSON), minerInfoTimeout)
}

// GetMinerInfo 获取缓存的矿机信息
func (c *MinerCache) GetMinerInfo(ctx context.Context, minerID int) (*model.Miner, error) {
	key := fmt.Sprintf("%s%d:info", minerKeyPrefix, minerID)
	minerJSON, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var miner model.Miner
	err = json.Unmarshal([]byte(minerJSON), &miner)
	return &miner, err
}

// SetMinerStats 缓存矿机状态信息
func (c *MinerCache) SetMinerStats(ctx context.Context, minerID int, stats map[string]interface{}) error {
	key := fmt.Sprintf("%s%d:stats", minerKeyPrefix, minerID)
	statsJSON, err := json.Marshal(stats)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, string(statsJSON), minerStatsTimeout)
}

// GetMinerStats 获取缓存的矿机状态信息
func (c *MinerCache) GetMinerStats(ctx context.Context, minerID int) (map[string]interface{}, error) {
	key := fmt.Sprintf("%s%d:stats", minerKeyPrefix, minerID)
	statsJSON, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var stats map[string]interface{}
	err = json.Unmarshal([]byte(statsJSON), &stats)
	return stats, err
}
