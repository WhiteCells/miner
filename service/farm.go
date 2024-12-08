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

// CreateFarm 创建矿场
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
		Perm:   perm.FarmOwner,
	}

	return farmID, s.userFarmDAO.CreateUserFarm(userFarm)
}

// GetFarmInfo 获取矿场信息
func (s *FarmService) GetFarmInfo(ctx context.Context, farmID int) (*model.Farm, error) {
	// 缓存获取
	farm, err := s.farmCache.GetFarmInfo(ctx, farmID)
	if err == nil {
		return farm, nil
	}

	// 缓存未命中，数据库获取
	farm, err = s.farmDAO.GetFarmByID(farmID)
	if err != nil {
		return nil, err
	}

	// 更新缓存
	if err := s.farmCache.SetFarmInfo(ctx, farm); err != nil {
		return nil, err
	}

	return farm, nil
}

// GetAllFarmInfo 获取所有矿场信息
func (s *FarmService) GetUserAllFarmInfo(ctx context.Context) (*[]model.Farm, error) {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return nil, errors.New("invalid user_id in context")
	}
	return s.farmDAO.GetUserAllFarm(userID)
}

// UpdateFarm 更新矿场信息
func (s *FarmService) UpdateFarm(ctx context.Context, req *dto.UpdateFarmReq) error {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	if !s.checkFarmPermission(userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
		return errors.New("permission denied")
	}

	// 查找矿场
	farm, err := s.farmDAO.GetFarmByID(req.FarmID)
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

	// todo bug，当缓存更新失败时，数据库回滚
	// 更新数据库
	if err := s.farmDAO.UpdateFarm(farm); err != nil {
		return err
	}

	// 更新缓存
	if err = s.farmCache.SetFarmInfo(ctx, farm); err != nil {
		return err
	}

	return nil
}

// DeleteFarm 删除矿场
func (s *FarmService) DeleteFarm(ctx context.Context, req *dto.DeleteFarmReq) error {
	// 检查用户对矿场的权限
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	if !s.checkFarmPermission(userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner}) {
		return errors.New("permission denied")
	}
	return s.farmDAO.DeleteFarmByID(req.FarmID)
}

// TransferFarmOwnership 转移矿场所有权
func (s *FarmService) TransferFarmOwnership(ctx context.Context, req *dto.TransferMinerReq) error {
	// 检查当前用户是否是矿场所有者
	if !s.checkFarmPermission(req.FromUserID, req.FarmID, []perm.FarmPerm{perm.FarmOwner}) {
		return errors.New("permission denied")
	}
	return s.userFarmDAO.TransferFarmOwnership(req.FromUserID, req.ToUserID, req.FarmID)
}

// AddFarmMember 添加矿场成员
func (s *FarmService) AddFarmMember(ctx context.Context, userID, farmID, memberID int, permission perm.FarmPerm) error {
	// 检查权限
	if !s.checkFarmPermission(userID, farmID, []perm.FarmPerm{perm.FarmOwner}) {
		return errors.New("permission denied")
	}

	if !isValidPerm(permission) {
		return errors.New("invalid role")
	}

	return s.userFarmDAO.CreateUserFarm(&model.UserFarm{
		UserID: memberID,
		FarmID: farmID,
		Perm:   permission,
	})
}

// RemoveFarmMember 移除矿场成员
func (s *FarmService) RemoveFarmMember(ctx context.Context, userID, farmID, memberID int) error {
	// 检查权限
	if !s.checkFarmPermission(userID, farmID, []perm.FarmPerm{perm.FarmOwner}) {
		return errors.New("permission denied")
	}

	memberRole, err := s.userFarmDAO.GetUserFarmPerm(memberID, farmID)
	if err != nil {
		return err
	}

	// 所有者不移除
	if memberRole == perm.FarmOwner {
		return errors.New("cannot remove farm owner")
	}

	return s.userFarmDAO.DeleteUserFarm(memberID, farmID)
}

// checkFarmPermission 检查用户对矿场的权限
func (s *FarmService) checkFarmPermission(userID, farmID int, allowedRoles []perm.FarmPerm) bool {
	role, err := s.userFarmDAO.GetUserFarmPerm(userID, farmID)
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

// isValidPerm 检查权限是否有效
func isValidPerm(role perm.FarmPerm) bool {
	validRoles := map[perm.FarmPerm]bool{
		perm.FarmOwner:   true,
		perm.FarmManager: true,
		perm.FarmViewer:  true,
	}
	return validRoles[role]
}
