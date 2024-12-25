package redis

import (
	"context"
	"miner/model/info"
)

type FsRDB struct{}

func NewFsRDB() *FsRDB {
	return &FsRDB{}
}

// 添加飞行表
// 更新飞行表
// +--------+-----------+-------+
// | field  |    key    |  val  |
// ---------+-----------+-------+
// | fs     | <user_id> |  info |
// +--------+-----------+-------+
//
// +--------------+-----------+-----------------+
// | field        |    key    |  val            |
// ---------------+-----------+-----------------+
// | fs_<user_id> | mine_pool |  <mine_pool_id> |
// +--------------+-----------+-----------------+
// | fs_<user_id> | wallet    |  <wallet_id>    |
// +------ -------+-----------+-----------------+
func (c *FsRDB) Set(ctx context.Context, userID string, fs *info.Fs) error {
	return nil
}

// 删除飞行表
func (c *FsRDB) Del(ctx context.Context, userID string, fsID string) error {
	return nil
}

// 查询飞行表
func (c *FsRDB) GetAll(ctx context.Context, userID string) (*[]info.Fs, error) {
	return nil, nil
}

// 通过 ID 查询
func (c *FsRDB) GetByID(ctx context.Context, userID string, fsID string) (*info.Fs, error) {
	return nil, nil
}
