package redis

import (
	"context"
	"encoding/json"
	"miner/model/info"
	"miner/utils"
)

type FsRDB struct{}

func NewFsRDB() *FsRDB {
	return &FsRDB{}
}

// 添加飞行表
// 更新飞行表
// +--------------+-----------+-----------------+
// | field        |    key    |  val            |
// ---------------+-----------+-----------------+
// | fs:<user_id> |  <fs_id>  |  info           |
// +--------------+-----------+-----------------+
func (c *FsRDB) Set(ctx context.Context, userID string, fs *info.Fs) error {
	field := MakeField(FsField, userID)
	fsJSON, err := json.Marshal(fs)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, field, fs.ID, string(fsJSON))
}

// 删除飞行表
func (c *FsRDB) Del(ctx context.Context, userID string, fsID string) error {
	field := MakeField(FsField, userID)
	return utils.RDB.HDel(ctx, field, fsID)
}

// 查询飞行表
func (c *FsRDB) GetAll(ctx context.Context, userID string) (*[]info.Fs, error) {
	field := MakeField(FsField, userID)
	idInfo, err := utils.RDB.HGetAll(ctx, field)
	if err != nil {
		return nil, err
	}
	var fss []info.Fs
	for fsID := range idInfo {
		fs, err := c.GetByID(ctx, userID, fsID)
		if err != nil {
			return nil, err
		}
		fss = append(fss, *fs)
	}
	return &fss, nil
}

// 通过 ID 查询
func (c *FsRDB) GetByID(ctx context.Context, userID string, fsID string) (*info.Fs, error) {
	field := MakeField(FsField, userID)
	fsJSON, err := utils.RDB.HGet(ctx, field, fsID)
	if err != nil {
		return nil, err
	}
	var fs info.Fs
	err = json.Unmarshal([]byte(fsJSON), &fs)
	return &fs, err
}

// 应用钱包
// +--------------------+--------------+
// | key                |   val        |
// +--------------------+--------------+
// | fs_wallet:<fs_id>  | <wallet_id>  |
// +--------------------+--------------+
func (c *FsRDB) ApplyWallet(ctx context.Context, userID, fsID, walletID string) error {
	key := MakeKey(FsWalletField, fsID)
	return utils.RDB.Set(ctx, key, walletID)
}

// 应用矿池
// +----------------------+----------------+
// | key                  |   val          |
// +----------------------+----------------+
// | fs_minepool:<fs_id>  | <minepool_id>  |
// +----------------------+----------------+
func (c *FsRDB) ApplyMinepool(ctx context.Context, userID, fsID, minepoolID string) error {
	key := MakeKey(FsMinepoolField, fsID)
	return utils.RDB.Set(ctx, key, minepoolID)
}
