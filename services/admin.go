package services

import (
	"context"
	"miner/common/status"
	"miner/dao/mysql"
	"miner/dao/redis"
	"miner/model"
)

type AdminService struct {
	adminDAO     *mysql.AdminDAO
	farmDAO      *mysql.FarmDAO
	adminRDB     *redis.AdminRDB
	bscApiKeyRDB *redis.BscApiKeyRDB
}

func NewAdminService() *AdminService {
	return &AdminService{
		adminDAO: mysql.NewAdminDAO(),
		farmDAO:  mysql.NewFarmDAO(),
		adminRDB: redis.NewAdminRDB(),
	}
}

func (m *AdminService) GetUsers(ctx context.Context, query map[string]any) ([]model.User, int64, error) {
	return m.adminDAO.GetUsers(ctx, query)
}

func (m *AdminService) DelUser(ctx context.Context, userID int) error {
	return m.adminDAO.DelUser(ctx, userID)
}

func (m *AdminService) GetUserStatus(ctx context.Context, userID int) (status.UserStatus, error) {
	return m.adminDAO.GetUserStatus(ctx, userID)
}

func (m *AdminService) SetUserStatus(ctx context.Context, userID int, s status.UserStatus) error {
	return m.adminDAO.SetUserStatus(ctx, userID, s)
}

func (m *AdminService) SetMnemonic(ctx context.Context, mn string) error {
	return m.adminRDB.SetMnemonic(ctx, mn)
}

func (m *AdminService) GetMnemonic(ctx context.Context) (string, error) {
	return m.adminRDB.GetMnemonic(ctx)
}

func (m *AdminService) GetAllMnemonic(ctx context.Context) ([]string, error) {
	return m.adminRDB.GetAllMnemonic(ctx)
}

func (m *AdminService) AddBscApiKey(ctx context.Context, apikey string) error {
	return m.bscApiKeyRDB.ZAdd(ctx, apikey)
}

func (m *AdminService) GetBscApiKey(ctx context.Context) (string, error) {
	return m.bscApiKeyRDB.ZRangeWithScore(ctx)
}

func (m *AdminService) GetAllBscApiKey(ctx context.Context) ([]string, error) {
	return m.bscApiKeyRDB.ZRange(ctx)
}

func (m *AdminService) DelBscApiKey(ctx context.Context, apikey string) error {
	return m.bscApiKeyRDB.ZRem(ctx, apikey)
}

// func (m *AdminService) AddPool(ctx context.Context, pool *model.Pool) error {
// 	return m.adminDAO.AddPool(pool)
// }

// func (m *AdminService) DelPool(ctx context.Context, poolID int) error {
// 	return m.adminDAO.DelPool(poolID)
// }

// func (m *AdminService) UpdatePool(ctx context.Context, pool model.Pool) error {
// 	return m.adminDAO.UpdatePool(pool)
// }

// func (m *AdminService) GetPool(ctx context.Context, poolID int) (*model.Pool, error) {
// 	return m.adminDAO.GetPool(poolID)
// }

// func (m *AdminService) GetAllPools(ctx context.Context, query map[string]any) ([]model.Pool, error) {
// 	return m.adminDAO.GetAllPools(query)
// }

func (m *AdminService) SetFreeGpuNum(ctx context.Context, num int) error {
	return m.adminDAO.SetFreeGpuNum(ctx, num)
}

func (m *AdminService) GetFreeGpuNum(ctx context.Context) (int, error) {
	return m.adminDAO.GetFreeGpuNum(ctx)
}

func (m *AdminService) GetUserOperlogs(ctx context.Context, query map[string]any) ([]model.Operlog, int64, error) {
	return m.adminDAO.GetUserOperlogs(ctx, query)
}

func (m *AdminService) GetUserPointslogs(ctx context.Context, query map[string]any) ([]model.Pointslog, int64, error) {
	return m.adminDAO.GetUserPointslogs(ctx, query)
}

func (m *AdminService) GetUserLoginlogs(ctx context.Context, query map[string]any) ([]model.Loginlog, int64, error) {
	return m.adminDAO.GetUserLoginlogs(ctx, query)
}

func (m *AdminService) GetFarms(ctx context.Context, query map[string]any) ([]model.Farm, int64, error) {
	return m.farmDAO.GetFarms(ctx, query)
}

func (m *AdminService) GetFarmsByUserID(ctx context.Context, userID int, query map[string]any) ([]model.Farm, int64, error) {
	return m.farmDAO.GetFarmsByUserID(ctx, userID, query)
}

func (m *AdminService) GetUserMiners(ctx context.Context, userID int, query map[string]any) ([]model.Miner, int64, error) {
	return m.adminDAO.GetUserMiners(ctx, userID, query)
}

// func (m *AdminService) CreateGlobalFs(ctx context.Context, req *dto.CreateFsReq) error {
// 	return m.adminDAO.CreateGlobalFs()
// }

func (m *AdminService) GetInviteReward(ctx context.Context) (float32, error) {
	return m.adminDAO.GetInviteReward(ctx)
}

func (m *AdminService) SetInviteReward(ctx context.Context, reward float32) error {
	return m.adminDAO.SetInviteReward(ctx, reward)
}

func (m *AdminService) GetRechargeRatio(ctx context.Context) (float32, error) {
	return m.adminDAO.GetRechargeRatio(ctx)
}

func (m *AdminService) SetRechargeRatio(ctx context.Context, ratio float32) error {
	return m.adminDAO.SetRechargeRatio(ctx, ratio)
}

func (m *AdminService) GetRechargeReward(ctx context.Context) (float32, error) {
	return m.adminDAO.GetRechargeReward(ctx)
}

func (m *AdminService) SetRechargeReward(ctx context.Context, reward float32) error {
	return m.adminDAO.SetRechargeReward(ctx, reward)
}

func (m *AdminService) GetSwitchRegister(ctx context.Context) (status.RegisterStatus, error) {
	return m.adminDAO.GetSwitchRegister(ctx)
}

func (m *AdminService) SetSwitchRegister(ctx context.Context, s status.RegisterStatus) error {
	return m.adminDAO.SetSwitchRegister(ctx, s)
}
