package service

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/common/perm"
	"miner/dao/redis"
	"miner/model/info"
)

type FarmService struct {
	farmRDB *redis.FarmRDB
}

func NewFarmService() *FarmService {
	return &FarmService{
		farmRDB: redis.NewFarmRDB(),
	}
}

// CreateFarm 创建矿场
func (s *FarmService) CreateFarm(ctx context.Context, req *dto.CreateFarmReq) (*info.Farm, error) {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}
	farm := &info.Farm{
		Name:     req.Name,
		TimeZone: req.TimeZone,
	}

	// 创建矿场
	err := s.farmRDB.Set(ctx, userID, farm, perm.FarmOwner)

	return farm, err
}

// DeleteFarm 删除矿场
func (s *FarmService) DeleteFarm(ctx context.Context, req *dto.DeleteFarmReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	// 检查用户对矿场的权限
	if !s.validPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner}) {
		return errors.New("permission denied")
	}
	// 删除
	if err := s.farmRDB.Del(ctx, userID, req.FarmID); err != nil {
		return errors.New("delete farm failed")
	}
	return nil
}

// UpdateFarm 更新矿场信息
func (s *FarmService) UpdateFarm(ctx context.Context, req *dto.UpdateFarmReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	// 权限
	if !s.validPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
		return errors.New("permission denied")
	}

	// 查找矿场
	farm, err := s.farmRDB.GetByID(ctx, userID, req.FarmID)
	if err != nil {
		return err
	}

	for key, value := range req.UpdateInfo {
		switch key {
		case "name":
			name := value.(string)
			if name == "" || len(name) > 100 {
				return errors.New("invalid farm name")
			}
			farm.Name = name
		case "time_zone":
			timeZone := value.(string)
			if timeZone == "" {
				return errors.New("invalid farm time zone")
			}
			farm.TimeZone = timeZone
		}
	}

	// 更新
	return s.farmRDB.Set(ctx, userID, farm, farm.Perm)
}

// GetFarm 获取用户的所有矿场信息
func (s *FarmService) GetFarm(ctx context.Context) (*[]info.Farm, error) {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}
	return s.farmRDB.GetAll(ctx, userID)
}

// GetFarmByID 通过 ID 获取矿场信息
func (s *FarmService) GetFarmByID(ctx context.Context, farmID string) (*info.Farm, error) {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}
	return s.farmRDB.GetByID(ctx, userID, farmID)
}

// ApplyFs 矿场应用飞行表
func (s *FarmService) ApplyFs(ctx context.Context, req *dto.ApplyFarmFlightsheetReq) error {
	// userID, exists := ctx.Value("user_id").(string)
	// if !exists {
	// 	return errors.New("invalid user_id in context")
	// }
	// if err := s.farmRDB.ApplyFs(ctx, req.FarmID, req.FlightsheetID); err != nil {
	// 	return errors.New("farm apply flightsheet faild")
	// }
	return nil
}

// Transfer 转移矿场所有权
func (s *FarmService) Transfer(ctx context.Context, req *dto.TransferFarmReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	// 检查权限
	if !s.validPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner}) {
		return errors.New("permission denied")
	}
	if err := s.farmRDB.Transfer(ctx, req.FarmID, userID, req.ToUserID); err != nil {
		return errors.New("transfer farm failed")
	}
	return nil
}

// AddMember 添加矿场成员
func (s *FarmService) AddMember(ctx context.Context, farmID string, memID string, permission perm.FarmPerm) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	// 检查权限
	if !s.validPerm(ctx, userID, farmID, []perm.FarmPerm{perm.FarmOwner}) {
		return errors.New("permission denied")
	}
	// 添加成员
	return s.farmRDB.AddMember(ctx, userID, farmID, memID)
}

// DelMember 删除矿场成员
func (s *FarmService) DelFarmMember(ctx context.Context, userID, farmID, memID string) error {
	// 检查权限
	if !s.validPerm(ctx, userID, farmID, []perm.FarmPerm{perm.FarmOwner}) {
		return errors.New("permission denied")
	}
	return s.farmRDB.DelMember(ctx, userID, farmID, memID)
}

// validPerm 检查用户对矿场的权限
func (s *FarmService) validPerm(ctx context.Context, userID string, farmID string, allowedPerms []perm.FarmPerm) bool {
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
