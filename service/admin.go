package service

import (
	"context"
	"miner/common/dto"
	"miner/common/perm"
	"miner/dao/mysql"
	"miner/model"
)

type AdminService struct {
	adminDAO *mysql.AdminDAO
}

func NewAdminService() *AdminService {
	return &AdminService{
		adminDAO: mysql.NewAdminDAO(),
	}
}

// GetAllUser 获取所有用户信息
func (s *AdminService) GetAllUser(ctx context.Context, query map[string]interface{}) (*[]model.User, int64, error) {
	return s.adminDAO.GetAllUsers(query)
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
func (s *AdminService) GetUserFarms(ctx context.Context, query map[string]interface{}) (*[]model.Farm, int64, error) {
	return s.adminDAO.GetUserFarms(query)
}

// GetUserMiners 获取用户的矿机
func (s *AdminService) GetUserMiners(ctx context.Context, query map[string]interface{}) (*[]model.Miner, int64, error) {
	return s.adminDAO.GetUserMiners(query)
}

// SwitchRegister 用户注册开关
func (s *AdminService) SwitchRegister(ctx context.Context, req *dto.AdminSwitchRegisterReq) error {
	return s.adminDAO.SwitchRegister(req.Status)
}

// SetGlobalFlightsheet 设置全局飞行表
func (s *AdminService) SetGlobalFlightsheet(ctx context.Context, req *dto.AdminSetGlobalFlightsheetReq) error {
	fs := &model.Flightsheet{
		Name:     req.Name,
		CoinType: req.CoinType,
		MinePool: req.MinerPool,
		MineSoft: req.MineSoft,
		Perm:     perm.Perm(perm.AdminFS),
	}
	return s.adminDAO.SetGlobalFlightsheet(fs)
}

// SetInviteReward 设置邀请积分奖励
func (s *AdminService) SetInviteReward(ctx context.Context, req *dto.AdminSetInviteRewardReq) error {
	return s.adminDAO.SetInviteReward(req.Reward)
}

// SetRechargeReward 设置充值积分奖励
func (s *AdminService) SetRechargeReward(ctx context.Context, req *dto.AdminSetRechargeRewardReq) error {
	return s.adminDAO.SetRechargeReward(req.Reward)
}

// SetUserStatus 设置用户状态
func (s *AdminService) SetUserStatus(ctx context.Context, req *dto.AdminSetUserStatusReq) error {
	return s.adminDAO.SetUserStatus(req.UserID, req.Status)
}

// SetMinerPoolCost 设置矿池消耗
func (s *AdminService) SetMinePoolCost(ctx context.Context, req *dto.AdminSetMinerPoolCostReq) error {
	return s.adminDAO.SetMinePoolCost(req.MinePoolID, req.Cost)
}
