package mysql

import (
	"miner/common/perm"
	"miner/model"
	"miner/utils"

	"gorm.io/gorm"
)

type FarmDAO struct{}

func NewFarmDAO() *FarmDAO {
	return &FarmDAO{}
}

// CreateFarm 创建矿场
func (dao *FarmDAO) CreateFarm(farm *model.Farm, userID int) error {
	// 创建矿场时就需要将用户与矿场关联
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 创建矿场
		if err := tx.Create(farm).Error; err != nil {
			return err
		}
		// 创建 用户-矿场 关联
		userFarm := &model.UserFarm{
			UserID: userID,
			FarmID: farm.ID,
			Perm:   perm.FarmOwner,
		}
		if err := tx.Create(userFarm).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteFarmByID 删除矿场
func (dao *FarmDAO) DeleteFarmByID(farmID int) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除 用户-矿场 关联
		if err := tx.Where("farm_id = ?", farmID).Delete(&model.UserFarm{}).Error; err != nil {
			return err
		}
		// 删除 矿场-矿机 关联
		// TODO 矿机如何处理
		if err := tx.Where("farm_id = ?", farmID).Delete(&model.FarmMiner{}).Error; err != nil {
			return err
		}
		// 删除矿场
		if err := tx.Delete(&model.Farm{}, farmID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateFarm 更新矿场信息
func (dao *FarmDAO) UpdateFarm(farm *model.Farm) error {
	// TODO 如果更新飞行表后，需要更新关联
	return utils.DB.Save(farm).Error
}

// GetUserAllFarm 获取用户的所有矿场
func (dao *FarmDAO) GetUserAllFarm(userID int) (*[]model.Farm, error) {
	var farms []model.Farm
	err := utils.DB.Joins("JOIN user_farm ON farm.id = user_farm.farm_id").
		Where("user_farm.user_id = ?", userID).
		Find(&farms).Error
	return &farms, err
}

// GetFarmByID 获取指定矿场
func (dao *FarmDAO) GetFarmByID(farmID int) (*model.Farm, error) {
	var farm model.Farm
	err := utils.DB.First(&farm, farmID).Error
	return &farm, err
}

// ApplyFlightSheet 应用 Flightsheet
func (dao *FarmDAO) ApplyFlightSheet(farmID int, fs_id int) error {

	return nil
}
