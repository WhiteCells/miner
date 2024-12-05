package service

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/common/perm"
	"miner/dao/mysql"
	"miner/dao/redis"
	"miner/model"
)

type MinerService struct {
	minerDAO     *mysql.MinerDAO
	farmMinerDAO *mysql.FarmMinerDAO
	userMinerDAO *mysql.UserMinerDAO
	minerCache   *redis.MinerCache
	farmService  *FarmService
}

func NewMinerService() *MinerService {
	return &MinerService{
		minerDAO:     mysql.NewMinerDAO(),
		farmMinerDAO: mysql.NewFarmMinerDAO(),
		userMinerDAO: mysql.NewUserMinerDAO(),
		minerCache:   redis.NewMinerCache(),
		farmService:  NewFarmService(),
	}
}

// 创建矿机
func (s *MinerService) CreateMiner(ctx context.Context, req *dto.CreateMinerReq) error {
	// 检查用户对矿场的权限
	if !s.farmService.checkFarmPermission(req.UserID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
		return errors.New("permission denied")
	}

	// 创建矿机
	miner := &model.Miner{
		Name:        req.Name,
		IPAddress:   req.IPAddress,
		SSHPort:     req.SSHPort,
		SSHUser:     req.SSHUser,
		SSHPassword: req.SSHPassword,
		Status:      1,
	}

	// 测试连接 todo

	// 创建矿机记录
	if err := s.minerDAO.CreateMiner(miner); err != nil {
		return err
	}

	// 创建矿场-矿机关联
	return s.farmMinerDAO.CreateFarmMiner(&model.FarmMiner{
		MinerID: miner.ID,
		FarmID:  req.FarmID,
	})
}

// 获取矿机信息
func (s *MinerService) GetMinerInfo(ctx context.Context, minerID int) (*model.Miner, error) {
	// 先从缓存获取
	miner, err := s.minerCache.GetMinerInfo(ctx, minerID)
	if err == nil {
		return miner, nil
	}

	// 缓存未命中，从数据库获取
	miner, err = s.minerDAO.GetMinerByID(minerID)
	if err != nil {
		return nil, err
	}

	// 更新缓存
	if err := s.minerCache.SetMinerInfo(ctx, miner); err != nil {
		return nil, err
	}

	return miner, nil
}

// 更新矿机信息
func (s *MinerService) UpdateMiner(ctx context.Context, userID, minerID int, updates map[string]interface{}) error {
	if !s.checkMinerPermission(userID, minerID, []perm.MinerPerm{perm.MinerOwner, perm.MinerManager}) {
		return errors.New("permission denied")
	}

	miner, err := s.minerDAO.GetMinerByID(minerID)
	if err != nil {
		return err
	}

	// 更新矿机信息
	for key, value := range updates {
		switch key {
		case "name":
			miner.Name = value.(string)
		case "ip_address":
			miner.IPAddress = value.(string)
		case "ssh_port":
			miner.SSHPort = value.(int)
		case "ssh_user":
			miner.SSHUser = value.(string)
		case "ssh_password":
			miner.SSHPassword = value.(string)
		case "status":
			miner.Status = value.(int)
		}
	}

	// todo 更新了SSH相关信息，需要测试连接

	// 保存更新
	if err := s.minerDAO.UpdateMiner(miner); err != nil {
		return err
	}

	// 清除缓存
	// 更新缓存
	// return s.minerCache.DeleteMinerCache(ctx, minerID)
	return err
}

// 转移矿机到其他矿场
func (s *MinerService) TransferMiner(ctx context.Context, userID, minerID, fromFarmID, toFarmID int) error {
	// 检查源矿场权限
	if !s.checkMinerPermission(userID, fromFarmID, []perm.MinerPerm{perm.MinerOwner}) {
		return errors.New("permission denied")
	}

	// 检查目标矿场权限
	if !s.farmService.checkFarmPermission(userID, toFarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
		return errors.New("permission denied for target farm")
	}

	// 更新矿场-矿机关联

	// 清除缓存
	// 更新缓存
	return nil
}

// GetMinerStats 获取矿机状态信息
func (s *MinerService) GetMinerStats(ctx context.Context, minerID int) (map[string]interface{}, error) {
	// 先从缓存获取
	stats, err := s.minerCache.GetMinerStats(ctx, minerID)
	if err == nil {
		return stats, nil
	}

	// 获取矿机信息
	// miner, err := s.minerDAO.GetMinerByID(minerID)
	// if err != nil {
	// 	return nil, err
	// }

	// 通过获取矿机状态 todo
	// stats, err = s.getMinerStatsViaSSH(miner)

	// 更新缓存
	if err := s.minerCache.SetMinerStats(ctx, minerID, stats); err != nil {
		return nil, err
	}

	return stats, nil
}

// 应用飞行表到矿机
func (s *MinerService) ApplyFlightSheet(ctx context.Context, userID, minerID, flightSheetID int) error {
	if !s.checkMinerPermission(userID, minerID, []perm.MinerPerm{perm.MinerOwner, perm.MinerManager}) {
		return errors.New("permission denied")
	}

	// 获取矿机信息

	// 获取飞行表信息

	// 应用飞行表配置

	// 更新关联关系
	return nil
}

func (s *MinerService) checkMinerPermission(userID int, minerID int, allowedRoles []perm.MinerPerm) bool {
	role, err := s.userMinerDAO.GetUserMinerRole(userID, minerID)
	if err != nil {
		return false
	}

	for _, allowedRole := range allowedRoles {
		if role == allowedRole {
			return true
		}
	}
	return false
}
