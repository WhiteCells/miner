package mysql

import (
	"miner/model"
	"miner/model/relation"
	"miner/utils"

	"gorm.io/gorm"
)

type MinerDAO struct{}

func NewMinerDAO() *MinerDAO {
	return &MinerDAO{}
}

// CreateMiner 创建矿机
func (dao *MinerDAO) CreateMiner(miner *model.Miner, userID int, farmID int) error {
	// 创建矿机时就需要将用户与用户与矿机进行联系
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 创建 miner
		if err := tx.Create(miner).Error; err != nil {
			return err
		}

		// 建立 user-miner 关联
		userMiner := &relation.UserMiner{
			UserID:  userID,
			MinerID: miner.ID,
			// Perm:    perm.MinerOwner,
		}
		if err := tx.Create(userMiner).Error; err != nil {
			return err
		}

		// 建立 farm-miner 关联
		farmMiner := &relation.FarmMiner{
			FarmID:  farmID,
			MinerID: miner.ID,
		}
		if err := tx.Create(farmMiner).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// GetMinerByID 通过矿机 ID 获取矿机信息
func (dao *MinerDAO) GetMinerByID(minerID int) (*model.Miner, error) {
	var miner model.Miner
	err := utils.DB.First(&miner, minerID).Error
	return &miner, err
}

// GetFarmMiners 获取矿场的所有矿机
func (dao *MinerDAO) GetFarmMiners(farmID int) (*[]model.Miner, error) {
	var miners []model.Miner
	err := utils.DB.
		Joins("JOIN farm_miner ON miner.id = farm_miner.miner_id").
		Where("farm_miner.farm_id = ?", farmID).
		Find(&miners).Error
	return &miners, err
}

// DeleteMiner 删除矿机
func (dao *MinerDAO) DeleteMiner(minerID int, farmID int, userID int) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 删除 user-miner 关联
		if err := tx.
			Where("user_id = ? AND miner_id = ?", userID, minerID).
			Delete(&relation.UserMiner{}).Error; err != nil {
			return err
		}
		// 删除 farm-miner 关联
		if err := tx.
			Where("miner_id = ?", minerID).
			Delete(&relation.FarmMiner{}).Error; err != nil {
			return err
		}
		// 删除 miner-flightsheet 关联
		if err := tx.
			Where("miner_id = ?", minerID).
			Delete(&relation.MinerFs{}).Error; err != nil {
			return err
		}
		// 删除 miner
		if err := tx.Delete(&model.Miner{}, minerID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMiner 更新矿机信息
func (dao *MinerDAO) UpdateMiner(miner *model.Miner) error {
	return utils.DB.Save(miner).Error
}

// UpdateMinerStatus 更新矿机状态
func (dao *MinerDAO) UpdateMinerStatus(minerID int, status int) error {
	return utils.DB.
		Model(&model.Miner{}).
		Where("id = ?", minerID).
		Update("status", status).Error
}

// GetMiner 获取矿机
func (dao *MinerDAO) GetMiner(userID int, query map[string]any) (*[]model.Miner, int64, error) {
	var miners []model.Miner
	var total int64

	pageNum := query["page_num"].(int)
	pageSize := query["page_size"].(int)
	farmID := query["farm_id"].(int)

	// 查询总数
	if err := utils.DB.
		Model(&model.Miner{}).
		Joins("JOIN farm_miner ON farm_miner.miner_id = miner.id").
		Joins("JOIN user_miner ON user_miner.miner_id = miner.id").
		Joins("JOIN user_farm ON user_farm.farm_id = farm_miner.farm_id").
		Where("user_farm.user_id = ? AND user_farm.farm_id = ? AND user_miner.user_id = ?", userID, farmID, userID).
		Count(&total).Error; err != nil {
		return nil, -1, err
	}

	// 分页查询
	err := utils.DB.
		Model(&model.Miner{}).
		Joins("JOIN farm_miner ON farm_miner.miner_id = miner.id").
		Joins("JOIN user_miner ON user_miner.miner_id = miner.id").
		Joins("JOIN user_farm ON user_farm.farm_id = farm_miner.farm_id").
		Where("user_farm.user_id = ? AND user_farm.farm_id = ? AND user_miner.user_id = ?", userID, farmID, userID).
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&miners).Error

	return &miners, total, err
}

// Transfer 转移矿机所有权
func (dao *MinerDAO) Transfer(minerID int, fromFarmID int, farmHash string) error {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		// 通过 Farm hash 找到矿场
		var farm model.Farm
		if err := tx.
			Where("hash = ?", farmHash).
			First(&farm).Error; err != nil {
			return err
		}
		// 更新 farm-miner 关联
		if err := tx.
			Model(&relation.FarmMiner{}).
			Where("farm_id = ? AND miner_id = ?", fromFarmID, minerID).
			Update("farm_id", farm.ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
