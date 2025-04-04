package service

// import (
// 	"context"
// 	"errors"
// 	"miner/common/dto"
// 	"miner/common/status"
// 	"miner/dao/mysql"
// 	"miner/dao/redis"
// 	"miner/model"
// 	"miner/model/info"
// )

// type AdminService struct {
// 	adminDAO     *mysql.AdminDAO
// 	adminRDB     *redis.AdminRDB
// 	bscApiKeyRDB *redis.BscApiKeyRDB
// 	coinRDB      *redis.CoinRDB
// 	poolRDB      *redis.PoolRDB
// 	softRDB      *redis.SoftRDB
// 	poolsRDB     *redis.PoolsRDB
// 	softAllRDB   *redis.SoftAllRDB
// }

// func NewAdminService() *AdminService {
// 	return &AdminService{
// 		adminDAO:     mysql.NewAdminDAO(),
// 		adminRDB:     redis.NewAdminRDB(),
// 		bscApiKeyRDB: redis.NewBscApiKeyRDB(),
// 		coinRDB:      redis.NewCoinRDB(),
// 		poolRDB:      redis.NewPoolRDB(),
// 		softRDB:      redis.NewSoftRDB(),
// 		poolsRDB:     redis.NewPoolsRDB(),
// 		softAllRDB:   redis.NewSoftAllRDB(),
// 	}
// }

// // GetAllUser 获取所有用户信息
// func (s *AdminService) GetAllUser(ctx context.Context) ([]info.User, error) {
// 	return s.adminRDB.GetAllUsers(ctx)
// }

// // GetUserOperLogs 获取用户操作日志
// func (s *AdminService) GetUserOperLogs(ctx context.Context, query map[string]any) ([]model.Operlog, int64, error) {
// 	return s.adminDAO.GetUserOperlogs(ctx, query)
// }

// // GetUserLoginLogs 获取用户登陆日志
// func (s *AdminService) GetUserLoginLogs(ctx context.Context, query map[string]any) ([]model.Loginlog, int64, error) {
// 	return s.adminDAO.GetUserLoginlogs(ctx, query)
// }

// // GetUserPointsRecords 获取用户积分记录
// func (s *AdminService) GetUserPointsRecords(ctx context.Context, query map[string]any) ([]model.Pointslog, int64, error) {
// 	return s.adminDAO.GetUserPointslogs(ctx, query)
// }

// // GetUserFarms 获取用户的矿场
// func (s *AdminService) GetUserFarms(ctx context.Context) ([]info.Farm, error) {
// 	userID, exists := ctx.Value("user_id").(string)
// 	if !exists {
// 		return nil, errors.New("invalid user_id in context")
// 	}
// 	return s.adminRDB.GetUserFarms(ctx, userID)
// }

// // GetUserMiners 获取用户的矿机
// func (s *AdminService) GetUserMiners(ctx context.Context, farmID string) ([]info.Miner, error) {
// 	return s.adminRDB.GetUserMiners(ctx, farmID)
// }

// // SwitchRegister 用户注册开关
// func (s *AdminService) SetSwitchRegister(ctx context.Context, req *dto.AdminSwitchRegisterReq) error {
// 	return s.adminRDB.SetSwitchRegister(ctx, req.Status)
// }

// // GetSwitchRegister 用户注册开关
// func (s *AdminService) GetSwitchRegister(ctx context.Context) (string, error) {
// 	return s.adminRDB.GetSwitchRegister(ctx)
// }

// // SetGlobalFlightsheet 设置全局飞行表
// func (s *AdminService) SetGlobalFs(ctx context.Context, req *dto.AdminSetGlobalFsReq) error {
// 	// fs := &info.Fs{
// 	// 	Name: req.Name,
// 	// 	Coin: req.Coin,
// 	// 	Pool: req.Pool,
// 	// 	Soft: req.Soft,
// 	// }
// 	// return s.adminRDB.SetGlobalFs(ctx, fs)
// 	return nil
// }

// // GetInviteReward 设置邀请奖励
// func (s *AdminService) GetInviteReward(ctx context.Context) (float32, error) {
// 	return s.adminRDB.GetInviteReward(ctx)
// }

// // SetInviteReward 设置邀请积分奖励
// func (s *AdminService) SetInviteReward(ctx context.Context, req *dto.AdminSetInviteRewardReq) error {
// 	return s.adminRDB.SetInviteReward(ctx, req.Reward)
// }

// // GetRechargeRatio 获取充值获取积分比率
// func (s *AdminService) GetRechargeRatio(ctx context.Context) (float64, error) {
// 	return s.adminRDB.GetRechargeRatio(ctx)
// }

// // SetRechargeReward 设置充值获取积分比率
// func (s *AdminService) SetRechargeRatio(ctx context.Context, req *dto.AdminSetRechargeRewardReq) error {
// 	return s.adminRDB.SetRechargeRatio(ctx, req.Ratio)
// }

// // GetUserStatus 获取用户状态
// func (s *AdminService) GetUserStatus(ctx context.Context, userID string) (status.UserStatus, error) {
// 	return s.adminRDB.GetUserStatus(ctx, userID)
// }

// // SetUserStatus 设置用户状态
// func (s *AdminService) SetUserStatus(ctx context.Context, req *dto.AdminSetUserStatusReq) error {
// 	return s.adminRDB.SetUserStatus(ctx, req.UserID, req.Status)
// }

// // SetMinerPoolCost 设置矿池消耗
// func (s *AdminService) SetMinepoolCost(ctx context.Context, req *dto.AdminSetMinePoolCostReq) error {
// 	return s.adminRDB.SetMinepoolCost(ctx, req.MinepoolID, req.Cost)
// }

// // SetMnemonic 设置助记词
// func (s *AdminService) SetMnemonic(ctx context.Context, req *dto.AdminSetMnemonicReq) error {
// 	return s.adminRDB.SetMnemonic(ctx, req.Mnemonic)
// }

// // GetMnemonic 获取活跃助记词
// func (s *AdminService) GetMnemonic(ctx context.Context) (string, error) {
// 	return s.adminRDB.GetMnemonic(ctx)
// }

// // GetAllMnemonic 获取所有助记词
// func (s *AdminService) GetAllMnemonic(ctx context.Context) ([]string, error) {
// 	return s.adminRDB.GetAllMnemonic(ctx)
// }

// // AddApiKey 添加 apikey
// func (s *AdminService) AddBscApiKey(ctx context.Context, apikey string) error {
// 	return s.bscApiKeyRDB.ZAdd(ctx, apikey)
// }

// // GetApiKey 获取 apikey（最少使用）
// func (s *AdminService) GetBscApiKey(ctx context.Context) (string, error) {
// 	return s.bscApiKeyRDB.ZRangeWithScore(ctx)
// }

// // GetAllApiKey 获取所有 apikey
// func (s *AdminService) GetAllBscApiKey(ctx context.Context) ([]string, error) {
// 	return s.bscApiKeyRDB.ZRange(ctx)
// }

// // DelApiKey 删除 apikey
// func (s *AdminService) DelBscApiKey(ctx context.Context, apikey string) error {
// 	return s.bscApiKeyRDB.ZRem(ctx, apikey)
// }

// // AddCoin 添加代币
// func (s *AdminService) AddCoin(ctx context.Context, coin *info.Coin) error {
// 	return s.coinRDB.Set(ctx, coin)
// }

// // DelCoin 删除代币
// func (s *AdminService) DelCoin(ctx context.Context, name string) error {
// 	return s.coinRDB.Del(ctx, name)
// }

// // GetCoinByName 获取代币
// func (s *AdminService) GetCoin(ctx context.Context, name string) (*info.Coin, error) {
// 	return s.coinRDB.Get(ctx, name)
// }

// // GetAllCoin 获取所有代币
// func (s *AdminService) GetAllCoin(ctx context.Context) ([]info.Coin, error) {
// 	return s.coinRDB.GetAll(ctx)
// }

// // AddPool 添加矿池
// func (s *AdminService) AddPool(ctx context.Context, CoinName string, pool *info.Pool) error {
// 	err := s.poolRDB.Set(ctx, CoinName, pool)
// 	if err != nil {
// 		return err
// 	}
// 	err = s.poolsRDB.Set(ctx, pool)
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }

// // DelPool 删除矿池
// func (s *AdminService) DelPool(ctx context.Context, coinName string, poolName string) error {
// 	err := s.poolRDB.Del(ctx, coinName, poolName)
// 	if err != nil {
// 		return err
// 	}
// 	err = s.poolsRDB.Del(ctx, poolName)
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }

// // GetPool 获取矿池
// func (s *AdminService) GetPool(ctx context.Context, coinName string, name string) (*info.Pool, error) {
// 	return s.poolRDB.Get(ctx, coinName, name)
// }

// // GetAllPool 获取所有矿池
// func (s *AdminService) GetAllPool(ctx context.Context) ([]info.Pool, error) {
// 	return s.poolsRDB.GetAll(ctx)
// }

// // GetAllPoolByCoin 根据coin获取所有矿池
// func (s *AdminService) GetAllPoolByCoin(ctx context.Context, coinName string) ([]info.Pool, error) {
// 	return s.poolRDB.GetAll(ctx, coinName)
// }

// //// AddSoft 应用 custom miner soft
// //func (s *AdminService) AddSoft(ctx context.Context, coin string, name string, soft *info.Soft) error {
// //	return s.softRDB.Set(ctx, coin, name, soft)
// //}

// // AddSoft 添加矿池
// func (s *AdminService) AddSoft(ctx context.Context, coinName string, soft *info.Soft) error {
// 	err := s.softRDB.Set(ctx, coinName, soft.MinerName, soft)
// 	if err != nil {
// 		return err
// 	}
// 	err = s.softAllRDB.Set(ctx, soft)
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }

// // DelSoft 删除矿池
// func (s *AdminService) DelSoft(ctx context.Context, coinName string, softName string) error {
// 	err := s.softRDB.Del(ctx, coinName, softName)
// 	if err != nil {
// 		return err
// 	}
// 	err = s.softAllRDB.Del(ctx, softName)
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }

// // GetSoft 获取矿池
// func (s *AdminService) GetSoft(ctx context.Context, coinName string, name string) (*info.Soft, error) {
// 	return s.softRDB.Get(ctx, coinName, name)
// }

// // GetAllSoft 获取所有矿池
// func (s *AdminService) GetAllSoft(ctx context.Context) ([]info.Soft, error) {
// 	return s.softAllRDB.GetAll(ctx)
// }

// // GetAllSoftByCoin 根据coin获取所有矿池
// func (s *AdminService) GetAllSoftByCoin(ctx context.Context, coinName string) ([]info.Soft, error) {
// 	return s.softRDB.GetAll(ctx, coinName)
// }

// // 设置卡数上限
// func (s *AdminService) SetFreeGpuNum(ctx context.Context, num int) error {
// 	return s.adminRDB.SetFreeGpuNum(ctx, num)
// }

// // 获取卡数上线
// func (s *AdminService) GetFreeGpuNum(ctx context.Context) (int, error) {
// 	return s.adminRDB.GetFreeGpuNum(ctx)
// }
