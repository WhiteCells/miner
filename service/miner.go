package service

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/common/perm"
	"miner/dao/redis"
	"miner/model/info"
)

type MinerService struct {
	minerRDB *redis.MinerRDB
	farmRDB  *redis.FarmRDB
}

func NewMinerService() *MinerService {
	return &MinerService{
		minerRDB: redis.NewMinerRDB(),
		farmRDB:  redis.NewFarmRDB(),
	}
}

// CreateMiner 创建矿机
func (s *MinerService) CreateMiner(ctx context.Context, req *dto.CreateMinerReq) (*info.Miner, error) {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}
	// 检查用户对矿场的权限
	if !s.validFarmPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
		return nil, errors.New("permission denied")
	}
	// TODO id

	// 创建矿机
	miner := &info.Miner{
		Name: req.Name,
		Perm: perm.MinerOwner,
	}

	// TODO 测试连接

	// 创建矿机
	err := s.minerRDB.Set(ctx, req.FarmID, miner)
	return miner, err
}

// DeleteMiner 删除矿机
func (s *MinerService) DeleteMiner(ctx context.Context, req *dto.DeleteMinerReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	// 检查用户对 Miner 的权限
	if !s.validPerm(ctx, userID, req.MinerID, []perm.MinerPerm{perm.MinerOwner}) {
		return errors.New("permission denied")
	}
	// 删除矿机
	if err := s.minerRDB.Del(ctx, req.FarmID, req.MinerID); err != nil {
		return errors.New("delete miner failed")
	}
	return nil
}

// GetMinerByID 获取矿机信息
func (s *MinerService) GetMinerByID(ctx context.Context, minerID string) (*info.Miner, error) {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}
	miner, err := s.minerRDB.GetByID(ctx, userID, minerID)
	return miner, err
}

// UpdateMiner 更新矿机信息
func (s *MinerService) UpdateMiner(ctx context.Context, req *dto.UpdateMinerReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	if !s.validPerm(ctx, userID, req.MinerID, []perm.MinerPerm{perm.MinerOwner, perm.MinerManager}) {
		return errors.New("permission denied")
	}

	miner, err := s.minerRDB.GetByID(ctx, req.FarmID, req.MinerID)
	if err != nil {
		return errors.New("miner not found")
	}

	// 更新矿机信息
	for key, value := range req.UpdateInfo {
		switch key {
		case "name":
			miner.Name = value.(string)
		}
	}

	// todo 需要测试连接

	// 保存更新
	if err := s.minerRDB.Set(ctx, req.FarmID, miner); err != nil {
		return err
	}

	return nil
}

// GetMiner 获取用户在矿场的所有矿机
func (s *MinerService) GetMiner(ctx context.Context, farmID string) (*[]info.Miner, error) {
	miners, err := s.minerRDB.GetAll(ctx, farmID)
	return miners, err
}

// Transfer 转移矿机
func (s *MinerService) Transfer(ctx context.Context, req *dto.TransferMinerReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	// 权限检查
	if !s.validPerm(ctx, userID, req.MinerID, []perm.MinerPerm{perm.MinerOwner}) {
		return errors.New("permission denied")
	}
	// 转移
	if err := s.minerRDB.Transfer(ctx, userID, req.FromFarmID, req.MinerID, req.ToUserID, req.ToFarmID); err != nil {
		return errors.New("transfer miner failed")
	}
	return nil
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

// ApplyFs 矿机应用飞行表
func (s *MinerService) ApplyFs(ctx context.Context, req *dto.ApplyMinerFlightsheetReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	if !s.validPerm(ctx, userID, req.MinerID, []perm.MinerPerm{perm.MinerOwner, perm.MinerManager}) {
		return errors.New("permission denied")
	}
	return s.minerRDB.ApplyFs(ctx, req.MinerID, req.FlightsheetID)
}

func (s *MinerService) validPerm(ctx context.Context, userID string, minerID string, allowedPerms []perm.MinerPerm) bool {
	farm, err := s.minerRDB.GetByID(ctx, userID, minerID)
	if err != nil {
		return false
	}

	for _, p := range allowedPerms {
		if farm.Perm == p {
			return true
		}
	}

	return false
}

func (s *MinerService) validFarmPerm(ctx context.Context, userID string, farmID string, allowedPerms []perm.FarmPerm) bool {
	farm, err := s.farmRDB.GetByID(ctx, userID, farmID)
	if err != nil {
		return false
	}

	for _, p := range allowedPerms {
		if farm.Perm == p {
			return true
		}
	}

	return false
}
