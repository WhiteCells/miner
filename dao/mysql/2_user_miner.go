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
func (dao *UserMinerDAO) CreateUserFarm(userFarm *model.UserFarm) error {
	return utils.DB.Create(userFarm).Error
}

// 获取用户在矿场中的权限
func (dao *UserMinerDAO) GetUserMinerRole(userID int, farmID int) (perm.MinerPerm, error) {
	var userFarm model.UserMiner
	err := utils.DB.Where("user_id = ? AND farm_id = ?", userID, farmID).First(&userFarm).Error
	return userFarm.Perm, err
}
