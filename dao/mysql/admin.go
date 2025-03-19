package mysql

import (
	"miner/common/status"
	"miner/model"
	"miner/model/relation"
	"miner/utils"
)

type AdminDAO struct{}

func NewAdminDAO() *AdminDAO {
	return &AdminDAO{}
}

// GetAllUser 获取所有用户信息
func (dao *AdminDAO) GetAllUsers(query map[string]any) (*[]model.User, int64, error) {
	var users []model.User
	var total int64

	db := utils.DB.Model(&model.User{})

	// query 的其他参数

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 查询总数
	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	// 分页查询
	err := db.
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("id"). // 目前用 ID，后续有需求在修改
		Find(&users).Error

	return &users, total, err
}

// GetUserOperLogs 获取用户日志
func (dao *AdminDAO) GetUserOperlogs(query map[string]any) (*[]model.Operlog, int64, error) {
	var logs []model.Operlog
	var total int64

	db := utils.DB.Model(&model.Operlog{})

	// query 的其他参数

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 查询总数
	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	// 分页查询
	err := db.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("time"). // 目前用 time，后续有需求在修改
		Find(&logs).Error

	return &logs, total, err
}

// GetUserLoginLogs 获取用户登陆日志
func (dao *AdminDAO) GetUserLoginlogs(query map[string]any) (*[]model.Loginlog, int64, error) {
	var logs []model.Loginlog
	var total int64

	db := utils.DB.Model(&model.Loginlog{})

	// query 的其他参数

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 查询总数
	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	// 分页查询
	err := db.
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("time"). // 目前用 time，后续有需求在修改
		Find(&logs).Error

	return &logs, total, err
}

// GetUserPointsRecords 获取用户积分记录
func (dao *AdminDAO) GetUserPointslogs(query map[string]any) (*[]model.Pointslog, int64, error) {
	var records []model.Pointslog
	var total int64

	// query 的其他参数

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 查询总数
	if err := utils.DB.
		Where(model.Pointslog{}).
		Count(&total).Error; err != nil {
		return nil, -1, err
	}

	// 分页查询
	err := utils.DB.
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("time"). // 目前用 time，后续有需求在修改
		Find(&records).Error

	return &records, total, err
}

// GetUserFarms 获取用户的矿场
func (dao *AdminDAO) GetUserFarms(query map[string]interface{}) (*[]model.Farm, int64, error) {
	var farms []model.Farm
	var total int64

	// query 的其他参数

	userID := query["user_id"].(int)
	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 获取用户拥有的矿场数量
	// 后续可以细分为：用户拥有，用户管理，用户查看
	if err := utils.DB.
		Model(relation.UserFarm{}).
		Where("user_id = ?", userID).
		Count(&total).Error; err != nil {
		return nil, -1, err
	}

	// 查询矿场详情
	err := utils.DB.
		Joins("JOIN user_farm ON user_farm.farm_id = farm.id").
		Where("user_farm.user_id = ?", userID).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&farms).Error

	return &farms, total, err
}

// GetUserMiners 获取用户的矿机
func (dao *AdminDAO) GetUserMiners(query map[string]interface{}) (*[]model.Miner, int64, error) {
	var miners []model.Miner
	var total int64

	userID := query["user_id"].(int)
	farmID := query["farm_id"].(int)
	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	if err := utils.DB.
		Model(model.Miner{}).
		Joins("JOIN farm_miner ON farm_miner.miner_id = miner.id").
		Joins("JOIN user_miner ON user_miner.miner_id = miner.id").
		Joins("JOIN user_farm ON user_farm.farm_id = farm_miner.farm_id").
		Where("user_farm.user_id = ? AND user_farm.farm_id = ? AND user_miner.user_id = ?", userID, farmID, userID).
		Count(&total).Error; err != nil {
		return nil, -1, err
	}

	err := utils.DB.
		Model(model.Miner{}).
		Joins("JOIN farm_miner ON farm_miner.miner_id = miner.id").
		Joins("JOIN user_miner ON user_miner.miner_id = miner.id").
		Joins("JOIN user_farm ON user_farm.farm_id = farm_miner.farm_id").
		Where("user_farm.user_id = ? AND user_farm.farm_id = ? AND user_miner.user_id = ?", userID, farmID, userID).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&miners).Error

	return &miners, total, err
}

// SwitchRegister 切换注册功能
func (dao *AdminDAO) SwitchRegister(status status.RegisterStatus) error {
	var system model.System
	if err := utils.DB.First(&system).Error; err != nil {
		return err
	}
	system.SwitchRegister = status

	if err := utils.DB.Save(system).Error; err != nil {
		return err
	}

	return nil
}

// SetGlobalFlightsheet 设置全局飞行表，所有用户都进行加载
func (dao *AdminDAO) SetGlobalFlightsheet(fs *model.Fs) error {
	return nil
}

// SetInviteReward 设置邀请积分奖励
func (dao *AdminDAO) SetInviteReward(reward int) error {
	var system model.System
	if err := utils.DB.First(&system).Error; err != nil {
		return err
	}
	system.InviteReward = reward

	if err := utils.DB.Save(system).Error; err != nil {
		return err
	}

	return nil
}

// SetRechargeReward 设置充值积分奖励
func (dao *AdminDAO) SetRechargeReward(reward int) error {
	var system model.System
	if err := utils.DB.First(&system).Error; err != nil {
		return err
	}

	system.RechargeReward = reward
	if err := utils.DB.Save(system).Error; err != nil {
		return err
	}

	return nil
}

// SetUserStatus 设置用户状态
func (dao *AdminDAO) SetUserStatus(userID int, status status.UserStatus) error {
	var user model.User
	if err := utils.DB.First(&user, userID).Error; err != nil {
		return err
	}

	user.Status = status
	if err := utils.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}

// SetMinerPoolCost 设置矿池消耗
func (dao *AdminDAO) SetMinePoolCost(minePoolID int, cost float64) error {
	var minePool model.Pool
	if err := utils.DB.First(&minePool, minePoolID).Error; err != nil {
		return err
	}

	minePool.Cost = cost
	if err := utils.DB.Save(minePool).Error; err != nil {
		return err
	}

	return nil
}
