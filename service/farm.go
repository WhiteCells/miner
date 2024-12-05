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

type FarmService struct {
	farmDAO     *mysql.FarmDAO
	userFarmDAO *mysql.UserFarmDAO
	farmCache   *redis.FarmCache
}

func NewFarmService() *FarmService {
	return &FarmService{
		farmDAO:     mysql.NewFarmDAO(),
		userFarmDAO: mysql.NewUserFarmDAO(),
		farmCache:   redis.NewFarmCache(),
	}
}

// 创建矿场
func (s *FarmService) CreateFarm(ctx context.Context, req *dto.CreateFarmReq) (int, error) {
	farm := &model.Farm{
		Name:     req.Name,
		TimeZone: req.TimeZone,
	}

	// 创建矿场
	farmID, err := s.farmDAO.CreateFarm(farm)
	if err != nil {
		return -1, err
	}

	// 创建用户-矿场关联
	userFarm := &model.UserFarm{
		UserID: ctx.Value("user_id").(int),
		FarmID: farmID,
		Role:   perm.FarmOwner,
	}

	return farmID, s.userFarmDAO.CreateUserFarm(userFarm)
}

// 获取矿场信息
func (s *FarmService) GetFarmInfo(ctx context.Context, req *dto.GetFarmInfoReq) (*model.Farm, error) {
	// 缓存获取
	farm, err := s.farmCache.GetFarmInfo(ctx, req.FarmID)
	if err == nil {
		return farm, nil
	}

	// 缓存未命中，数据库获取
	farm, err = s.farmDAO.GetFarmByID(req.FarmID)
	if err != nil {
		return nil, err
	}

	// 更新缓存
	if err := s.farmCache.SetFarmInfo(ctx, farm); err != nil {
		return nil, err
	}

	return farm, nil
}

// 更新矿场信息
func (s *FarmService) UpdateFarm(ctx context.Context, userID, farmID int, updates map[string]interface{}) error {
	// 检查权限
	if !s.checkFarmPermission(userID, farmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
		return errors.New("permission denied")
	}

	farm, err := s.farmDAO.GetFarmByID(farmID)
	if err != nil {
		return err
	}

	// 更新矿场信息
	for key, value := range updates {
		switch key {
		case "name":
			farm.Name = value.(string)
		case "time_zone":
			farm.TimeZone = value.(string)
		}
	}

	// 保存更新
	if err := s.farmDAO.UpdateFarm(farm); err != nil {
		return err
	}

	// 清除缓存
	return s.farmCache.DeleteFarmCache(ctx, farmID)
}

// 转移矿场所有权
func (s *FarmService) TransferFarmOwnership(ctx context.Context, req *dto.TransferMinerReq) error {
	// 检查当前用户是否是矿场所有者
	if !s.checkFarmPermission(req.FromUserID, req.FarmID, []perm.FarmPerm{perm.FarmOwner}) {
		return errors.New("permission denied")
	}
	return s.userFarmDAO.TransferFarmOwnership(req.FromUserID, req.ToUserID, req.FarmID)
}

// 获取用户的所有矿场
func (s *FarmService) GetUserAllFarm(ctx context.Context, userID int) ([]model.Farm, error) {
	return s.farmDAO.GetUserAllFarm(userID)
}

// 添加矿场成员
func (s *FarmService) AddFarmMember(ctx context.Context, userID, farmID, memberID int, role perm.FarmPerm) error {
	// 检查权限
	if !s.checkFarmPermission(userID, farmID, []perm.FarmPerm{perm.FarmOwner}) {
		return errors.New("permission denied")
	}

	if !isValidRole(role) {
		return errors.New("invalid role")
	}

	return s.userFarmDAO.CreateUserFarm(&model.UserFarm{
		UserID: memberID,
		FarmID: farmID,
		Role:   role,
	})
}

// 移除矿场成员
func (s *FarmService) RemoveFarmMember(ctx context.Context, userID, farmID, memberID int) error {
	// 检查权限
	if !s.checkFarmPermission(userID, farmID, []perm.FarmPerm{perm.FarmOwner}) {
		return errors.New("permission denied")
	}

	memberRole, err := s.userFarmDAO.GetUserFarmRole(memberID, farmID)
	if err != nil {
		return err
	}

	// 所有者不移除
	if memberRole == perm.FarmOwner {
		return errors.New("cannot remove farm owner")
	}

	return s.userFarmDAO.DeleteUserFarm(memberID, farmID)
}

// 检查用户对矿场的权限
func (s *FarmService) checkFarmPermission(userID, farmID int, allowedRoles []perm.FarmPerm) bool {
	role, err := s.userFarmDAO.GetUserFarmRole(userID, farmID)
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

// 检查角色是否有效
func isValidRole(role perm.FarmPerm) bool {
	validRoles := map[perm.FarmPerm]bool{
		"owner":   true,
		"manager": true,
		"viewer":  true,
	}
	return validRoles[role]
}
