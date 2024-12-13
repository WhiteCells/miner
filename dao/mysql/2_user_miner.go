package mysql

import (
	"miner/common/perm"
	"miner/model"
	"miner/utils"
)

// 用户 - 矿机
type UserMinerDAO struct{}

func NewUserMinerDAO() *UserMinerDAO {
	return &UserMinerDAO{}
}

// 建立用户与矿机的联系
func (dao *UserMinerDAO) CreateUserMiner(userFarm *model.UserMiner) error {
	return utils.DB.Create(userFarm).Error
}

// 删除用户与矿机的联系
func (dao *UserMinerDAO) DeleteUserMiner(userID int, minerID int) error {
	return utils.DB.Where("user_id = ? AND miner_id = ?", userID, minerID).Delete(model.UserMiner{}).Error
}

// GetUserMiner 通过 ID 获取矿机信息
func (dao *UserMinerDAO) GetUserMinerByID(userID int, minerID int) (*model.UserMiner, error) {
	var userMiner model.UserMiner
	err := utils.DB.Where("user_id = ? AND miner_id = ?", userID, minerID).First(&userMiner).Error
	return &userMiner, err
}

// GetUserMinerPerm 获取用户在矿机中的权限
func (dao *UserMinerDAO) GetUserMinerPerm(userID int, minerID int) (perm.MinerPerm, error) {
	var userFarm model.UserMiner
	err := utils.DB.Where("user_id = ? AND miner_id = ?", userID, minerID).First(&userFarm).Error
	return userFarm.Perm, err
}

// GetUserAllMinerInFarm 获取矿场中指定用户的所有矿机
func (dao UserMinerDAO) GetUserAllMinerInFarm(userID int, farmID int) (*[]model.Miner, error) {
	var miners []model.Miner
	err := utils.DB.
		Select("DISTINCT miner.*").
		Joins("JOIN farm_miner ON farm_miner.miner_id = miner.id").
		Joins("JOIN user_farm ON farm_miner.farm_id = user_farm.farm_id").
		Where("user_farm.user_id = ? AND farm_miner.farm_id = ?", userID, farmID).
		Find(&miners).Error
	return &miners, err
}
