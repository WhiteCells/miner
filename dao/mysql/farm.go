package mysql

import (
	"miner/model"
	"miner/utils"

	"gorm.io/gorm"
)

type FarmDAO struct{}

func NewFarmDAO() *FarmDAO {
	return &FarmDAO{}
}

// 创建矿场
func (dao *FarmDAO) CreateFarm(farm *model.Farm) (int, error) {
	return farm.ID, utils.DB.Create(farm).Error
}

// 获取矿场信息
func (dao *FarmDAO) GetFarmByID(userID int) (*model.Farm, error) {
	var farm model.Farm
	err := utils.DB.First(&farm, userID).Error
	return &farm, err
}

// 获取用户的所有矿场
func (dao *FarmDAO) GetUserAllFarm(userID int) (*[]model.Farm, error) {
	var farms []model.Farm
	err := utils.DB.Joins("JOIN user_farm ON farm.id = user_farm.farm_id").
		Where("user_farm.user_id = ?", userID).
		Find(&farms).Error
	return &farms, err
}

// 更新矿场信息
func (dao *FarmDAO) UpdateFarm(farm *model.Farm) error {
	return utils.DB.Save(farm).Error
}

// 删除矿场
func (dao *FarmDAO) DeleteFarmByID(farmID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除矿场关联
		if err := tx.Where("farm_id = ?", farmID).Delete(&model.UserFarm{}).Error; err != nil {
			return err
		}
		if err := tx.Where("farm_id = ?", farmID).Delete(&model.FarmMiner{}).Error; err != nil {
			return err
		}
		// 删除矿场
		return tx.Delete(&model.Farm{}, farmID).Error
	})
}
