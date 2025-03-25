package mysql

import (
	"context"
	"errors"
	"miner/common/status"
	"miner/model"
	"miner/utils"
)

type AdminDAO struct{}

func NewAdminDAO() *AdminDAO {
	return &AdminDAO{}
}

// 删除用户
func (dao *AdminDAO) DelUser(ctx context.Context, userID int) error {
	if err := utils.DB.WithContext(ctx).Delete(&model.User{}, userID).Error; err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}

// 获取用户
func (dao *AdminDAO) GetUsers(ctx context.Context, query map[string]any) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	db := utils.DB.WithContext(ctx).Model(&model.User{})

	// query 的其他参数

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 查询总数
	if err := db.Count(&total).Error; err != nil {
		return nil, -1, errors.New("no user found")
	}

	// 分页查询
	if err := db.
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("id"). // 目前用 ID，后续有需求在修改
		Find(&users).Error; err != nil {
		return nil, -1, errors.New("no user found")
	}

	return users, total, nil
}

// 获取所有用户
func (dao *AdminDAO) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	if err := utils.DB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, errors.New("")
	}
	return users, nil
}

// 获取用户状态
func (m *AdminDAO) GetUserStatus(ctx context.Context, userID int) (status.UserStatus, error) {
	var user model.User
	if err := utils.DB.WithContext(ctx).
		Where("id=?", userID).
		Error; err != nil {
		return status.UserNone, errors.New("")
	}
	return user.Status, nil
}

// 设置免费 GPU 数量
func (m *AdminDAO) SetFreeGpuNum(ctx context.Context, num int) error {
	var system model.System
	if err := utils.DB.WithContext(ctx).
		Find(&system).Error; err != nil {
		return errors.New("")
	}
	system.FreeGpuNum = num
	if err := utils.DB.WithContext(ctx).Save(&system).Error; err != nil {
		return errors.New("")
	}
	return nil
}

// 获取免费 GPU 数量
func (m *AdminDAO) GetFreeGpuNum(ctx context.Context) (int, error) {
	var system model.System
	if err := utils.DB.WithContext(ctx).Find(&system).Error; err != nil {
		return -1, errors.New("")
	}
	return system.FreeGpuNum, nil
}

// 获取用户日志
func (dao *AdminDAO) GetUserOperlogs(ctx context.Context, query map[string]any) ([]model.Operlog, int64, error) {
	var logs []model.Operlog
	var total int64

	db := utils.DB.WithContext(ctx).Model(&model.Operlog{})

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 查询总数
	if err := db.Count(&total).Error; err != nil {
		return nil, -1, errors.New("")
	}

	// 分页查询
	if err := db.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&logs).Error; err != nil {
		return nil, -1, errors.New("")
	}

	return logs, total, nil
}

// 获取用户登陆日志
func (dao *AdminDAO) GetUserLoginlogs(ctx context.Context, query map[string]any) ([]model.Loginlog, int64, error) {
	var logs []model.Loginlog
	var total int64

	db := utils.DB.WithContext(ctx).Model(&model.Loginlog{})

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 查询总数
	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	// 分页查询
	if err := db.
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&logs).Error; err != nil {
		return nil, -1, errors.New("")
	}

	return logs, total, nil
}

// 获取用户积分记录
func (dao *AdminDAO) GetUserPointslogs(ctx context.Context, query map[string]any) ([]model.Pointslog, int64, error) {
	var records []model.Pointslog
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 查询总数
	if err := utils.DB.WithContext(ctx).
		Where(model.Pointslog{}).
		Count(&total).Error; err != nil {
		return nil, -1, errors.New("")
	}

	// 分页查询
	if err := utils.DB.WithContext(ctx).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&records).Error; err != nil {
		return nil, -1, errors.New("")
	}

	return records, total, nil
}

// 获取用户的矿场
func (dao *AdminDAO) GetFarms(ctx context.Context, query map[string]any) ([]model.Farm, int64, error) {
	var farms []model.Farm
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	db := utils.DB.WithContext(ctx).Model(&model.Farm{})

	if err := db.Count(&total).Error; err != nil {
		return nil, -1, errors.New("")
	}

	if err := db.
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&farms).Error; err != nil {
		return nil, -1, errors.New("")
	}

	return farms, total, nil
}

// 获取用户的矿机
func (dao *AdminDAO) GetUserMiners(ctx context.Context, userID int, query map[string]any) ([]model.Miner, int64, error) {
	var miners []model.Miner
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	db := utils.DB.WithContext(ctx).
		Model(model.Miner{}).
		Joins("JOIN user_miner ON user_miner.user_id=miner.id").
		Where("user_miner.user_id=?", userID)

	if err := db.
		Count(&total).Error; err != nil {
		return nil, -1, errors.New("")
	}

	if err := db.
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&miners).Error; err != nil {
		return nil, -1, errors.New("")
	}

	return miners, total, nil
}

// 设置用户状态
func (dao *AdminDAO) SetUserStatus(ctx context.Context, userID int, status status.UserStatus) error {
	var user model.User
	if err := utils.DB.WithContext(ctx).First(&user, userID).Error; err != nil {
		return errors.New("")
	}

	user.Status = status
	if err := utils.DB.WithContext(ctx).Save(user).Error; err != nil {
		return errors.New("")
	}

	return nil
}

func (m *AdminDAO) CreateGlobalFs(ctx context.Context, fs *model.Fs) error {
	return nil
}

func (m *AdminDAO) DeleteGlobalFs(ctx context.Context, fsID int) error {
	return nil
}

func (m *AdminDAO) UpdateGlobalFs(ctx context.Context, fsID int) error {
	return nil
}

func (m *AdminDAO) GetGlobalFs(ctx context.Context) ([]model.Fs, error) {
	var globalFs []model.Fs
	err := utils.DB.WithContext(ctx).
		Where("is_global=?", 1).
		Find(&globalFs).Error
	return globalFs, err
}

// 获取充值返现
func (m *AdminDAO) GetInviteReward(ctx context.Context) (float32, error) {
	var system model.System
	if err := utils.DB.WithContext(ctx).First(&system).Error; err != nil {
		return -1, errors.New("")
	}
	return system.InviteReward, nil
}

// 设置充值返现
func (m *AdminDAO) SetInviteReward(ctx context.Context, reward float32) error {
	var system model.System
	if err := utils.DB.WithContext(ctx).First(&system).Error; err != nil {
		return errors.New("")
	}
	system.InviteReward = reward
	return utils.DB.WithContext(ctx).Save(system).Error
}

// 获取充值比率
func (m *AdminDAO) GetRechargeRatio(ctx context.Context) (float32, error) {
	var system model.System
	if err := utils.DB.WithContext(ctx).First(&system).Error; err != nil {
		return -1, errors.New("")
	}
	return system.RechargeRatio, nil
}

// 设置充值比率
func (m *AdminDAO) SetRechargeRatio(ctx context.Context, ratio float32) error {
	var system model.System
	if err := utils.DB.WithContext(ctx).First(&system).Error; err != nil {
		return errors.New("")
	}
	system.RechargeRatio = ratio
	if err := utils.DB.WithContext(ctx).Save(system).Error; err != nil {
		return errors.New("")
	}
	return nil
}

// 获取充值返现
func (m *AdminDAO) GetRechargeReward(ctx context.Context) (float32, error) {
	var system model.System
	if err := utils.DB.WithContext(ctx).First(&system).Error; err != nil {
		return -1, errors.New("")
	}
	return system.InviteReward, nil
}

// 设置充值返现
func (m *AdminDAO) SetRechargeReward(ctx context.Context, reward float32) error {
	var system model.System
	if err := utils.DB.WithContext(ctx).First(&system).Error; err != nil {
		return errors.New("")
	}
	system.RechargeReward = reward
	if err := utils.DB.WithContext(ctx).Save(system).Error; err != nil {
		return errors.New("")
	}
	return nil
}

// 获取注册开关
func (m *AdminDAO) GetSwitchRegister(ctx context.Context) (status.RegisterStatus, error) {
	var system model.System
	if err := utils.DB.WithContext(ctx).First(&system).Error; err != nil {
		return status.RegisterNone, errors.New("")
	}
	return system.SwitchRegister, nil
}

// 设置注册开关
func (m *AdminDAO) SetSwitchRegister(ctx context.Context, s status.RegisterStatus) error {
	var system model.System
	if err := utils.DB.WithContext(ctx).First(&system).Error; err != nil {
		return errors.New("")
	}
	system.SwitchRegister = s
	if err := utils.DB.WithContext(ctx).Save(system).Error; err != nil {
		return errors.New("")
	}
	return nil
}
