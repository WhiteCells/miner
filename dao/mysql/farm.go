package mysql

import (
	"miner/model"
	"miner/model/relation"
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
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 创建矿场
		if err := tx.Create(farm).Error; err != nil {
			return err
		}
		// 创建 用户-矿场 关联
		userFarm := &relation.UserFarm{
			UserID: userID,
			FarmID: farm.ID,
		}
		if err := tx.Create(userFarm).Error; err != nil {
			return err
		}
		return nil
	})
}

// DeleteFarmByID 删除矿场
func (dao *FarmDAO) DeleteFarmByID(farmID int, userID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除 用户-矿场 关联
		if err := tx.Delete(&relation.UserFarm{}, "user_id = ? AND farm_id = ?", userID, farmID).Error; err != nil {
			return err
		}
		// 删除 矿场-矿机 关联
		// TODO 矿机如何处理
		if err := tx.Delete(&relation.FarmMiner{}, "farm_id = ?", farmID).Error; err != nil {
			return err
		}
		// 删除矿场
		if err := tx.Delete(&model.Farm{}, farmID).Error; err != nil {
			return err
		}
		return nil
	})
}

// UpdateFarm 更新矿场信息
func (dao *FarmDAO) UpdateFarm(farm *model.Farm) error {
	return utils.DB.Save(farm).Error
}

// GetFarm 获取用户的矿场
func (dao *FarmDAO) GetFarm(userID int, query map[string]any) (*[]model.Farm, int64, error) {
	var farms []model.Farm
	var total int64

	// 查询总数
	if err := utils.DB.Model(relation.UserFarm{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, -1, err
	}

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)

	// 分页查询
	err := utils.DB.
		Joins("JOIN user_farm ON farm.id = user_farm.farm_id").
		Where("user_farm.user_id = ?", userID).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&farms).Error

	return &farms, total, err
}

func (m *FarmDAO) GetUserAllFarms(userID int) (*[]model.Farm, error) {
	var farms []model.Farm
	err := utils.DB.Find(&farms).Error
	return &farms, err
}

// 获取矿场
func (dao *FarmDAO) GetFarmByID(farmID int) (*model.Farm, error) {
	var farm model.Farm
	err := utils.DB.First(&farm, farmID).Error
	return &farm, err
}

// 矿场应用飞行表
func (dao *FarmDAO) ApplyFs(farmID int, fsID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除原有 farm-fs 关联
		if err := tx.Delete(&relation.FarmFs{}, "farm_id = ?", farmID).Error; err != nil {
			return err
		}
		// 建立新的 farm-fs 关联
		farmFlightsheet := &relation.FarmFs{
			FarmID: farmID,
			FsID:   fsID,
		}
		if err := tx.Create(farmFlightsheet).Error; err != nil {
			return err
		}
		// TODO 对矿场下没有设置飞行表的矿机的应用
		return nil
	})
}

// 转移矿场
func (dao *FarmDAO) TransferFarm(farmID int, fromUserID int, toUserID int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 更新 user-farm 关联
		if err := tx.Model(&relation.UserFarm{}).
			Where("user_id = ?", fromUserID).
			Update("user_id", toUserID).
			Error; err != nil {
			return err
		}
		// 更新 user-miner 关联
		if err := tx.Model(&relation.UserMiner{}).
			Where("user_id = ?", fromUserID).
			Updates(map[string]any{"user_id": toUserID}).
			Error; err != nil {
			return err
		}
		return nil
	})
}
