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
func (dao *FarmDAO) DeleteFarmByID(farmID int, userID int) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除 用户-矿场 关联
		if err := tx.Delete(&model.UserFarm{}, "user_id = ? AND farm_id = ?", userID, farmID).Error; err != nil {
			return err
		}
		// 删除 矿场-矿机 关联
		// TODO 矿机如何处理
		if err := tx.Delete(&model.FarmMiner{}, "farm_id = ?", farmID).Error; err != nil {
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
	err := utils.DB.
		Joins("JOIN user_farm ON farm.id = user_farm.farm_id").
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

// ApplyFlightsheet 应用 Flightsheet
func (dao *FarmDAO) ApplyFlightsheet(farmID int, fsID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除原有 farm-flightsheet 关联
		if err := tx.Delete(&model.FarmFlightsheet{}, "farm_id = ? AND flightsheet_id = ?", farmID, fsID).Error; err != nil {
			return err
		}
		// 建立新的 farm-flightsheet 关联
		farmFlightsheet := &model.FarmFlightsheet{
			FarmID:        farmID,
			FlightsheetID: fsID,
		}
		if err := tx.Create(farmFlightsheet).Error; err != nil {
			return err
		}
		return nil
	})
}

// TransferFarm 转移矿场，矿场下的矿机也会转移
func (dao *FarmDAO) TransferFarm(farmID int, fromUserID int, toUserID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 更新 user-farm 关联
		if err := tx.Model(&model.UserFarm{}).
			Where("user_id = ?", fromUserID).
			Update("user_id", toUserID).
			Error; err != nil {
			return err
		}
		// 更新 user-miner 关联
		if err := tx.Model(&model.UserMiner{}).
			Where("user_id = ?", fromUserID).
			Updates(map[string]interface{}{"user_id": toUserID}).
			Error; err != nil {
			return err
		}
		return nil
	})
}
