package redis

import (
	"context"
	"miner/common/status"
	"miner/model/info"
	"miner/utils"
)

type AdminRDB struct {
	userRDB   *UserRDB
	farmRDB   *FarmRDB
	minerRDB  *MinerRDB
	SystemRDB *SystemRDB
}

func NewAdminRDB() *AdminRDB {
	return &AdminRDB{
		userRDB:   NewUserRDB(),
		farmRDB:   NewFarmRDB(),
		minerRDB:  NewMinerRDB(),
		SystemRDB: NewSystemRDB(),
	}
}

// 获取所有用户信息
func (c *AdminRDB) GetAllUsers(ctx context.Context) (*[]info.User, error) {
	idInfo, err := utils.RDB.HGetAll(ctx, userField)
	if err != nil {
		return nil, err
	}
	var users []info.User
	for userID := range idInfo {
		user, err := c.userRDB.GetByID(ctx, userID)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return &users, nil
}

// 获取用户操作日志
// func (c *AdminRDB) GetUserOperLogs(ctx context.Context) (*)

// 获取用户登陆日志
// func (c *AdminRDB) GetUserLoginLogs(ctx context.Context) (*)

// 获取用户积分记录
// func (c *AdminRDB) GetUserPointsRecords(ctx context.Context) (*)

// 获取指定用户的所有矿场
func (c *AdminRDB) GetUserFarms(ctx context.Context, userID string) (*[]info.Farm, error) {
	return c.farmRDB.GetAll(ctx, userID)
}

// 获取指定用户的所有矿机
func (c *AdminRDB) GetUserMiners(ctx context.Context, farmID string) (*[]info.Miner, error) {
	return c.minerRDB.GetAll(ctx, farmID)
}

// +--------+-----------------+------+
// + field  | key             | val  |
// +--------+-----------------+------+
// + admin  | reward_invite   | 111  |
// + admin  | reward_recharge | 111  |
// + admin  | switch_register | 1    |
// +--------+-----------------+------+

// 修改注册开关
func (c *AdminRDB) SwitchRegister(ctx context.Context, status status.RegisterStatus) error {
	return nil
}

// 修改邀请积分奖励
func (c *AdminRDB) RewardInvite(ctx context.Context, reward int) error {
	return nil
}

// 修改充值积分奖励
func (c *AdminRDB) RewardRecharge(ctx context.Context, reward int) error {
	return nil
}

// 设置全局飞行表
func (c *AdminRDB) SetGlobalFs(ctx context.Context, fs *info.Fs) error {
	return nil
}
