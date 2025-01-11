package service

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/dao/mysql"
	"miner/dao/redis"
	"miner/model"
	"miner/model/info"
)

type AdminService struct {
	adminDAO *mysql.AdminDAO
	adminRDB *redis.AdminRDB
}

func NewAdminService() *AdminService {
	return &AdminService{
		adminDAO: mysql.NewAdminDAO(),
		adminRDB: redis.NewAdminRDB(),
	}
}

// GetAllUser 获取所有用户信息
func (s *AdminService) GetAllUser(ctx context.Context) (*[]info.User, error) {
	return s.adminRDB.GetAllUsers(ctx)
}

// GetUserOperLogs 获取用户操作日志
func (s *AdminService) GetUserOperLogs(ctx context.Context, query map[string]interface{}) (*[]model.OperLog, int64, error) {
	return s.adminDAO.GetUserOperLogs(query)
}

// GetUserLoginLogs 获取用户登陆日志
func (s *AdminService) GetUserLoginLogs(ctx context.Context, query map[string]interface{}) (*[]model.LoginLog, int64, error) {
	return s.adminDAO.GetUserLoginLogs(query)
}

// GetUserPointsRecords 获取用户积分记录
func (s *AdminService) GetUserPointsRecords(ctx context.Context, query map[string]interface{}) (*[]model.PointsRecord, int64, error) {
	return s.adminDAO.GetUserPointsRecords(query)
}

// GetUserFarms 获取用户的矿场
func (s *AdminService) GetUserFarms(ctx context.Context) (*[]info.Farm, error) {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}
	return s.adminRDB.GetUserFarms(ctx, userID)
}

// GetUserMiners 获取用户的矿机
func (s *AdminService) GetUserMiners(ctx context.Context, farmID string) (*[]info.Miner, error) {
	return s.adminRDB.GetUserMiners(ctx, farmID)
}

// SwitchRegister 用户注册开关
func (s *AdminService) SetSwitchRegister(ctx context.Context, req *dto.AdminSwitchRegisterReq) error {
	return s.adminRDB.SetSwitchRegister(ctx, req.Status)
}

// SetGlobalFlightsheet 设置全局飞行表
func (s *AdminService) SetGlobalFs(ctx context.Context, req *dto.AdminSetGlobalFsReq) error {
	fs := &info.Fs{
		Name:   req.Name,
		CoinID: req.Coin,
		MineID: req.Mine,
		SoftID: req.Soft,
	}
	return s.adminRDB.SetGlobalFs(ctx, fs)
}

// SetInviteReward 设置邀请积分奖励
func (s *AdminService) SetInviteReward(ctx context.Context, req *dto.AdminSetInviteRewardReq) error {
	return s.adminRDB.SetInviteReward(ctx, req.Reward)
}

// SetRechargeReward 设置充值获取积分比率
func (s *AdminService) SetRechargeRatio(ctx context.Context, req *dto.AdminSetRechargeRewardReq) error {
	return s.adminRDB.SetRechargeRatio(ctx, req.Ratio)
}

// SetUserStatus 设置用户状态
func (s *AdminService) SetUserStatus(ctx context.Context, req *dto.AdminSetUserStatusReq) error {
	return s.adminRDB.SetUserStatus(ctx, req.UserID, req.Status)
}

// SetMinerPoolCost 设置矿池消耗
func (s *AdminService) SetMinepoolCost(ctx context.Context, req *dto.AdminSetMinePoolCostReq) error {
	return s.adminRDB.SetMinepoolCost(ctx, req.MinepoolID, req.Cost)
}

// SetMnemonic 设置助记词
func (s *AdminService) SetMnemonic(ctx context.Context, req *dto.AdminSetMnemonicReq) error {
	return s.adminRDB.SetMnemonic(ctx, req.Mnemonic)
}

// GetMnemonic 获取活跃助记词
func (s *AdminService) GetMnemonic(ctx context.Context) (string, error) {
	return s.adminRDB.GetMnemonic(ctx)
}

// GetAllMnemonic 获取所有助记词
func (s *AdminService) GetAllMnemonic(ctx context.Context) (*[]string, error) {
	return s.adminRDB.GetAllMnemonic(ctx)
}
