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
// +--------------+-----------+--------+
// | field        |    key    |  val   |
// ---------------+-----------+--------+
// | fs:<user_id> |  <fs_id>  |  info  |
// +--------------+-----------+--------+
func (c *FsRDB) Set(ctx context.Context, userID string, fs *info.Fs) error {
	pipe := utils.RDB.Client.TxPipeline()

	// 设置飞行表
	field := MakeField(FsField, userID)
	fsByte, err := json.Marshal(fs)
	if err != nil {
		return err
	}
	pipe.HSet(ctx, field, fs.ID, string(fsByte))
	// Wallet
	key := MakeField(FsWalletField, fs.ID)
	pipe.Set(ctx, key, fs.WalletID, 0)
	// Pool
	key = MakeField(FsPoolField, fs.ID)
	pipe.Set(ctx, key, fs.Pool, 0)
	// Soft
	key = MakeField(FsSoftField, fs.ID)
	pipe.Set(ctx, key, fs.Soft, 0)

	_, err = pipe.Exec(ctx)

	return err
}

// 删除飞行表
func (c *FsRDB) Del(ctx context.Context, userID string, fsID string) error {
	field := MakeField(FsField, userID)
	return utils.RDB.HDel(ctx, field, fsID)
}

// 查询飞行表
func (c *FsRDB) GetAll(ctx context.Context, userID string) ([]info.Fs, error) {
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
	return fss, nil
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
// | fs:wallet:<fs_id>  | <wallet_id>  |
// +--------------------+--------------+
// func (c *FsRDB) ApplyWallet(ctx context.Context, userID, fsID, walletID string) error {
// 	key := MakeKey(FsWalletField, fsID)
// 	return utils.RDB.Set(ctx, key, walletID)
// }

// 应用矿池
// +------------------+----------------+
// | key              |   val          |
// +------------------+----------------+
// | fs:pool:<fs_id>  | <pool>         |
// +------------------+----------------+
func (c *FsRDB) ApplyPool(ctx context.Context, userID, fsID, pool string) error {
	key := MakeKey(FsPoolField, fsID)
	return utils.RDB.Set(ctx, key, pool)
}

// 应用软件
// +-----------------+---------+
// | key             |   val   |
// +-----------------+---------+
// | fs:soft:<fs_id> |  <soft> |
// +-----------------+---------+
func (c *FsRDB) ApplySoft(ctx context.Context, userID, fsID, soft string) error {
	key := MakeKey(FsSoftField, fsID)
	return utils.RDB.Set(ctx, key, soft)
}
