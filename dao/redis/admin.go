package redis

import (
	"context"
	"encoding/json"
	"miner/common/status"
	"miner/model/info"
	"miner/utils"
)

type AdminRDB struct {
	userRDB     *UserRDB
	farmRDB     *FarmRDB
	minerRDB    *MinerRDB
	SystemRDB   *SystemRDB
	minepoolRDB *MinepoolRDB
}

func NewAdminRDB() *AdminRDB {
	return &AdminRDB{
		userRDB:     NewUserRDB(),
		farmRDB:     NewFarmRDB(),
		minerRDB:    NewMinerRDB(),
		SystemRDB:   NewSystemRDB(),
		minepoolRDB: NewMinpoolRDB(),
	}
}

// 获取所有用户信息
func (c *AdminRDB) GetAllUsers(ctx context.Context) (*[]info.User, error) {
	idInfo, err := utils.RDB.HGetAll(ctx, UserField)
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

// +-----------------------+------+
// | key                   | val  |
// +-----------------------+------+
// | admin_reward_invite   | 111  |
// +-----------------------+------+
// | admin_reward_recharge | 111  |
// +-----------------------+------+
// | admin_switch_register | 1    |
// +-----------------------+------+

// 修改注册开关
func (c *AdminRDB) SetSwitchRegister(ctx context.Context, status status.RegisterStatus) error {
	return utils.RDB.Set(ctx, AdminSwitchRegisterField, status)
}

// 获取注册开关
func (c *AdminRDB) GetSwitchRegister(ctx context.Context) (string, error) {
	return utils.RDB.Get(ctx, AdminSwitchRegisterField)
}

// 修改邀请积分奖励
func (c *AdminRDB) SetRewardInvite(ctx context.Context, reward int) error {
	return utils.RDB.Set(ctx, AdminRewardInviteField, reward)
}

// 修改充值积分奖励
func (c *AdminRDB) SetRewardRecharge(ctx context.Context, reward int) error {
	return utils.RDB.Set(ctx, AdminRewardRechargeField, reward)
}

// +-----------+---------+------+
// + field     | key     | val  |
// +-----------+---------+------+
// + admin_gfs | <fs_id> | info |
// +-----------+---------+------+

// 设置全局飞行表
func (c *AdminRDB) SetGlobalFs(ctx context.Context, fs *info.Fs) error {
	fsJSON, err := json.Marshal(fs)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, AdminGfsField, fs.ID, string(fsJSON))
}

// 设置用户状态
func (c *AdminRDB) SetUserStatus(ctx context.Context, userID string, s status.UserStatus) error {
	user, err := c.userRDB.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	user.Status = s
	return c.userRDB.Set(ctx, user)
}

// 设置矿池的消耗
func (c *AdminRDB) SetMinepoolCost(ctx context.Context, mpID string, cost float64) error {
	mp, err := c.minepoolRDB.GetByID(ctx, mpID)
	if err != nil {
		return err
	}
	mp.Cost = cost
	return c.minepoolRDB.Set(ctx, mp)
}
