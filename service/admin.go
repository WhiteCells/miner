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
	adminDAO     *mysql.AdminDAO
	adminRDB     *redis.AdminRDB
	bscApiKeyRDB *redis.BscApiKeyRDB
	coinRDB      *redis.CoinRDB
	poolRDB      *redis.PoolRDB
	softRDB      *redis.SoftRDB
}

func NewAdminService() *AdminService {
	return &AdminService{
		adminDAO:     mysql.NewAdminDAO(),
		adminRDB:     redis.NewAdminRDB(),
		bscApiKeyRDB: redis.NewBscApiKeyRDB(),
		coinRDB:      redis.NewCoinRDB(),
		poolRDB:      redis.NewPoolRDB(),
		softRDB:      redis.NewSoftRDB(),
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
		Name: req.Name,
		Coin: req.Coin,
		Pool: req.Pool,
		Soft: req.Soft,
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

// AddApiKey 添加 apikey
func (s *AdminService) AddBscApiKey(ctx context.Context, apikey string) error {
	return s.bscApiKeyRDB.ZAdd(ctx, apikey)
}

// GetApiKey 获取 apikey（最少使用）
func (s *AdminService) GetBscApiKey(ctx context.Context) (string, error) {
	return s.bscApiKeyRDB.ZRangeWithScore(ctx)
}

// GetAllApiKey 获取所有 apikey
func (s *AdminService) GetAllBscApiKey(ctx context.Context) (*[]string, error) {
	return s.bscApiKeyRDB.ZRange(ctx)
}

// DelApiKey 删除 apikey
func (s *AdminService) DelBscApiKey(ctx context.Context, apikey string) error {
	return s.bscApiKeyRDB.ZRem(ctx, apikey)
}

// AddCoin 添加代币
func (s *AdminService) AddCoin(ctx context.Context, coin *info.Coin) error {
	return s.coinRDB.Set(ctx, coin)
}

// DelCoin 删除代币
func (s *AdminService) DelCoin(ctx context.Context, name string) error {
	return s.coinRDB.Del(ctx, name)
}

// GetCoinByName 获取代币
func (s *AdminService) GetCoin(ctx context.Context, name string) (*info.Coin, error) {
	return s.coinRDB.Get(ctx, name)
}

// GetAllCoin 获取所有代币
func (s *AdminService) GetAllCoin(ctx context.Context) (*[]info.Coin, error) {
	return s.coinRDB.GetAll(ctx)
}

// AddPool 添加矿池
func (s *AdminService) AddPool(ctx context.Context, CoinName string, pool *info.Pool) error {
	return s.poolRDB.Set(ctx, CoinName, pool)
}

// DelPool 删除矿池
func (s *AdminService) DelPool(ctx context.Context, coinName string, poolName string) error {
	return s.poolRDB.Del(ctx, coinName, poolName)
}

// GetPool 获取矿池
func (s *AdminService) GetPool(ctx context.Context, coinName string, name string) (*info.Pool, error) {
	return s.poolRDB.Get(ctx, coinName, name)
}

// GetAllPool 获取所有矿池
func (s *AdminService) GetAllPool(ctx context.Context, coinName string) (*[]info.Pool, error) {
	return s.poolRDB.GetAll(ctx, coinName)
}

// AddSoft
func (s *AdminService) AddSoft(ctx context.Context, soft *info.Soft) error {
	return s.softRDB.Set(ctx, soft)
}

// DelSoft
func (s *AdminService) DelSoft(ctx context.Context, name string) error {
	return s.softRDB.Del(ctx, name)
}

// GetSoftByName
func (s *AdminService) GetSoftByName(ctx context.Context, name string) (*info.Soft, error) {
	return s.softRDB.GetByName(ctx, name)
}

// GetAllSoft
func (s *AdminService) GetAllSoft(ctx context.Context) (*[]info.Soft, error) {
	return s.softRDB.GetAll(ctx)
}

// 设置卡数上限
func (s *AdminService) SetFreeGpuNum(ctx context.Context, num int) error {
	return s.adminRDB.SetFreeGpuNum(ctx, num)
}

// 获取卡数上线
func (s *AdminService) GetFreeGpuNum(ctx context.Context) (int, error) {
	return s.adminRDB.GetFreeGpuNum(ctx)
}
