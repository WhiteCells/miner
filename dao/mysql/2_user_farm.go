package mysql

import (
	"miner/common/perm"
	"miner/model"
	"miner/utils"

	"gorm.io/gorm"
)

// 用户 - 矿场
type UserFarmDAO struct{}

func NewUserFarmDAO() *UserFarmDAO {
	return &UserFarmDAO{}
}

// 建立用户与矿场的联系
func (dao *UserFarmDAO) CreateUserFarm(userFarm *model.UserFarm) error {
	return utils.DB.Create(userFarm).Error
}

// 获取用户在矿场的权限
func (dao *UserFarmDAO) GetUserFarmPerm(userID int, farmID int) (perm.FarmPerm, error) {
	var userFarm model.UserFarm
	err := utils.DB.Where("user_id = ? AND farm_id = ?", userID, farmID).First(&userFarm).Error
	return userFarm.Perm, err
}

// 更新用户在矿场的权限
func (dao *UserFarmDAO) UpdateUserFarmRole(userID int, role perm.FarmPerm) error {
	return utils.DB.Model(&model.Farm{}).Where("id = ?", userID).Update("role", role).Error
}

// 删除用户与矿场的联系
func (dao *UserFarmDAO) DeleteUserFarm(userID int, farmID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Where("user_id = ? AND farm_id = ?", userID, farmID).
			Delete(&model.UserFarm{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// 转移所有权
func (dao *UserFarmDAO) TransferFarmOwnership(farmID int, fromUserID int, toUserID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// if err :=
		return nil
	})
}
