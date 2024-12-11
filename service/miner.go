package service

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/common/perm"
	"miner/common/status"
	"miner/dao/mysql"
	"miner/dao/redis"
	"miner/model"
)

type MinerService struct {
	minerDAO     *mysql.MinerDAO
	minerCache   *redis.MinerCache
	userFarmDAO  *mysql.UserFarmDAO
	farmMinerDAO *mysql.FarmMinerDAO
	userMinerDAO *mysql.UserMinerDAO
	farmService  *FarmService
}

func NewMinerService() *MinerService {
	return &MinerService{
		minerDAO:     mysql.NewMinerDAO(),
		minerCache:   redis.NewMinerCache(),
		userFarmDAO:  mysql.NewUserFarmDAO(),
		farmMinerDAO: mysql.NewFarmMinerDAO(),
		userMinerDAO: mysql.NewUserMinerDAO(),
		farmService:  NewFarmService(),
	}
}

// 创建矿机
func (s *MinerService) CreateMiner(ctx context.Context, req *dto.CreateMinerReq) (*model.Miner, error) {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}
	// 检查用户对矿场的权限
	if !s.farmService.checkFarmPermission(userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
		return nil, errors.New("permission denied")
	}

	// 创建矿机
	miner := &model.Miner{
		Name:        req.Name,
		IP:          req.IP,
		SSHPort:     req.SSHPort,
		SSHUser:     req.SSHUser,
		SSHPassword: req.SSHPassword,
		Status:      status.MinerOn,
	}

	// TODO 测试连接

	// 创建矿机
	if err := s.minerDAO.CreateMiner(miner, userID, req.FarmID); err != nil {
		return nil, err
	}

	return miner, nil
}

// 删除矿机
func (s *MinerService) DeleteMiner(ctx context.Context, req *dto.DeleteMinerReq) error {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	// 检查用户对 Miner 的权限
	if !s.checkMinerPermission(userID, req.MinerID, []perm.MinerPerm{perm.MinerOwner}) {
		return errors.New("permission denied")
	}

	if err := s.minerDAO.DeleteMiner(req.MinerID, req.FarmID, userID); err != nil {
		return errors.New("delete miner failed")
	}

	return nil
}

// 获取矿机信息
func (s *MinerService) GetMinerInfo(ctx context.Context, minerID int) (*model.Miner, error) {
	// 先从缓存获取
	miner, err := s.minerCache.GetMinerInfoByID(ctx, minerID)
	if err == nil {
		return miner, nil
	}

	// 缓存未命中，从数据库获取
	miner, err = s.minerDAO.GetMinerByID(minerID)
	if err != nil {
		return nil, err
	}

	// 更新缓存
	if err := s.minerCache.SetMinerInfoByID(ctx, miner); err != nil {
		return nil, err
	}

	return miner, nil
}

// 更新矿机信息
func (s *MinerService) UpdateMiner(ctx context.Context, req *dto.UpdateMinerReq) error {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	if !s.checkMinerPermission(userID, req.MinerID, []perm.MinerPerm{perm.MinerOwner, perm.MinerManager}) {
		return errors.New("permission denied")
	}

	miner, err := s.minerDAO.GetMinerByID(req.MinerID)
	if err != nil {
		return errors.New("miner not found")
	}

	// 更新矿机信息
	for key, value := range req.UpdateInfo {
		switch key {
		case "name":
			miner.Name = value.(string)
		case "ip_address":
			miner.IP = value.(string)
		case "ssh_port":
			miner.SSHPort = value.(int)
		case "ssh_user":
			miner.SSHUser = value.(string)
		case "ssh_password":
			miner.SSHPassword = value.(string)
		case "status":
			miner.Status = value.(status.MinerStatus)
		}
	}

	// todo 需要测试连接

	// 保存更新
	if err := s.minerDAO.UpdateMiner(miner); err != nil {
		return err
	}

	// 更新缓存
	// return s.minerCache.DeleteMinerCache(ctx, minerID)
	return err
}

// 获取用户在矿场的所有矿机
func (s *MinerService) GetUserAllMinerInFarm(ctx context.Context, req *dto.GetUserAllMinerInFarmReq) (*[]model.Miner, error) {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}
	// 缓存
	// 数据库
	return s.userMinerDAO.GetUserAllMinerInFarm(userID, req.FarmID)
}

// 转移矿机到其他矿场
// func (s *MinerService) TransferMiner(ctx context.Context, userID, minerID, fromFarmID, toFarmID int) error {
// 	// 检查源矿场权限
// 	if !s.checkMinerPermission(userID, fromFarmID, minerID, []perm.MinerPerm{perm.MinerOwner}) {
// 		return errors.New("permission denied")
// 	}

// 	// 检查目标矿场权限
// 	if !s.farmService.checkFarmPermission(userID, toFarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
// 		return errors.New("permission denied for target farm")
// 	}

// 	// 更新矿场-矿机关联

// 	// 清除缓存
// 	// 更新缓存
// 	return nil
// }

// GetMinerStats 获取矿机状态信息
func (s *MinerService) GetMinerStats(ctx context.Context, minerID int) (map[string]interface{}, error) {
	// 先从缓存获取

	// 获取矿机信息
	// miner, err := s.minerDAO.GetMinerByID(minerID)
	// if err != nil {
	// 	return nil, err
	// }

	// 通过获取矿机状态 todo
	// stats, err = s.getMinerStatsViaSSH(miner)
	return nil, nil
}

// 应用飞行表到矿机
func (s *MinerService) ApplyFlightSheet(ctx context.Context, req *dto.ApplyMinerFlightsheetReq) error {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	if !s.checkMinerPermission(userID, req.MinerID, []perm.MinerPerm{perm.MinerOwner, perm.MinerManager}) {
		return errors.New("permission denied")
	}

	// 获取矿机信息

	// 获取飞行表信息

	// 应用飞行表配置

	// 更新关联关系
	return nil
}

func (s *MinerService) checkMinerPermission(userID int, minerID int, allowedPerms []perm.MinerPerm) bool {
	// user 对 farm 的权限
	// userFarmPerm, err := s.userFarmDAO.GetUserFarmPerm(userID, FarmID)
	// if err != nil {
	// 	return false
	// }
	// hasUserFarmPerm := false
	// for _, allowedPerm := range allowedPerms {
	// 	if perm.Perm(userFarmPerm) == perm.Perm(allowedPerm) {
	// 		hasUserFarmPerm = true
	// 	}
	// }
	// if !hasUserFarmPerm {
	// 	return false
	// }

	// user 对 miner 的权限
	userMinerPerm, err := s.userMinerDAO.GetUserMinerPerm(userID, minerID)
	if err != nil {
		return false
	}
	for _, allowedPerm := range allowedPerms {
		if userMinerPerm == allowedPerm {
			return true
		}
	}
	return false
}
