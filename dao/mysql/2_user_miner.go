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

// 建立用户与矿场的联系
func (dao *UserMinerDAO) CreateUserMiner(userFarm *model.UserMiner) error {
	return utils.DB.Create(userFarm).Error
}

// 获取用户在矿机中的权限
func (dao *UserMinerDAO) GetUserMinerPerm(userID int, minerID int) (perm.MinerPerm, error) {
	var userFarm model.UserMiner
	err := utils.DB.Where("user_id = ? AND miner_id = ?", userID, minerID).First(&userFarm).Error
	return userFarm.Perm, err
}

// 获取矿场中指定用户的所有矿机
func (dao UserMinerDAO) GetUserAllMinerInFarm(userID int, farmID int) (*[]model.Miner, error) {
	var miners []model.Miner
	err := utils.DB.Joins("JOIN farm_miner ON miner.id = farm_miner.miner_id").
		Joins("JOIN user_farm ON farm_miner.farm_id = user_farm.farm_id").
		Where("user_farm.user_id = ? AND farm_miner.farm_id = ?", userID, farmID).
		Find(&miners).Error
	return &miners, err
}
