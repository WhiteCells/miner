package mysql

import (
	"miner/common/perm"
	"miner/model"
	"miner/utils"
)

// 矿场 - 矿机
type FarmMinerDAO struct{}

func NewFarmMinerDAO() *FarmMinerDAO {
	return &FarmMinerDAO{}
}

// 建立矿场与矿机的联系
func (dao *FarmMinerDAO) CreateFarmMiner(farmMiner *model.FarmMiner) error {
	return utils.DB.Create(farmMiner).Error
}

// todo 更新用户在矿机中的权限
func (dao *FarmMinerDAO) UpdateUserMinerRole(userID int, minerID int, perm perm.MinerPerm) error {
	return nil
}
