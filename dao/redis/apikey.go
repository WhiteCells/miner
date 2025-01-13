package redis

import (
	"context"
	"errors"
	"miner/utils"
)

/*
  设计目标
  管理员可以添加多个 apikey 用于检查用户代币余额
  在多个 apikey 时，实现简单的负载均衡
  使用 rdb 的分数
*/

// Bsc
type BscApiKeyRDB struct {
}

func NewBscApiKeyRDB() *BscApiKeyRDB {
	return &BscApiKeyRDB{}
}

// ZAdd 添加 apikey（如果添加重复的，会导致会更新分数，所以需要将添加和更新拆分）
func (c *BscApiKeyRDB) ZAdd(ctx context.Context, apikey string) error {
	// 检查 apikey 是否存在
	_, err := utils.RDB.ZScore(ctx, ApiKeyBscField, apikey)
	if err == nil {
		return errors.New("api key exists")
	}
	return utils.RDB.ZAdd(ctx, ApiKeyBscField, apikey)
}

// ZIncrBy 更新 apikey 的分数
func (c *BscApiKeyRDB) ZIncrBy(ctx context.Context, apikey string, increment float64) error {
	// 增加时也需要检查 apikey 是否存在（因为如果不存在会创建）
	_, err := utils.RDB.ZScore(ctx, ApiKeyBscField, apikey)
	if err != nil {
		return errors.New("api key not exists")
	}
	return utils.RDB.ZIncrBy(ctx, ApiKeyBscField, apikey, increment)
}

// ZRem 删除 apikey
func (c *BscApiKeyRDB) ZRem(ctx context.Context, apikey string) error {
	return utils.RDB.ZRem(ctx, ApiKeyBscField, apikey)
}

// ZRangeWithScore 获取 apikey（分数最小的）
func (c *BscApiKeyRDB) ZRangeWithScore(ctx context.Context) (string, error) {
	return utils.RDB.ZRangeWithScore(ctx, ApiKeyBscField)
}

// ZScore 获取 apikey 分数
func (c *BscApiKeyRDB) ZScore(ctx context.Context, apikey string) (float64, error) {
	return utils.RDB.ZScore(ctx, ApiKeyBscField, apikey)
}

// xxx
// type xxxApiKeyRDB struct {
// }
