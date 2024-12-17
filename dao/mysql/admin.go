package mysql

import (
	"miner/common/status"
	"miner/model"
	"miner/utils"
)

type AdminDAO struct{}

func NewAdminDAO() *AdminDAO {
	return &AdminDAO{}
}

// GetAllUser 获取所有用户信息
func (dao *AdminDAO) GetAllUsers(query map[string]interface{}) (*[]model.User, int64, error) {
	var users []model.User
	var total int64

	db := utils.DB.Model(&model.User{})

	// query 的其他参数

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	err := db.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("id"). // 目前用 ID，后续有需求在修改
		Find(&users).Error

	return &users, total, err
}

// GetUserOperLogs 获取用户日志
func (dao *AdminDAO) GetUserOperLogs(query map[string]interface{}) (*[]model.OperLog, int64, error) {
	var logs []model.OperLog
	var total int64

	db := utils.DB.Model(&model.OperLog{})

	// query 的其他参数

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	err := db.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("time"). // 目前用 time，后续有需求在修改
		Find(&logs).Error

	return &logs, total, err
}

// GetUserLoginLogs 获取用户登陆日志
func (dao *AdminDAO) GetUserLoginLogs(query map[string]interface{}) (*[]model.LoginLog, int64, error) {
	var logs []model.LoginLog
	var total int64

	db := utils.DB.Model(&model.LoginLog{})

	// query 的其他参数

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	err := db.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("time"). // 目前用 time，后续有需求在修改
		Find(&logs).Error

	return &logs, total, err
}

// GetUserPointsRecords 获取用户积分记录
func (dao *AdminDAO) GetUserPointsRecords(query map[string]interface{}) (*[]model.PointsRecord, int64, error) {
	var records []model.PointsRecord
	var total int64

	db := utils.DB.Model(&model.LoginLog{})

	// query 的其他参数

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, -1, err
	}

	err := db.Offset((pageNum - 1) * pageSize).
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
	// pageNum := query["page_num"].(int)
	// pageSize := query["page_size"].(int)

	// 获取用户拥有的矿场数量
	// 后续可以细分为：用户拥有，用户管理，用户查看
	if err := utils.DB.Model(model.UserFarm{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, -1, err
	}

	// 查询矿场详情
	// if err := utils.DB.Joins("JOIN user_farm ON user_farm.farm_id = farm.id").
	// 	Where("user_farm.user_id = ?", userID).
	// 	Find(&farms).Error; err != nil {
	// 	return nil, -1, err
	// }

	return &farms, total, nil
}

// GetUserMiners 获取用户的矿机
func (dao *AdminDAO) GetUserMiners(query map[string]interface{}) (*[]model.Miner, int64, error) {
	var miners []model.Miner
	var total int64

	userID := query["user_id"].(int)
	// pageNum := query["page_num"].(int)
	// pageSize := query["page_size"].(int)

	if err := utils.DB.Model(model.UserMiner{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, -1, err
	}

	return &miners, total, nil
}

// SwitchRegister 切换注册功能
func (dao *AdminDAO) SwitchRegister(status status.RegisterStatus) error {

	return nil
}

// SetGlobalFlightsheet 设置全局飞行表，所有用户都进行加载
func (dao *AdminDAO) SetGlobalFlightsheet(fs *model.Flightsheet) error {
	return nil
}

// SetInviteReward 设置邀请积分奖励
func (dao *AdminDAO) SetInviteReward(reward int) error {
	return nil
}

// SetRechargeReward 设置充值积分奖励
func (dao *AdminDAO) SetRechargeReward(reward int) error {
	return nil
}

// SetUserStatus 设置用户状态
func (dao *AdminDAO) SetUserStatus(status status.UserStatus) error {
	return nil
}

// SetMinerPoolCost 设置矿池消耗
func (dao *AdminDAO) SetMinerPoolCost(minePoolID int, cost float64) error {
	return nil
}
