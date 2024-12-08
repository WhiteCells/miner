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
	minerInfoTimeout = 30 * time.Minute
)

type MinerCache struct{}

func NewMinerCache() *MinerCache {
	return &MinerCache{}
}

func (c *MinerCache) SetMinerInfoByID(ctx context.Context, miner *model.Miner) error {
	key := fmt.Sprintf("miner:%d:info", miner.ID)
	minerJSON, err := json.Marshal(miner)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, string(minerJSON), minerInfoTimeout)
}

func (c *MinerCache) GetMinerInfoByID(ctx context.Context, minerID int) (*model.Miner, error) {
	key := fmt.Sprintf("miner:%d:info", minerID)
	minerJSON, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var miner model.Miner
	err = json.Unmarshal([]byte(minerJSON), &miner)
	return &miner, err
}
