package redis

import (
	"context"
	"encoding/json"
	"miner/common/status"
	"miner/model/info"
	"miner/utils"
	"strconv"
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

// 修改邀请积分奖励数量
func (c *AdminRDB) SetInviteReward(ctx context.Context, reward int) error {
	return utils.RDB.Set(ctx, AdminInviteRewardField, reward)
}

// 获取邀请积分奖励数量
func (c *AdminRDB) GetInviteReward(ctx context.Context) (int, error) {
	rewardStr, err := utils.RDB.Get(ctx, AdminInviteRewardField)
	if err != nil {
		return 0, err
	}
	reward, err := strconv.Atoi(rewardStr)
	if err != nil {
		return 0, err
	}
	return reward, nil
}

// 修改充值积分奖励比例
func (c *AdminRDB) SetRechargeRatio(ctx context.Context, ratio float64) error {
	return utils.RDB.Set(ctx, AdminRechargeRatioField, ratio)
}

// 获取充值积分奖励比例
func (c *AdminRDB) GetRechargeRatio(ctx context.Context) (float64, error) {
	ratioStr, err := utils.RDB.Get(ctx, AdminRechargeRatioField)
	if err != nil {
		return 0, err
	}
	ratio, err := strconv.ParseFloat(ratioStr, 64)
	if err != nil {
		return 0, err
	}
	return ratio, nil
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
